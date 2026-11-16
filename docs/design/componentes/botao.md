# Botão

**Produto:** os dois · **Estado:** estável

Dispara uma ação ou leva a outra página. Três pesos, pra deixar claro o que é principal.

## Prévia
<div class="previa not-content">
<div class="p-tema-claro p-canvas"><span class="previa-rotulo">Claro</span><span class="previa-linha"><a class="p-botao primario" href="#previa">Começar pela documentação</a></span><span class="previa-linha"><a class="p-botao" href="#previa">Compartilhar a ideia</a><button class="p-botao-icone" type="button" aria-label="Buscar"><svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.9" stroke-linecap="round" aria-hidden="true"><circle cx="11" cy="11" r="7"></circle><path d="M20 20l-3.5-3.5"></path></svg></button></span></div>
<div class="p-tema-escuro p-canvas"><span class="previa-rotulo">Escuro</span><span class="previa-linha"><a class="p-botao primario" href="#previa">Começar pela documentação</a></span><span class="previa-linha"><a class="p-botao" href="#previa">Compartilhar a ideia</a><button class="p-botao-icone" type="button" aria-label="Buscar"><svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.9" stroke-linecap="round" aria-hidden="true"><circle cx="11" cy="11" r="7"></circle><path d="M20 20l-3.5-3.5"></path></svg></button></span></div>
</div>

## Anatomia
1. **Primário** (`.p-botao.primario`): fundo na cor do texto. No máximo um por contexto.
2. **Secundário** (`.p-botao`): só contorno.
3. **Terciário**: texto sublinhado, sem caixa. Também serve de aba.
4. **Só ícone** (`.p-botao-icone`): quadrado de 44 px, sempre com `aria-label`.

## Estados
| Estado | Como fica |
| --- | --- |
| Normal | Como na prévia |
| Foco | Contorno visível do navegador; nunca remova sem pôr outro |
| Desligado | Não usamos: se a ação não pode acontecer, explique por quê no lugar do botão |

## Em tela pequena
Botões de texto ocupam a largura toda e empilham. O alvo continua com 44 px de altura.

## Acessibilidade
- `<a href>` quando leva a outra página, `<button>` quando faz algo na mesma página.
- Botão só de ícone sempre com `aria-label`.
- Altura mínima de 44 px (veja [Acessibilidade](../fundamentos/acessibilidade.md)).

## Tokens que usa
| Token | Onde |
| --- | --- |
| `--p-ink` | Fundo do primário, texto do secundário |
| `--p-bg` | Texto do primário |
| `--p-line-soft` | Contorno do secundário e do só ícone |
| `--p-raio-btn` | Cantos |
| `--p-fonte-texto` | Texto |

## Código
<!-- codigo:site/src/estilos/tokens.css#p-botao,p-botao-icone -->

## Implementações
| Onde | Situação |
| --- | --- |
| HTML e CSS (referência) | Classes globais em `site/src/estilos/tokens.css` |
| React, Angular, outros | Ainda não existe. Veja [como propor](../como-propor-um-componente.md) |
