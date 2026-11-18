# 0011 · Design system no repositório e no site

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
O design system nasceu num canvas do claude.ai e chegou ao repositório como cópia (`docs/design/`), com o aviso "não edite à mão". Enquanto isso, as peças do site foram nascendo direto no código, nos PRs, sem virar componente documentado. Resultado: o design system está partido em dois. O app só existe como documentação, o site só como código, e os dois usam nomes de token diferentes (`card-solid` num, `--p-solid` no outro).

O projeto quer que pessoas e IAs leiam as mesmas regras, que qualquer um proponha componente por PR ([ADR-0001](0001-arquivos-como-fonte-da-verdade.md): arquivo como fonte da verdade) e que nada dependa de ferramenta proprietária sem caminho nativo ([ADR-0005](0005-integracoes-terceirizadas-caminho-nativo.md)).

## Decisão
1. **O repositório é a fonte.** `docs/design/` passa a ser editado por PR, como o resto da documentação. O canvas do claude.ai vira rascunho: serve pra desenhar, mas o que vale é o que entra no repositório.
2. **Seção própria no site**, com quatro partes: Fundamentos (princípios, cor, tipografia, espaço, movimento, acessibilidade), Componentes, Padrões e Implementações.
3. **Componentes separados pelo produto, não pelo aparelho:** App (o canvas do Prumo), Site (página inicial e documentação) e Os dois. Tela pequena é uma seção dentro de cada componente.
4. **Molde fixo por componente:** pra que serve, prévia clara e escura, anatomia, estados, em tela pequena, acessibilidade, tokens que usa, código e implementações. O mesmo molde serve pra quem lê no site e pra IA que lê o Markdown.
5. **Tokens no formato Design Tokens 2025.10**, do Design Tokens Community Group[^dtcg]. É a primeira versão estável, mas não é norma do W3C nem está no caminho de virar uma, e ainda recebe refinamentos[^dtcg-faq]; o Prumo acompanha as versões novas como qualquer dependência. As variáveis CSS do site são geradas a partir desse arquivo, com um nome só por token.
6. **Implementação de referência em HTML e CSS.** Os componentes do site são arquivos `.astro` ([ADR-0010](0010-site-com-starlight.md)); a doc mostra o que o navegador recebe (HTML, CSS e custom elements nativos, como `<prumo-idioma>`), extraído do arquivo no build. Ninguém copia código pra doc à mão.
7. **Trechos de código dos componentes em CC0.** O código mostrado nas páginas de componente pode ser copiado sem obrigação, seguindo a recomendação da Free Software Foundation pra trechos pequenos em documentação[^fsf]. O resto do repositório continua AGPL-3.0-only ([ADR-0004](0004-licenca-agpl-3.md)).
8. **Miniaturas de link fazem parte do design system:** imagem 1200 × 630, o tamanho que a Meta recomenda pra telas de alta resolução[^meta], gerada no build pra cada página, composição centrada (sobrevive ao corte quadrado), tema claro e texto alternativo em `og:image:alt`, que o protocolo Open Graph pede sempre que houver `og:image`[^ogp].

Port pra outros frameworks (React, Angular…) segue o molde e entra em Implementações. A escolha entre um port por framework e Web Components como núcleo fica pra quando o app tiver interface web.

## Consequências
- O aviso "não edite à mão" sai de `docs/design/README.md`.
- O v0 migra **todos** os componentes que já existem, do site e do app, pro molde. A lista está na [#65](https://github.com/vitorsantss/prumo/issues/65).
- Mudança visual ou componente novo só entra com a página do componente atualizada (regra no `AGENTS.md` e no modelo de PR, junto com o v0).
- Renomear tokens mexe no CSS do site inteiro de uma vez; vai num PR só, conferido nos dois temas e nas três larguras.
- Quem copia um trecho de componente não herda a AGPL; quem copia um arquivo inteiro do repositório, sim.

## Fontes
[^dtcg]: Design Tokens Community Group, *Design Tokens Format Module 2025.10* · https://www.designtokens.org/tr/2025.10/
[^dtcg-faq]: Design Tokens Community Group, *FAQ* (estado da especificação) · https://www.designtokens.org/faq/
[^fsf]: Free Software Foundation, *How to Choose a License for Your Own Work* · https://www.gnu.org/licenses/license-recommendations
[^ogp]: *The Open Graph protocol* · https://ogp.me
[^meta]: Meta for Developers, *Images in Link Shares* · https://developers.facebook.com/docs/sharing/webmasters/images
