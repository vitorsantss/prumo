import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

const site = process.env.PRUMO_SITE_URL ?? 'https://vitorsantss.github.io';
const base = process.env.PRUMO_SITE_BASE ?? '/prumo';

export default defineConfig({
  site,
  base,
  trailingSlash: 'always',
  build: { assets: 'recursos' },
  integrations: [
    starlight({
      title: 'Prumo',
      description: 'Pesquisa aberta sobre construir software com IA em etapas com portões humanos. Em estudo, ainda não pronto pra uso.',
      defaultLocale: 'root',
      locales: {
        root: { label: 'Português', lang: 'pt-BR' },
        en: { label: 'English', lang: 'en' },
      },
      social: [{ icon: 'github', label: 'GitHub', href: 'https://github.com/vitorsantss/prumo' }],
      customCss: ['./src/estilos/prumo.css'],
      lastUpdated: false,
      sidebar: [
        {
          label: '1 · Entenda a ideia',
          translations: { en: '1 · Understand the idea' },
          items: ['docs/comece-aqui', 'docs/a-ideia', 'docs/processo'],
        },
        {
          label: '2 · Veja onde está',
          translations: { en: '2 · See where it stands' },
          items: [
            'docs/estado-atual',
            { label: 'Decisões (ADRs)', translations: { en: 'Decisions (ADRs)' }, collapsed: true, items: [{ autogenerate: { directory: 'docs/adr' } }] },
          ],
        },
        {
          label: '3 · Participe',
          translations: { en: '3 · Take part' },
          items: ['docs/formas-de-contribuir', 'docs/primeira-contribuicao'],
        },
        {
          label: 'Referência',
          translations: { en: 'Reference' },
          items: ['docs/referencias', 'docs'],
        },
      ],
    }),
  ],
});
