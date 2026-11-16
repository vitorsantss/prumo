# Design system

As regras visuais do Prumo, num lugar só. Pessoas leem aqui; uma IA que clonar o repositório lê os mesmos arquivos em `docs/design/`. **Esta pasta é a fonte** (ADR-0011): mudanças entram por PR, como o resto da documentação, e o CSS do site é gerado a partir dela.

## Quatro partes
| Parte | O que tem |
| --- | --- |
| [Fundamentos](fundamentos/principios.md) | Princípios, cor, tipografia, espaço e raio, movimento, acessibilidade. Tudo sai de [`tokens.json`](tokens.json). |
| Componentes | Cada peça no mesmo molde: pra que serve, prévia clara e escura, anatomia, estados, tela pequena, acessibilidade, tokens e código. |
| Padrões | Componentes combinados pra resolver uma situação. Chegam na próxima entrega. |
| Implementações | HTML e CSS de referência, no próprio site. Outros frameworks entram seguindo o molde. |

## Onde cada componente vive
Componentes se separam pelo produto, não pelo aparelho. Tela pequena é uma seção dentro de cada componente.

| Produto | O que é | Componentes nesta entrega |
| --- | --- | --- |
| Site | Página inicial e documentação | [Seletor de idioma](componentes/seletor-de-idioma.md) |
| App | O canvas do Prumo | Na próxima entrega (hoje em [`sistema/`](https://github.com/vitorsantss/prumo/tree/main/docs/design/sistema)) |
| Os dois | Mesma regra, dois lugares | [Botão](componentes/botao.md), [Ponto de status](componentes/ponto-de-status.md) |

O que ainda falta migrar está na #65.

## Se você é uma IA
Comece por [`tokens.json`](tokens.json), no formato Design Tokens 2025.10 do W3C[^dtcg], e pelas [regras da camada Site](site.md). Cada componente tem uma página com o mesmo molde em `componentes/`. Use os tokens (`var(--p-…)`), nunca valores soltos.

## Arquivos
| Arquivo | O que tem |
| --- | --- |
| `tokens.json` | Fonte dos tokens. O site gera `site/src/estilos/tokens-gerados.css` a partir dele no build. |
| `site.md` | Regras da camada Site (página inicial e documentação). |
| `fundamentos/`, `componentes/` | As páginas desta seção. |
| `sistema/` | Versão anterior, vinda do canvas do claude.ai: componentes do app com prévia em HTML. Migra pro molde na próxima entrega. |

[^dtcg]: Design Tokens Community Group, *Design Tokens Format Module 2025.10* · https://www.designtokens.org/tr/2025.10/
