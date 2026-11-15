import { readFileSync, writeFileSync, mkdirSync, rmSync, readdirSync, statSync } from 'node:fs';
import { join, dirname, relative, posix } from 'node:path';
import { fileURLToPath } from 'node:url';

const aqui = dirname(fileURLToPath(import.meta.url));
const origem = join(aqui, '..', '..', 'docs');
const destino = join(aqui, '..', 'src', 'content', 'docs');
const repositorio = 'https://github.com/vitorsantss/prumo/blob/main/';
const ignorar = new Set(['design']);
const marcaIA = /^> \*\*Translated by AI[^\n]*\n?/m;

function listarMarkdown(pasta, base = '') {
  return readdirSync(pasta).flatMap((nome) => {
    const caminho = join(pasta, nome);
    const rel = posix.join(base, nome);
    if (statSync(caminho).isDirectory()) return ignorar.has(rel) ? [] : listarMarkdown(caminho, rel);
    return nome.endsWith('.md') ? [rel] : [];
  });
}

function rota(rel) {
  const semExt = rel.replace(/\.md$/, '').replace(/(^|\/)README$/, '$1');
  const idioma = semExt.startsWith('en/') ? 'en/' : '';
  const resto = idioma ? semExt.slice(3) : semExt;
  return `/${idioma}docs/${resto}`.replace(/\/+$/, '') + '/';
}

function reescreverLinks(texto, rel) {
  const daqui = rota(rel);
  return texto.replace(/\]\((?!https?:|#|mailto:)([^)\s]+?)(#[^)]*)?\)/g, (todo, alvo, ancora = '') => {
    const resolvido = posix.normalize(posix.join(posix.dirname(rel), alvo));
    if (resolvido.startsWith('..')) {
      const noRepo = posix.normalize(posix.join('docs', posix.dirname(rel), alvo));
      return `](${repositorio}${noRepo}${ancora})`;
    }
    const ehPasta = alvo.endsWith('/');
    if (!ehPasta && !resolvido.endsWith('.md')) return `](${repositorio}docs/${resolvido}${ancora})`;
    const destinoRota = ehPasta ? rota(posix.join(resolvido, 'README.md')) : rota(resolvido);
    let relativo = posix.relative(daqui, destinoRota) || '.';
    if (!relativo.endsWith('/')) relativo += '/';
    return `](${relativo}${ancora})`;
  });
}

const adrs = new Map(
  readdirSync(join(origem, 'adr'))
    .filter((nome) => /^\d{4}-.+\.md$/.test(nome))
    .map((nome) => [nome.slice(0, 4), nome]),
);
const issues = 'https://github.com/vitorsantss/prumo/issues/';
const adrsSemArquivo = new Set();

function ligarCitacoes(texto, rel) {
  const pasta = posix.dirname(rel);
  const prefixoAdr = rel.startsWith('en/') ? 'en/adr/' : 'adr/';
  const ligar = (trecho) =>
    trecho
      .replace(/(?<![\w/-])ADR-(\d{4})(?![\w-])/g, (todo, numero) => {
        const arquivo = adrs.get(numero);
        if (!arquivo) adrsSemArquivo.add(`${todo} (em docs/${rel})`);
        if (!arquivo || rel.endsWith(arquivo)) return todo;
        return `[${todo}](${posix.relative(pasta, prefixoAdr + arquivo)})`;
      })
      .replace(/(?<![\w/&\[])#(\d{1,4})\b/g, (todo, numero) => `[${todo}](${issues}${numero})`);
  const protegido = /(\[[^\]]*\]\([^)]*\)|`[^`\n]*`|<[^>\n]+>|https?:\/\/\S+)/;
  return texto
    .split(/(^```[\s\S]*?^```)/m)
    .map((bloco) => (bloco.startsWith('```') ? bloco : bloco.split(protegido).map((parte, i) => (i % 2 ? parte : ligar(parte))).join('')))
    .join('');
}

function converter(rel) {
  const bruto = readFileSync(join(origem, rel), 'utf8').replace(/\r\n/g, '\n');
  const linhas = bruto.split('\n');
  const indiceTitulo = linhas.findIndex((l) => l.startsWith('# '));
  const titulo = indiceTitulo >= 0 ? linhas[indiceTitulo].slice(2).trim() : rel;
  if (indiceTitulo >= 0) linhas.splice(indiceTitulo, 1);
  const corpo = reescreverLinks(ligarCitacoes(linhas.join('\n').trimStart(), rel), rel);
  const traducao = rel.startsWith('en/') ? (marcaIA.test(corpo) ? 'ia' : 'revisada') : null;
  const texto = corpo.replace(marcaIA, '').trimStart();
  const frontmatter = ['---', `title: ${JSON.stringify(titulo)}`, ...(traducao ? [`traducao: ${traducao}`] : []), ...(rel.endsWith('0000-modelo.md') ? ['sidebar:', '  hidden: true', 'pagefind: false'] : []), ...(/^(en\/)?adr\/README\.md$/.test(rel) ? ['sidebar:', '  label: ' + (rel.startsWith('en/') ? '"All decisions"' : '"Todas as decisões"'), '  order: 0'] : []), `origem: ${JSON.stringify('docs/' + rel)}`, `editUrl: ${JSON.stringify(repositorio + 'docs/' + rel)}`, '---', ''];
  return frontmatter.join('\n') + texto;
}

rmSync(destino, { recursive: true, force: true });
for (const rel of listarMarkdown(origem)) {
  const saidaRel = rel.startsWith('en/') ? posix.join('en', 'docs', rel.slice(3)) : posix.join('docs', rel);
  const saida = join(destino, saidaRel.replace(/(^|\/)README\.md$/, '$1index.md'));
  mkdirSync(dirname(saida), { recursive: true });
  writeFileSync(saida, converter(rel));
}
console.log(`docs sincronizados em ${relative(process.cwd(), destino)}`);
if (adrsSemArquivo.size) console.warn(`ADRs citados que ainda não existem em docs/adr (ficam sem link):\n  ${[...adrsSemArquivo].join('\n  ')}`);
