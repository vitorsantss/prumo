# Como o site é publicado

O site fica em **https://prumo-docs.pages.dev**, no Cloudflare Pages, no plano grátis. Cada merge na `main` publica uma versão nova sozinho. O Cloudflare é uma ponte: o ADR-0010 aponta o GitHub Pages como destino, e ele fica parado enquanto o GitHub Actions da conta não volta (#41). A volta está registrada em #61.

## O caminho de uma mudança até o site

1. A mudança entra por PR, passa pela revisão e vira um commit na `main` (merge por squash, ADR-0008).
2. O Cloudflare percebe o push na `main`, copia o repositório e roda `npm run build` dentro de `site/`. Esse comando sincroniza `docs/`, gera as páginas e monta o índice da busca.
3. Se o build passa, a pasta `dist/` vira a versão no ar. Se falha, nada é publicado e o log fica no painel do Cloudflare.

Branches que não são a `main` ganham uma prévia automática, com endereço próprio[^previas]. PR aberto a partir de fork não ganha prévia[^github].

## Configuração atual

| Campo | Valor |
| --- | --- |
| Projeto | `prumo-docs` (o nome `prumo` já era de outra pessoa) |
| Endereço | https://prumo-docs.pages.dev |
| Repositório | `vitorsantss/prumo`; o app do Cloudflare no GitHub só enxerga este repositório |
| Branch de produção | `main` |
| Framework preset | None |
| Pasta raiz | `site` |
| Comando de build | `npm run build` |
| Pasta publicada | `dist` |
| Variáveis (produção e prévia) | `PRUMO_SITE_URL=https://prumo-docs.pages.dev`, `PRUMO_SITE_BASE=/`, `NODE_VERSION=22` |
| Plano | Free, sem cartão cadastrado |

Nenhuma dessas variáveis é segredo: o site é estático e elas só dizem o endereço. Não existe token nem senha guardados no projeto.

## Limites do plano grátis

| Limite | Valor |
| --- | --- |
| Builds por mês | 500 |
| Builds ao mesmo tempo | 1 |
| Tempo máximo de um build | 20 minutos |
| Arquivos por site | 20.000 |
| Tamanho de um arquivo | 25 MiB |
| Domínios próprios por projeto | 100 |

Fonte dos números: a página de limites do Cloudflare Pages[^limites]. Hoje o site tem por volta de cem arquivos e cada merge gera um build. A conta só aperta se tiver dezenas de pushes por dia; nesse caso, dá pra desligar as prévias automáticas de branch.

## Desfazer uma publicação

No painel, em **Workers & Pages → prumo-docs → Deployments**, abra o menu de três pontos da versão desejada e escolha **Rollback to this deployment**[^rollback]. Só versões de produção que passaram no build servem de destino. A correção de verdade continua sendo um PR que conserta a `main`.

## Domínio

O endereço `prumo-docs.pages.dev` é grátis e já vem com HTTPS. Domínio próprio é opcional e o registro do domínio é pago à parte. Se um dia fizer sentido:

1. Em **Custom domains → Set up a custom domain**, informe o domínio.
2. No DNS, crie um `CNAME` apontando pra `prumo-docs.pages.dev`. Se o DNS do domínio estiver no Cloudflare, o registro é criado sozinho[^dominios].
3. Troque a variável `PRUMO_SITE_URL` pro domínio novo e publique de novo.
4. Opcional: redirecionar o `pages.dev` pro domínio, com o recurso Bulk Redirects da conta[^redirecionar].

## Voltar pro GitHub Pages

Vale avaliar a volta (#61) quando o Actions estiver funcionando e pelo menos um destes acontecer: o Cloudflare mudar as condições do plano grátis, a ponte começar a dar trabalho, ou o projeto quiser reduzir serviços de terceiros, como pede o ADR-0005.

| | Cloudflare Pages (hoje) | GitHub Pages |
| --- | --- | --- |
| Custo | Grátis | Grátis pra repositório público |
| Depende do Actions | Não | Sim |
| Prévia por branch | Automática | Precisa montar |
| Limites principais | 500 builds/mês | site até 1 GB, 100 GB de banda por mês (limite flexível), deploy em até 10 minutos[^limites-gh] |
| Endereço | `prumo-docs.pages.dev` | `vitorsantss.github.io/prumo` |

Passos da migração:

1. Criar `.github/workflows/deploy.yml` com a action oficial do Astro, `withastro/action`, que é a forma recomendada pelo Astro[^astro], apontando o `path` pra `site`.
2. Em **Settings → Pages**, escolher **GitHub Actions** como origem.
3. Não definir `PRUMO_SITE_URL` nem `PRUMO_SITE_BASE`: os valores padrão do `astro.config.mjs` já são `https://vitorsantss.github.io` e `/prumo`.
4. Conferir o site novo nas três larguras (celular, notebook, tela larga), com busca e troca de idioma.
5. Atualizar o endereço no README, no ADR-0010 e nesta página.
6. Apagar o projeto `prumo-docs` no Cloudflare e remover o app do Cloudflare das integrações do GitHub.

[^previas]: Cloudflare, *Branch deployment controls* · https://developers.cloudflare.com/pages/configuration/branch-build-controls/
[^github]: Cloudflare, *GitHub integration* · https://developers.cloudflare.com/pages/configuration/git-integration/github-integration/
[^limites]: Cloudflare, *Pages limits* · https://developers.cloudflare.com/pages/platform/limits/
[^rollback]: Cloudflare, *Rollbacks* · https://developers.cloudflare.com/pages/configuration/rollbacks/
[^dominios]: Cloudflare, *Custom domains* · https://developers.cloudflare.com/pages/configuration/custom-domains/
[^redirecionar]: Cloudflare, *Redirecting \*.pages.dev to a Custom Domain* · https://developers.cloudflare.com/pages/how-to/redirect-to-custom-domain/
[^limites-gh]: GitHub, *GitHub Pages limits* · https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits
[^astro]: Astro, *Deploy your Astro Site to GitHub Pages* · https://docs.astro.build/en/guides/deploy/github/
