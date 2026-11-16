import { spawn, execFileSync } from 'node:child_process';
import { mkdtempSync, readFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

const PAGINAS = ['', 'docs/a-ideia/', 'docs/estado-atual/', 'docs/referencias/', 'en/docs/adr/0001-arquivos-como-fonte-da-verdade/'];
const MODOS = ['mobile', 'desktop'];
const MINIMO = 90;
const base = (process.env.PRUMO_SITE_BASE ?? '/prumo').replace(/\/?$/, '/');
const origem = process.argv[2] ?? `http://localhost:4321${base}`;
const pasta = mkdtempSync(join(tmpdir(), 'prumo-auditoria-'));

const servidor = process.argv[2] ? null : spawn('npx', ['astro', 'preview', '--port', '4321'], { stdio: 'ignore' });
if (servidor) await new Promise((pronto) => setTimeout(pronto, 3000));

const abaixo = [];
try {
  for (const pagina of PAGINAS) {
    for (const modo of MODOS) {
      const saida = join(pasta, `${pagina.replaceAll('/', '_') || 'inicio'}-${modo}.json`);
      execFileSync('npx', [
        '--yes', 'lighthouse@13', `${origem}${pagina}`,
        ...(modo === 'desktop' ? ['--preset=desktop'] : []),
        '--quiet', '--output=json', `--output-path=${saida}`,
        '--chrome-flags=--headless=new --no-sandbox',
      ], { stdio: 'ignore' });
      const { categories } = JSON.parse(readFileSync(saida, 'utf8'));
      const notas = Object.values(categories).map((c) => [c.title, Math.round((c.score ?? 0) * 100)]);
      console.log(`${(pagina || '/').padEnd(52)} ${modo.padEnd(8)} ${notas.map(([t, n]) => `${t} ${n}`).join(' · ')}`);
      notas.filter(([, n]) => n < MINIMO).forEach(([t, n]) => abaixo.push(`${pagina || '/'} (${modo}): ${t} ${n}`));
    }
  }
} finally {
  servidor?.kill();
}

if (abaixo.length) {
  console.error(`\nAbaixo de ${MINIMO}:\n${abaixo.join('\n')}\nRelatórios completos em ${pasta}`);
  process.exit(1);
}
console.log(`\nTudo acima de ${MINIMO}. Relatórios completos em ${pasta}`);
