# Como mudar o site e a documentação

Qualquer pessoa pode propor mudança na página inicial, na documentação ou no visual. O caminho é o mesmo que o Prumo defende pra software: a ideia é escrita e aprovada antes do código.

## Só texto da documentação
Corrigiu um erro, deixou algo mais claro, acrescentou uma fonte? O rodapé de cada página diz de qual arquivo ela vem (por exemplo, `docs/a-ideia.md`). Abra esse arquivo [no repositório](https://github.com/vitorsantss/prumo/tree/main/docs), edite pelo próprio GitHub e abra o PR direto, sem issue.

Prefere só avisar? [Abra uma issue](https://github.com/vitorsantss/prumo/issues/new?template=site.yml) dizendo a página e o que está errado.

## Página nova
Todo arquivo em `docs/` vira página do site, e toda página precisa de um lugar no menu (`site/sidebar.mjs`). Se ela ainda não está pronta pra sair, ponha o caminho em `naoPublicado`, em `site/scripts/sincronizar-docs.mjs`. Aí ela fica só no repositório, e os links pra ela levam pro GitHub. O build falha se aparecer uma página publicada fora do menu, porque a busca acharia uma página que ninguém decidiu publicar.

## Como escrever

- **Toda citação tem link.** Afirmação sobre o mundo leva a fonte, de preferência como nota de rodapé (`texto[^1]` e, no fim, `[^1]: Autor, *Título* · URL`). Projeto, ferramenta ou desafio citado pelo nome leva link pra página dele.
- **Que link usar.** Página oficial quando existir (projeto, especificação, documentação). Artigo ou norma: o DOI (`https://doi.org/...`). Livro: o registro no [Open Library](https://openlibrary.org), que mostra ISBN e onde encontrar. Prefira a fonte primária ao texto que comenta ela.
- **Fonte sem link: como saber se é confiável.**
  1. Procure um identificador: DOI, ISBN, número do arXiv ou da norma. Achou, use como link.
  2. Veja quem publicou: autor com trabalho reconhecido na área, editora, periódico com revisão por pares, órgão oficial ou o próprio dono do projeto.
  3. Confira se ainda vale: página marcada como "não mantida", norma inativa ou versão antiga mudam o que dá pra afirmar.
  4. Leia o trecho: a fonte diz mesmo o que estamos afirmando? Resumo de terceiro não conta.
  5. Não passou? Não cite. Se é experiência nossa, escreva como experiência nossa, sem atribuir a ninguém.
- **Citação interna também.** Escreva `ADR-0007` ou `#50` normalmente: o site transforma em link pro ADR e pra issue. Se o ADR ainda não foi mergeado, o build avisa e o link aparece sozinho depois do merge.
- **Português do Brasil, frases curtas, um assunto por parágrafo.**

## Conteúdo da página inicial
Hipóteses, linha do tempo e textos da página inicial ficam em `site/src/data/landing.json`, em português e inglês. Mudou o estado do projeto (uma hipótese foi testada, um marco foi cumprido)? Atualize esse arquivo no mesmo PR da mudança. Conferir isso automaticamente é a issue #50.

## Visual, componente ou navegação
1. **Abra uma issue** com o modelo ["Mudança no site ou na documentação"](https://github.com/vitorsantss/prumo/issues/new?template=site.yml): o problema que você viu, a proposta e quem é afetado.
2. **Mostre como fica**, nos temas claro e escuro. Vale print, desenho no Excalidraw ou protótipo; o importante é dar pra comparar antes e depois.
3. **Confira as regras** da [camada Site do design system](https://github.com/vitorsantss/prumo/blob/main/docs/design/site.md). Se a proposta quebra uma regra, diga qual e por quê: às vezes a regra é que precisa mudar.
4. **Espere a aprovação na issue** antes de programar. É o portão humano: evita PR grande recusado por causa do desenho.
5. **Abra o PR** com:
   - o código;
   - `docs/design/site.md` atualizado, se surgiu regra ou componente novo;
   - um ADR, se entrou dependência ou padrão novo;
   - prints nos temas claro e escuro, em celular, notebook e tela larga;
   - `cd site && npm run build` passando;
   - `npm run auditar` sem nota abaixo de 90 (ver abaixo).
6. **Revisão e merge por squash**, como todo PR (ADR-0008).

## Testar no seu computador
```
cd site
npm install
npm run dev
```
A busca só funciona depois de `npm run build && npm run preview`, porque o índice é gerado no build.

## Desempenho, acessibilidade e SEO
O site é medido com o Lighthouse, a auditoria do Chrome[^lighthouse]: desempenho, acessibilidade, boas práticas, SEO e navegação por agentes de IA. Depois do build:
```
cd site
npm run build
npm run auditar
```
O comando sobe o `astro preview`, audita a página inicial e cinco páginas da documentação no modo celular e no modo computador, e falha se alguma nota ficar abaixo de 90. Precisa do Chrome instalado; se ele não for achado, aponte o caminho em `CHROME_PATH`. Pra medir o site publicado, passe o endereço: `npm run auditar -- https://prumo-docs.pages.dev/`.

Também dá pra rodar à mão no Chrome: abra a página, `F12`, aba **Lighthouse**, **Analyze page load**.

[^lighthouse]: Chrome for Developers, *Introduction to Lighthouse* · https://developer.chrome.com/docs/lighthouse/overview

