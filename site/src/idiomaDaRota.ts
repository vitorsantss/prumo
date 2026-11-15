import { defineRouteMiddleware, type StarlightRouteData } from '@astrojs/starlight/route-data';
import { getCollection } from 'astro:content';

type Entrada = StarlightRouteData['sidebar'][number];

const base = import.meta.env.BASE_URL.replace(/\/?$/, '/');

const idDoLink = (href: string) => href.slice(base.length).replace(/\/$/, '');

const paraPortugues = (entrada: Entrada): Entrada => {
  if (entrada.type === 'group') return { ...entrada, entries: entrada.entries.map(paraPortugues) };
  const id = idDoLink(entrada.href).replace(/^en\//, '');
  return {
    ...entrada,
    href: `${base}${id}/`,
    isCurrent: false,
    attrs: { ...entrada.attrs, lang: 'pt-BR', hreflang: 'pt-BR' },
    badge: { text: 'PT', variant: 'default', class: 'selo-idioma' },
  };
};

const separar = (entradas: Entrada[], traduzidas: Set<string>): { ficam: Entrada[]; saem: Entrada[] } => {
  const ficam: Entrada[] = [];
  const saem: Entrada[] = [];
  for (const entrada of entradas) {
    if (entrada.type === 'link') {
      (traduzidas.has(idDoLink(entrada.href)) ? ficam : saem).push(entrada);
      continue;
    }
    const filhos = separar(entrada.entries, traduzidas);
    if (filhos.ficam.length === 0) {
      saem.push(entrada);
      continue;
    }
    ficam.push({ ...entrada, entries: filhos.ficam });
    saem.push(...filhos.saem);
  }
  return { ficam, saem };
};

export const onRequest = defineRouteMiddleware(async (contexto) => {
  const rota = contexto.locals.starlightRoute;
  if (rota.locale !== 'en') return;

  if (rota.isFallback) rota.toc = undefined;

  const traduzidas = new Set((await getCollection('docs', (e) => e.id.startsWith('en/'))).map((e) => e.id));
  const { ficam, saem } = separar(rota.sidebar, traduzidas);
  rota.sidebar = saem.length
    ? [...ficam, { type: 'group', label: 'Only in Portuguese', entries: saem.map(paraPortugues), collapsed: false, badge: undefined }]
    : ficam;
});
