import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import { menu } from './sidebar.mjs';

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
      head: [
        { tag: 'link', attrs: { rel: 'expect', href: '#starlight__sidebar', blocking: 'render', media: '(min-width: 50rem)' } },
        { tag: 'link', attrs: { rel: 'expect', href: '#_top', blocking: 'render' } },
        { tag: 'link', attrs: { rel: 'expect', href: '#starlight__on-this-page', blocking: 'render', media: '(min-width: 72rem)' } },
      ],
      lastUpdated: false,
      pagination: false,
      routeMiddleware: './src/idiomaDaRota.ts',
      components: {
        Head: './src/components/starlight/Cabeca.astro',
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
      sidebar: menu,
    }),
  ],
});
