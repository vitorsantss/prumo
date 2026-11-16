# Seletor de idioma

**Produto:** site · **Estado:** estável

Troca o idioma da página atual e mostra, antes do clique, o estado de cada tradução. Fica no cabeçalho, ao lado do botão de tema.

## Prévia
<div class="previa not-content">
<div class="p-tema-claro p-canvas"><span class="previa-rotulo">Claro</span><div class="p-ilha previa-painel"><span class="previa-titulo">Idioma desta página</span><span class="previa-opcao atual"><span><strong>Português</strong><br><small>original</small></span><span class="marca" aria-hidden="true"></span></span><span class="previa-opcao"><span><strong>English</strong><br><small>traduzida por IA</small></span></span></div></div>
<div class="p-tema-escuro p-canvas"><span class="previa-rotulo">Escuro, inglês sem tradução</span><div class="p-ilha previa-painel"><span class="previa-titulo">Idioma desta página</span><span class="previa-opcao atual"><span><strong>Português</strong><br><small>original</small></span><span class="marca" aria-hidden="true"></span></span><span class="previa-opcao falta"><span><strong>English</strong><br><small>ainda sem tradução</small><br><a href="#previa">Por que e como ajudar →</a></span></span></div></div>
</div>

## Anatomia
1. **Gatilho:** o idioma atual escrito no próprio idioma, sem bandeira (bandeira é país, não idioma).
2. **Título:** deixa claro que a escolha vale pra esta página.
3. **Opção atual:** fundo suave e ponto cheio.
4. **Opção sem tradução:** borda tracejada e link pra página que explica como ajudar a traduzir.

## Estados
| Estado | Como fica |
| --- | --- |
| Fechado | Só o gatilho |
| Aberto | Painel abaixo do gatilho; se faltar espaço, abre pra dentro da tela |
| Opção traduzida por IA | Texto "traduzida por IA" embaixo do nome |
| Opção sem tradução | Tracejado e link de ajuda, sem trocar de página |

## Em tela pequena
Abaixo de 50rem o painel vira uma folha que sobe da base, com alça e fundo escurecido. O gatilho mora no menu, junto do tema.

## Acessibilidade
- O gatilho é um `<button>` com `aria-expanded` e `aria-controls`.
- Cada opção leva `lang` e `hreflang` do próprio idioma.
- Fecha com Esc e com clique fora.
- Sem animação com `prefers-reduced-motion`.

## Tokens que usa
| Token | Onde |
| --- | --- |
| `--p-card-solid` | Fundo do painel |
| `--p-soft` | Opção atual e gatilho aberto |
| `--p-line-soft` | Bordas e tracejado da opção sem tradução |
| `--p-duracao-pop` | Entrada do painel |

## Código
<!-- codigo:site/src/components/SeletorIdioma.astro -->

## Implementações
| Onde | Situação |
| --- | --- |
| HTML e CSS (referência) | `site/src/components/SeletorIdioma.astro` |
| React, Angular, outros | Ainda não existe. Veja [como propor](../como-propor-um-componente.md) |
