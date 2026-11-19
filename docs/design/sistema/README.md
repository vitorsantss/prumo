# Prumo · design system

> Versão anterior, vinda do canvas do claude.ai. A fonte agora é [`docs/design/`](../README.md) ([ADR-0011](../../adr/0011-design-system-no-repositorio.md)); estes componentes do app migram pro molde na próxima entrega (#65).

Prumo é o fio com peso que o pedreiro usa pra ver se a parede está reta; no Nordeste, "tomar prumo" é entrar nos eixos. O sistema faz isso com o trabalho: cada tarefa passa pelas etapas no rumo certo. Canvas pessoal de fluxo de trabalho com IA: Jira → Research → SPEC → PLAN → Código → Review → Entrega.

## Princípios
1. **O canvas é a tela.** Nada de sidebar fixa nem barras no topo e na base. Painéis flutuam sobre o canvas pontilhado.
2. **Só pede atenção o que precisa de você.** Laranja (aguardando aprovação) é a única cor que chama ação. Botão primário aparece no máximo uma vez por contexto.
3. **Calmo, não neon.** Blur leve (6px), degradê só nas linhas entre etapas, sombras curtas. Sem cores saturadas de enfeite.
4. **Manuscrito só em destaque.** Architects Daughter no logo, títulos de etapa e nomes nos quadradinhos. Todo o resto em Inter; código e tokens em JetBrains Mono.
5. **Forma + cor.** Status muda de forma (vazado, anel, cheio com anel, cheio), não só de cor.
6. **Sempre claro e escuro.** Todo componente é desenhado e revisado nos dois temas.

## Superfícies
- **Card:** fundo translúcido + blur 6px + textura hachurada a 135°, cantos levemente irregulares (14 12 16 12 / 12 16 12 14).
- **Botões:** primário preenchido com a cor do texto; secundário só contorno; terciário é texto sublinhado (também usado como aba).
- **Tokens:** sempre no canto inferior direito. Legenda de status no centro inferior. Navegação no canto inferior esquerdo.

## Movimento
pop 240ms (popovers e cards), sobe 420ms (sessão), troca de cor de linha 600ms. Tudo desligado com prefers-reduced-motion.

## O que evitar
Neon, roxo em botão, barras de pílula ou segmentadas, scrollbar padrão do sistema (use a fina, com tokens), fonte manuscrita em texto corrido, traço "desenhado à mão" exagerado.

Protótipo de referência: canvas "Prumo · simulação".
