# Cor

Destaque é a cor do texto (`--p-ink`). Cor só aparece quando diz um estado, e o estado muda de forma também.

<div class="previa not-content">
<div class="p-tema-claro p-canvas"><span class="previa-rotulo">Claro</span><span class="previa-linha"><span class="p-ponto falta"></span> não iniciado</span><span class="previa-linha"><span class="p-ponto andamento"></span> em andamento</span><span class="previa-linha"><span class="p-ponto acao"></span> pede sua ação</span><span class="previa-linha"><span class="p-ponto feito"></span> feito</span></div>
<div class="p-tema-escuro p-canvas"><span class="previa-rotulo">Escuro</span><span class="previa-linha"><span class="p-ponto falta"></span> não iniciado</span><span class="previa-linha"><span class="p-ponto andamento"></span> em andamento</span><span class="previa-linha"><span class="p-ponto acao"></span> pede sua ação</span><span class="previa-linha"><span class="p-ponto feito"></span> feito</span></div>
</div>

## Tokens
Gerados a partir de [`tokens.json`](../tokens.json). Cada token tem um valor pro tema claro e outro pro escuro; o site troca sozinho.

<!-- tokens:cor -->

## Regras
- Texto sobre fundo precisa de contraste de pelo menos 4,5:1 (veja [Acessibilidade](acessibilidade.md)).
- `--p-status-wait` (laranja) é a única cor que pede ação. Não use pra decorar.
- O adesivo (`--p-sticky`) tem texto sempre `#2B2620`, nos dois temas.
