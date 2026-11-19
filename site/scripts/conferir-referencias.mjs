import { readdir, readFile, stat } from 'node:fs/promises';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const raiz = resolve(dirname(fileURLToPath(import.meta.url)), '../..');
const alvos = ['docs', 'AGENTS.md'];
const conferirExternos = process.argv.includes('--externos');

const erros = [];
const avisos = [];

async function arquivosMarkdown(caminho) {
  const info = await stat(caminho);
  if (info.isFile()) return caminho.endsWith('.md') ? [caminho] : [];
  const entradas = await readdir(caminho, { withFileTypes: true });
  const listas = await Promise.all(entradas.map((e) => arquivosMarkdown(join(caminho, e.name))));
  return listas.flat();
}

function semCodigo(texto) {
  return texto
    .replace(/^```[\s\S]*?^```/gm, (bloco) => bloco.replace(/[^\n]/g, ' '))
    .replace(/`[^`\n]*`/g, (trecho) => ' '.repeat(trecho.length));
}

function semLinks(linha) {
  return linha.replace(/!?\[[^\]]*\]\([^)]*\)/g, (link) => ' '.repeat(link.length));
}

function registrar(lista, arquivo, numeroDaLinha, mensagem) {
  lista.push(`${relative(raiz, arquivo)}:${numeroDaLinha}: ${mensagem}`);
}

async function existe(caminho) {
  try {
    await stat(caminho);
    return true;
  } catch {
    return false;
  }
}

async function abre(url) {
  try {
    const resposta = await fetch(url, { method: 'GET', redirect: 'follow', signal: AbortSignal.timeout(15000) });
    return resposta.ok || resposta.status === 403 || resposta.status === 429;
  } catch {
    return false;
  }
}

async function conferir(arquivo, externos) {
  const linhas = semCodigo(await readFile(arquivo, 'utf8')).split('\n');
  const citadas = new Map();
  const definidas = new Map();

  for (const [indice, linha] of linhas.entries()) {
    const n = indice + 1;
    const definicao = linha.match(/^\[\^([^\]]+)\]:/);
    if (definicao) definidas.set(definicao[1], n);

    const corpo = definicao ? linha.slice(definicao[0].length) : linha;
    for (const citacao of corpo.matchAll(/\[\^([^\]]+)\]/g)) {
      if (!citadas.has(citacao[1])) citadas.set(citacao[1], n);
    }

    if (!linha.startsWith('#')) {
      const fora = semLinks(corpo);
      for (const adr of fora.matchAll(/\bADR-\d{4}\b/g)) {
        registrar(erros, arquivo, n, `${adr[0]} sem link`);
      }
      for (const numero of fora.matchAll(/(?<![\w/&#])#(\d+)\b/g)) {
        registrar(avisos, arquivo, n, `#${numero[1]} sem link (a classificação dos antigos é a #121)`);
      }
    }

    for (const link of linha.matchAll(/!?\[[^\]]*\]\(([^)\s]+)[^)]*\)/g)) {
      const destino = link[1];
      if (/^https?:\/\//.test(destino)) {
        externos.push({ arquivo, n, destino });
        continue;
      }
      if (/^(mailto:|#|\/)/.test(destino)) continue;
      const caminho = resolve(dirname(arquivo), decodeURIComponent(destino.split('#')[0]));
      if (!(await existe(caminho))) registrar(erros, arquivo, n, `link interno quebrado: ${destino}`);
    }
  }

  for (const [nome, n] of citadas) {
    if (!definidas.has(nome)) registrar(erros, arquivo, n, `nota [^${nome}] citada e não definida`);
  }
  for (const [nome, n] of definidas) {
    if (!citadas.has(nome)) registrar(erros, arquivo, n, `nota [^${nome}] definida e não citada`);
  }
}

const arquivos = (await Promise.all(alvos.map((a) => arquivosMarkdown(join(raiz, a))))).flat().sort();
const externos = [];
for (const arquivo of arquivos) await conferir(arquivo, externos);

if (conferirExternos) {
  const unicos = [...new Map(externos.map((e) => [e.destino, e])).values()];
  const resultados = await Promise.all(unicos.map(async (e) => ({ ...e, ok: await abre(e.destino) })));
  for (const r of resultados.filter((r) => !r.ok)) registrar(erros, r.arquivo, r.n, `link externo não abriu: ${r.destino}`);
}

for (const aviso of avisos) console.warn(`aviso  ${aviso}`);
for (const erro of erros) console.error(`erro   ${erro}`);
console.log(`\n${arquivos.length} arquivos · ${erros.length} erros · ${avisos.length} avisos`);
if (erros.length > 0) process.exit(1);
