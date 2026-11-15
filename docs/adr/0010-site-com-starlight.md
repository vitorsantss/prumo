# 0010 · Site e documentação com Starlight

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
O projeto precisa de uma página que explique a ideia e de uma documentação navegável, com busca, em português e inglês, sem esconder que ainda está em estudo. Os textos já vivem em Markdown em `docs/` e devem continuar legíveis direto no GitHub. O Material for MkDocs, opção mais comum até pouco tempo, chegou ao fim da vida em 05/11/2026.

## Decisão
O site é feito com **Starlight** (Astro, MIT), na pasta `site/`:
- `docs/` continua sendo a fonte. Um script (`site/scripts/sincronizar-docs.mjs`) copia os Markdown pro site no build, gera o título a partir do primeiro cabeçalho e reescreve os links. Ninguém edita a cópia.
- Páginas em inglês ficam em `docs/en/`, com o mesmo nome do arquivo em português. Página sem tradução aparece em português, com o aviso do próprio Starlight.
- A página inicial é própria (`site/src/components/Landing.astro`), com o visual do design system. O que muda com o projeto (hipóteses, linha do tempo) vem de `site/src/data/landing.json`.
- Endereço e caminho base vêm das variáveis `PRUMO_SITE_URL` e `PRUMO_SITE_BASE`, pra publicar no GitHub Pages ou em outro serviço sem mudar código.

Alternativas avaliadas: Hugo com Hextra (sem Node, mas templates numa linguagem a mais), VitePress (ecossistema Vue, diferente do React planejado pra interface) e Zensical (sucessor do Material, ainda novo).

## Consequências
- O repositório passa a ter Node, só dentro de `site/`. A interface futura do Prumo também será TypeScript e React, então não é uma dependência a mais no longo prazo.
- Trocar de gerador é trocar a moldura: o conteúdo é Markdown comum em `docs/`.
- A busca (Pagefind) roda no próprio site, sem serviço externo.
- Publicar depende de um serviço de hospedagem. O caminho livre é o GitHub Pages (ADR-0005); enquanto o Actions da conta estiver parado, ver a issue de hospedagem.

## Fontes
- Starlight · https://starlight.astro.build
- Material for MkDocs, aviso de fim da vida · https://github.com/squidfunk/mkdocs-material/issues/8523
- Pagefind · https://pagefind.app
