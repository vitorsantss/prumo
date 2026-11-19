# 0010 · Site e documentação com Starlight

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
O projeto precisa de uma página que explique a ideia e de uma documentação navegável, com busca, em português e inglês, sem esconder que ainda está em estudo. Os textos já vivem em Markdown em `docs/` e devem continuar legíveis direto no GitHub. O Material for MkDocs, opção mais comum até pouco tempo, entrou em modo de manutenção: só correções críticas e de segurança por pelo menos 12 meses, sem recursos novos, e o próprio time recomenda migrar pro Zensical.

## Decisão
O site é feito com **Starlight** (Astro, MIT), na pasta `site/`:
- `docs/` continua sendo a fonte. Um script (`site/scripts/sincronizar-docs.mjs`) copia os Markdown pro site no build, gera o título a partir do primeiro cabeçalho e reescreve os links. Ninguém edita a cópia.
- Páginas em inglês ficam em `docs/en/`, com o mesmo nome do arquivo em português. Página sem tradução aparece em português, com o aviso do próprio Starlight.
- A página inicial é própria (`site/src/components/Landing.astro`), com o visual do design system. O que muda com o projeto (hipóteses, linha do tempo) vem de `site/src/data/landing.json`.
- Cabeçalho, busca, seletor de idioma, tema e aviso de página sem tradução são componentes nossos, no visual do design system (regras em `docs/design/site.md`). A busca usa a API do Pagefind direto, pra página inicial e documentação terem o mesmo modal.
- Endereço e caminho base vêm das variáveis `PRUMO_SITE_URL` e `PRUMO_SITE_BASE`, pra publicar no GitHub Pages ou em outro serviço sem mudar código.

Alternativas avaliadas: Astro puro, sem o Starlight (visual todo nosso, mas busca, rotas de idioma e sidebar viram código pra manter; custo inicial 3 a 4 vezes maior), Hugo com Hextra (sem Node, mas templates numa linguagem a mais), VitePress (ecossistema Vue, diferente do React planejado pra interface) e Zensical (sucessor do Material, ainda novo).

## Consequências
- O repositório passa a ter Node, só dentro de `site/`. A interface futura do Prumo também será TypeScript e React, então não é uma dependência a mais no longo prazo.
- Trocar de gerador é trocar a moldura: o conteúdo é Markdown comum em `docs/`.
- A busca (Pagefind) roda no próprio site, sem serviço externo.
- Publicar depende de um serviço de hospedagem. Hoje o site está no Cloudflare Pages (https://prumo-docs.pages.dev), que publica a cada push na `main`, porque o Actions da conta está parado. O caminho livre é o GitHub Pages ([ADR-0005](0005-integracoes-terceirizadas-caminho-nativo.md)), acompanhado na #61.

## Quando reavaliar
O Starlight ainda está na versão 0.x e muda entre versões. A versão fica travada no `package.json` e só sobe de propósito, lendo o changelog. Migrar pra Astro puro entra em pauta se uma atualização quebrar um componente trocado duas vezes seguidas, se precisarmos trocar mais da metade dos componentes, ou se uma necessidade do Prumo não couber no Starlight. Acompanhamento na issue #58.

## Fontes
- Starlight · https://starlight.astro.build
- Material for MkDocs, aviso de modo de manutenção · https://github.com/squidfunk/mkdocs-material/issues/8523
- Pagefind · https://pagefind.app
- Starlight, componentes substituíveis · https://starlight.astro.build/reference/overrides
