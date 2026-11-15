import { readFileSync, writeFileSync, mkdirSync, rmSync, readdirSync, statSync } from 'node:fs';
import { join, dirname, relative, posix } from 'node:path';
import { fileURLToPath } from 'node:url';

const aqui = dirname(fileURLToPath(import.meta.url));
const origem = join(aqui, '..', '..', 'docs');
const destino = join(aqui, '..', 'src', 'content', 'docs');
const repositorio = 'https://github.com/vitorsantss/prumo/blob/main/';
const ignorar = new Set(['design']);
const marcaIA = /^> \*\*Translated by AI/m;

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

function converter(rel) {
  const bruto = readFileSync(join(origem, rel), 'utf8').replace(/\r\n/g, '\n');
  const linhas = bruto.split('\n');
  const indiceTitulo = linhas.findIndex((l) => l.startsWith('# '));
  const titulo = indiceTitulo >= 0 ? linhas[indiceTitulo].slice(2).trim() : rel;
  if (indiceTitulo >= 0) linhas.splice(indiceTitulo, 1);
  const corpo = reescreverLinks(linhas.join('\n').trimStart(), rel);
  const traducao = rel.startsWith('en/') ? (marcaIA.test(corpo) ? 'ia' : 'revisada') : null;
  const frontmatter = ['---', `title: ${JSON.stringify(titulo)}`, ...(traducao ? [`traducao: ${traducao}`] : []), ...(rel.endsWith('0000-modelo.md') ? ['sidebar:', '  hidden: true', 'pagefind: false'] : []), `editUrl: ${JSON.stringify(repositorio.replace('/blob/', '/edit/') + 'docs/' + rel)}`, '---', ''];
  return frontmatter.join('\n') + corpo;
}

rmSync(destino, { recursive: true, force: true });
for (const rel of listarMarkdown(origem)) {
  const saidaRel = rel.startsWith('en/') ? posix.join('en', 'docs', rel.slice(3)) : posix.join('docs', rel);
  const saida = join(destino, saidaRel.replace(/(^|\/)README\.md$/, '$1index.md'));
  mkdirSync(dirname(saida), { recursive: true });
  writeFileSync(saida, converter(rel));
}
console.log(`docs sincronizados em ${relative(process.cwd(), destino)}`);
