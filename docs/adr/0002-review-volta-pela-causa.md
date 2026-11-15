# 0002 · Review reprovado volta pela causa

- **Status:** Proposto
- **Data:** 2026-11-15

## Contexto
Quando o review reprova uma entrega, o caminho óbvio é mandar corrigir o código. Mas nem todo erro nasce no código: se a regra de negócio foi entendida errado, a SPEC está errada, o PLAN foi feito em cima dela, e consertar o código trata o sintoma. Quanto mais tarde se corrige um erro de origem, mais caro fica (Boehm), e com IA isso também vira tokens gastos à toa.

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

Mais tarde, o próprio agente de review pode **sugerir** a causa, e a pessoa confirma ou troca. A decisão continua humana.

## Consequências
- Cada reprovação gera um dado: com o tempo dá pra ver em que etapa os erros nascem e quanto custa cada volta (#40).
- O gate do Review precisa de uma opção a mais que o gate das outras etapas. Entra quando a etapa de Review existir (C2).
- Classificar errado custa uma volta a mais, mas não quebra nada: a etapa seguinte pode reprovar de novo com outra causa.

## Fontes
- Barry Boehm, *Software Engineering Economics* (1981), custo de corrigir defeitos por fase
- IEEE 1044-2009, *Standard Classification for Software Anomalies*, classificar a origem de um defeito
- Dex Horthy, *Advanced Context Engineering for Coding Agents*, erro na pesquisa ou no plano vira muitas linhas de código erradas · https://humanlayer.dev/blog/advanced-context-engineering
