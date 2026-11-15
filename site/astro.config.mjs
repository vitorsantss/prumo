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
      customCss: ['./src/estilos/tokens.css', './src/estilos/prumo.css'],
      lastUpdated: false,
      pagination: false,
      routeMiddleware: './src/idiomaDaRota.ts',
      components: {
        PageFrame: './src/components/starlight/Moldura.astro',
        TwoColumnContent: './src/components/starlight/Colunas.astro',
        EditLink: './src/components/starlight/Origem.astro',
        MarkdownContent: './src/components/starlight/Conteudo.astro',
        PageTitle: './src/components/starlight/Titulo.astro',
        Header: './src/components/starlight/Cabecalho.astro',
        Search: './src/components/starlight/Pesquisa.astro',
        LanguageSelect: './src/components/starlight/Idioma.astro',
        ThemeSelect: './src/components/starlight/Tema.astro',
        FallbackContentNotice: './src/components/starlight/AvisoSemTraducao.astro',
      },
      sidebar: [
        {
          label: 'Comece',
          translations: { en: 'Start' },
          items: ['docs/comece-aqui', 'docs/primeira-contribuicao'],
        },
        {
          label: 'Entenda',
          translations: { en: 'Understand' },
          items: [
            'docs/a-ideia',
            'docs/processo',
            { label: 'Decisões (ADRs)', translations: { en: 'Decisions (ADRs)' }, collapsed: true, items: [{ autogenerate: { directory: 'docs/adr' } }] },
          ],
        },
        {
          label: 'Faça',
          translations: { en: 'Do' },
          items: ['docs/formas-de-contribuir', 'docs/traduzir', 'docs/como-mudar-o-site'],
        },
        {
          label: 'Consulte',
          translations: { en: 'Look up' },
          items: ['docs/estado-atual', 'docs/referencias', 'docs'],
        },
      ],
    }),
  ],
});
