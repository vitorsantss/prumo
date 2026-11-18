# 0002 · Review reprovado volta pela causa

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
Quando o review reprova uma entrega, o caminho óbvio é mandar corrigir o código. Mas nem todo erro nasce no código: se a regra de negócio foi entendida errado, a SPEC está errada, o PLAN foi feito em cima dela, e consertar o código trata o sintoma. A causa continua lá e reaparece na próxima tarefa que usar a mesma SPEC.

Três apoios pra voltar pela causa:
- **Classificar a causa melhora o processo.** A classificação ortogonal de defeitos (ODC), da IBM, usa o tipo de cada defeito pra dar retorno sobre o processo enquanto ele acontece[^odc]. A IEEE 1044 classifica anomalias pra análise causal e melhoria do processo[^ieee1044].
- **Com IA, o erro de origem se multiplica.** Uma linha ruim de plano pode virar centenas de linhas de código ruins, e revisar pesquisa e plano rende mais do que revisar código[^humanlayer]. É relato de um time praticante, não estudo controlado.
- **O que este ADR não usa como argumento:** a curva de Boehm, a ideia de que corrigir mais tarde custa muito mais. É crença comum, mas o maior estudo publicado sobre ela, com 171 projetos, não achou esse efeito de forma consistente[^menzies]. O Prumo pode testar isso nos próprios dados (ver Decisão e Consequências).

## Decisão
Ao reprovar, quem revisa diz **onde o erro nasceu**, e a tarefa volta pra etapa dessa causa:

| Causa | Exemplo (Rinha de Backend) | Volta pra |
| --- | --- | --- |
| Regra de negócio | A SPEC diz que o saldo não pode ficar negativo; a regra é que não pode passar do limite negativo do cliente | SPEC, e o que vem depois refaz a partir dela |
| Implementação | A SPEC está certa, mas com requisições simultâneas o saldo passa do limite | PLAN |
| Ajuste pequeno | Nome de campo, mensagem de erro, formatação | Fluxo curto: corrige e revisa de novo, sem refazer etapas |

**Quem classifica, por enquanto, é o mantenedor.** No gate do Review, a opção "reprovar" pergunta a causa:

```
[a] aprovar · [j] pedir ajuste · [r] reprovar > r
Causa: [n] regra de negócio · [i] implementação · [p] ajuste pequeno > i
O que está errado? > saldo passa do limite com requisições simultâneas
→ volta pro PLAN com o comentário
```

A causa e o comentário ficam gravados no `gates.jsonl` e no `etapa.json` da etapa (ADR-0007).

Toda reprovação, em qualquer gate, grava a etapa onde o erro nasceu e a etapa onde foi achado. O custo da correção vem do registro de uso (ADR-0014). Com esses três campos, o Prumo consegue repetir a medição de Menzies e colegas[^menzies] nos próprios dados.

Mais tarde, o próprio agente de review pode **sugerir** a causa, e a pessoa confirma ou troca. A decisão continua humana.

## Consequências
- Cada reprovação gera um dado: com o tempo dá pra ver em que etapa os erros nascem e quanto custa cada volta ([#40](https://github.com/vitorsantss/prumo/issues/40)).
- O Prumo vai poder testar se corrigir mais tarde custa mais nos próprios projetos, com o critério do estudo original: só conclui sobre um par de etapas com pelo menos 30 defeitos. Antes disso, o resultado é `sem_dado`. Mesmo depois, vale só pro Prumo.
- O gate do Review precisa de uma opção a mais que o gate das outras etapas. Entra quando a etapa de Review existir (C2).
- Classificar errado custa uma volta a mais, mas não quebra nada: a etapa seguinte pode reprovar de novo com outra causa.

## Fontes
[^odc]: Ram Chillarege et al., *Orthogonal Defect Classification: A Concept for In-Process Measurements*, IEEE TSE 18(11), 1992 · https://chillarege.com/docs/Papers/concept
[^ieee1044]: IEEE, *Std 1044-2009, Standard Classification for Software Anomalies*, hoje marcada como histórica nos catálogos de normas · https://webstore.ansi.org/standards/ieee/ieee10442009
[^humanlayer]: Dex Horthy (HumanLayer), *Advanced Context Engineering for Coding Agents*, 2025 · https://humanlayer.dev/blog/advanced-context-engineering
[^menzies]: Tim Menzies, William Nichols, Forrest Shull e Lucas Layman, *Are Delayed Issues Harder to Resolve? Revisiting Cost-to-Fix of Defects throughout the Lifecycle*, Empirical Software Engineering 22(4), 2017 · https://arxiv.org/abs/1609.04886
