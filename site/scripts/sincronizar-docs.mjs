import { readFileSync, writeFileSync, mkdirSync, rmSync, readdirSync, statSync, existsSync } from 'node:fs';
import { join, dirname, relative, posix } from 'node:path';
import { corCss, medidaCss } from './gerar-tokens.mjs';
import { fileURLToPath } from 'node:url';

const aqui = dirname(fileURLToPath(import.meta.url));
const origem = join(aqui, '..', '..', 'docs');
const destino = join(aqui, '..', 'src', 'content', 'docs');
const repositorio = 'https://github.com/vitorsantss/prumo/blob/main/';
const ignorar = new Set(['design/sistema']);
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

const raiz = join(aqui, '..', '..');
const tokens = JSON.parse(readFileSync(join(origem, 'design', 'tokens.json'), 'utf8'));
const folhas = (grupo) => Object.entries(grupo).filter(([nome]) => !nome.startsWith('$'));
const esc = (t) => String(t).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
const tabela = (cabecalho, linhas) =>
  `<table class="empilha tabela-tokens">\n<thead><tr>${cabecalho.map((c) => `<th>${c}</th>`).join('')}</tr></thead>\n<tbody>\n${linhas.map((l) => `<tr>${l.map((c, i) => `<td data-rotulo="${cabecalho[i]}">${c}</td>`).join('')}</tr>`).join('\n')}\n</tbody>\n</table>`;
const amostra = (valor) => `<span class="amostra"><span class="cor" style="background: ${valor}" aria-hidden="true"></span><code>${esc(valor)}</code></span>`;

const geradores = {
  cor: () => tabela(['Token', 'Claro', 'Escuro', 'Uso'], folhas(tokens.cor.claro).map(([nome, t]) => [`<code>--p-${nome}</code>`, amostra(corCss(t.$value)), amostra(corCss(tokens.cor.escuro[nome].$value)), esc(t.$description)])),
  espaco: () => tabela(['Token', 'Valor', 'Uso'], folhas(tokens.espaco).map(([nome, t]) => [`<code>--p-espaco-${nome}</code>`, `<span class="regua" style="width: ${medidaCss(t.$value)}" aria-hidden="true"></span> ${medidaCss(t.$value)}`, esc(t.$description)])),
  raio: () => tabela(['Token', 'Valor', 'Uso'], folhas(tokens.raio).map(([nome, t]) => [`<code>--p-raio-${nome}</code>`, `<span class="canto" style="border-radius: ${medidaCss(t.$value)}" aria-hidden="true"></span> ${medidaCss(t.$value)}`, esc(t.$description)])),
  sombra: () => tabela(['Token', 'Uso'], folhas(tokens.sombra.claro).map(([nome, t]) => [`<code>--p-sombra-${nome}</code>`, esc(t.$description)])),
  movimento: () => tabela(['Token', 'Valor', 'Uso'], [...folhas(tokens.duracao).map(([nome, t]) => [`<code>--p-duracao-${nome}</code>`, medidaCss(t.$value), esc(t.$description)]), ...folhas(tokens.curva).map(([nome, t]) => [`<code>--p-curva-${nome}</code>`, `cubic-bezier(${t.$value.join(', ')})`, esc(t.$description)])]),
  tipografia: () => {
    const familias = { '{fonte.manuscrita}': 'var(--p-fonte-manuscrita)', '{fonte.texto}': 'var(--p-fonte-texto)', '{fonte.mono}': 'var(--p-fonte-mono)' };
    return tabela(['Estilo', 'Amostra', 'Tamanho e peso', 'Uso'], folhas(tokens.tipografia).map(([nome, t]) => {
      const v = t.$value;
      return [`<code>${nome}</code>`, `<span style="font-family: ${familias[v.fontFamily]}; font-size: ${medidaCss(v.fontSize)}; font-weight: ${v.fontWeight}; line-height: ${v.lineHeight}">Prumo</span>`, `${medidaCss(v.fontSize)} · ${v.fontWeight} · ${v.lineHeight}`, esc(t.$description)];
    }));
  },
};

function trechoDeCodigo(marca, alvo) {
  const [caminho, classes] = alvo.split('#');
  if (!existsSync(join(raiz, caminho))) return marca;
  const fonte = readFileSync(join(raiz, caminho), 'utf8').replace(/\r\n/g, '\n');
  const url = `${repositorio}${caminho}`;
  const partes = [];
  if (caminho.endsWith('.css')) {
    const regras = (classes ?? '').split(',').flatMap((classe) => [...fonte.matchAll(new RegExp(`^[^{}\\n]*\\.${classe}(?![\\w-])[^{}\\n]*\\{[^}]*\\}`, 'gm'))].map((m) => m[0]));
    partes.push(`Trecho de [\`${caminho}\`](${url}), extraído no build.`, '', '```css', [...new Set(regras)].join('\n\n'), '```');
  } else {
    const semFrontmatter = fonte.replace(/^---\n[\s\S]*?\n---\n/, '');
    const template = semFrontmatter.split(/\n<(?:script|style)[\s>]/)[0].trim();
    const estilo = (semFrontmatter.match(/<style[^>]*>\n?([\s\S]*?)<\/style>/) ?? [])[1]?.replace(/^ {2}/gm, '').trim();
    partes.push(`Trecho de [\`${caminho}\`](${url}), extraído no build. As chaves \`{ }\` são valores que o Astro preenche ao gerar a página.`, '', '```html', template, '```');
    if (estilo) partes.push('', '<details>', `<summary>CSS do componente (${estilo.split('\n').length} linhas)</summary>`, '', '```css', estilo, '```', '', '</details>');
  }
  partes.push('', 'Os trechos de código desta página podem ser copiados sem obrigação, em [CC0](https://creativecommons.org/publicdomain/zero/1.0/deed.pt-br), conforme o ADR-0011.');
  return partes.join('\n');
}

function expandirMarcadores(texto) {
  return texto
    .replace(/<!-- tokens:([a-z]+) -->/g, (marca, tipo) => (geradores[tipo] ? geradores[tipo]() : marca))
    .replace(/<!-- codigo:([^ ]+) -->/g, (marca, alvo) => trechoDeCodigo(marca, alvo));
}

function converter(rel) {
  const bruto = readFileSync(join(origem, rel), 'utf8').replace(/\r\n/g, '\n');
  const linhas = bruto.split('\n');
  const indiceTitulo = linhas.findIndex((l) => l.startsWith('# '));
  const titulo = indiceTitulo >= 0 ? linhas[indiceTitulo].slice(2).trim() : rel;
  if (indiceTitulo >= 0) linhas.splice(indiceTitulo, 1);
  const corpo = reescreverLinks(ligarCitacoes(linhas.join('\n').trimStart(), rel), rel);
  const traducao = rel.startsWith('en/') ? (marcaIA.test(corpo) ? 'ia' : 'revisada') : null;
  const texto = expandirMarcadores(corpo.replace(marcaIA, '').trimStart());
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
