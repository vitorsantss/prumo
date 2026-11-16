import { readFileSync, writeFileSync } from 'node:fs';
import { join, dirname } from 'node:path';
import { fileURLToPath } from 'node:url';

const aqui = dirname(fileURLToPath(import.meta.url));
const tokens = JSON.parse(readFileSync(join(aqui, '..', '..', 'docs', 'design', 'tokens.json'), 'utf8'));
const saida = join(aqui, '..', 'src', 'estilos', 'tokens-gerados.css');

const folhas = (grupo) => Object.entries(grupo).filter(([nome]) => !nome.startsWith('$'));
const numero = (n) => Number(n.toFixed(4)).toString();

export function corCss({ components, alpha, hex }) {
  if (alpha === undefined || alpha === 1) return hex;
  const [r, g, b] = components.map((c) => Math.round(c * 255));
  return `rgba(${r}, ${g}, ${b}, ${numero(alpha)})`;
}
export const medidaCss = ({ value, unit }) => `${value}${unit}`;
const sombraCss = (valor) =>
  [valor].flat().map((s) => `${medidaCss(s.offsetX)} ${medidaCss(s.offsetY)} ${medidaCss(s.blur)} ${medidaCss(s.spread)} ${corCss(s.color)}`).join(', ');
const fonteCss = (lista) => lista.map((f) => (f.includes(' ') ? `'${f}'` : f)).join(', ');

function porTema(tema) {
  return [
    ...folhas(tokens.cor[tema]).map(([nome, t]) => `  --p-${nome}: ${corCss(t.$value)};`),
    ...folhas(tokens.sombra[tema]).map(([nome, t]) => `  --p-sombra-${nome}: ${sombraCss(t.$value)};`),
  ].join('\n');
}

const comuns = [
  ...folhas(tokens.espaco).map(([nome, t]) => `  --p-espaco-${nome}: ${medidaCss(t.$value)};`),
  ...folhas(tokens.raio).map(([nome, t]) => `  --p-raio-${nome}: ${medidaCss(t.$value)};`),
  ...folhas(tokens.duracao).map(([nome, t]) => `  --p-duracao-${nome}: ${medidaCss(t.$value)};`),
  ...folhas(tokens.curva).map(([nome, t]) => `  --p-curva-${nome}: cubic-bezier(${t.$value.join(', ')});`),
  ...folhas(tokens.fonte).map(([nome, t]) => `  --p-fonte-${nome}: ${fonteCss(t.$value)};`),
].join('\n');

const css = `/* Gerado por site/scripts/gerar-tokens.mjs a partir de docs/design/tokens.json. Não edite. */
:root {
${comuns}
}

:root,
.p-tema-escuro {
${porTema('escuro')}
}

:root[data-theme='light'],
.p-tema-claro {
${porTema('claro')}
}
`;

if (process.argv[1] === fileURLToPath(import.meta.url)) {
  writeFileSync(saida, css);
  console.log(`tokens: ${saida}`);
}
