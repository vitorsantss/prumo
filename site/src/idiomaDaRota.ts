import { defineRouteMiddleware, type StarlightRouteData } from '@astrojs/starlight/route-data';
import { getCollection } from 'astro:content';

type Entrada = StarlightRouteData['sidebar'][number];

const base = import.meta.env.BASE_URL.replace(/\/?$/, '/');

export const ehDesignSystem = (id: string) => /^(en\/)?docs\/design(\/|$)/.test(id);

const idDoLink = (href: string) => href.slice(base.length).replace(/\/$/, '');

const marcar = (entrada: Entrada, traduzidas: Set<string>): Entrada => {
  if (entrada.type === 'group') return { ...entrada, entries: entrada.entries.map((filha) => marcar(filha, traduzidas)) };
  if (traduzidas.has(idDoLink(entrada.href))) return entrada;
  return {
    ...entrada,
    attrs: { ...entrada.attrs, lang: 'pt-BR' },
    badge: { text: 'PT', variant: 'default', class: 'selo-idioma' },
  };
};

export const onRequest = defineRouteMiddleware(async (contexto) => {
  const rota = contexto.locals.starlightRoute;
  if (rota.toc) rota.toc.items = rota.toc.items.filter((item) => item.slug !== 'footnote-label');
  const grupoDS = rota.sidebar.findIndex((entrada) => entrada.type === 'group' && entrada.label === 'Design system');
  if (grupoDS >= 0) {
    const grupo = rota.sidebar[grupoDS];
    rota.sidebar = ehDesignSystem(rota.id) && grupo.type === 'group' ? grupo.entries : rota.sidebar.filter((_, i) => i !== grupoDS);
  }
  if (rota.locale !== 'en') return;
  const traduzidas = new Set((await getCollection('docs', (e) => e.id.startsWith('en/'))).map((e) => e.id));
  rota.sidebar = rota.sidebar.map((entrada) => marcar(entrada, traduzidas));
});
