# Ponto de status

**Produto:** os dois · **Estado:** estável

Mostra em que pé algo está. A forma muda junto com a cor, pra funcionar sem enxergar cor.

## Prévia
<div class="previa not-content">
<div class="p-tema-claro p-canvas"><span class="previa-rotulo">Claro</span><span class="previa-linha"><span class="p-ponto falta"></span> não iniciado</span><span class="previa-linha"><span class="p-ponto andamento"></span> em andamento</span><span class="previa-linha"><span class="p-ponto acao"></span> pede sua ação</span><span class="previa-linha"><span class="p-ponto feito"></span> feito</span></div>
<div class="p-tema-escuro p-canvas"><span class="previa-rotulo">Escuro</span><span class="previa-linha"><span class="p-ponto falta"></span> não iniciado</span><span class="previa-linha"><span class="p-ponto andamento"></span> em andamento</span><span class="previa-linha"><span class="p-ponto acao"></span> pede sua ação</span><span class="previa-linha"><span class="p-ponto feito"></span> feito</span></div>
</div>

## Anatomia
| Classe | Forma | Cor |
| --- | --- | --- |
| `.p-ponto.falta` | Vazado | `--p-status-todo` |
| `.p-ponto.andamento` | Anel com miolo | `--p-status-run` |
| `.p-ponto.acao` | Cheio com anel, único laranja da tela | `--p-status-wait` |
| `.p-ponto.feito` | Cheio | `--p-status-done` |

## Estados
São os quatro da anatomia. O app tem um quinto, bloqueado (`--p-status-block`), que entra na migração dos componentes do app.

## Em tela pequena
Igual. O ponto tem 12 px e nunca aparece sozinho: sempre com o texto do estado ao lado.

## Acessibilidade
- O ponto é decorativo (`aria-hidden="true"`); o estado vai em texto, visível ou em `.p-oculto`.
- A diferença de forma garante leitura pra quem não distingue cores (WCAG 1.4.1).

## Tokens que usa
`--p-status-todo`, `--p-status-run`, `--p-status-wait`, `--p-status-done`, `--p-bg` (anel do estado de ação).

## Código
<!-- codigo:site/src/estilos/tokens.css#p-ponto -->

## Implementações
| Onde | Situação |
| --- | --- |
| HTML e CSS (referência) | Classes globais em `site/src/estilos/tokens.css` |
| React, Angular, outros | Ainda não existe. Veja [como propor](../como-propor-um-componente.md) |
