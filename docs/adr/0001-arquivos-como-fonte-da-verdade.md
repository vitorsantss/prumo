# 0001 · Arquivos em tasks/<id>/ como fonte da verdade

- **Status:** Aceito
- **Data:** 2026-11-14

## Contexto

Cada tarefa que passa pelo Prumo gera coisas: a tarefa de origem, a especificação, o plano, as decisões dos gates, o consumo de tokens. Precisamos decidir onde esse estado mora de verdade.

O Prumo roda na máquina de quem usa (ADR-0006), quer ser estudado e auditado por qualquer pessoa e não pode depender de serviço pago (ADR-0005).

## Decisão

O estado de cada tarefa mora em **arquivos de texto** numa pasta própria:

```
tasks/<id>/
  tarefa.md      # a tarefa de origem, como foi recebida
  spec.md        # saída da etapa SPEC
  plan.md        # saída da etapa PLAN
  gates.jsonl    # uma linha por decisão de gate (aprovar, pedir ajuste, refazer)
  usage.jsonl    # uma linha por chamada de IA (etapa, modelo, tokens, custo)
  entrega.md     # resumo final
```

Qualquer banco ou índice que vier depois (pra tela, busca ou relatórios) é **derivado** desses arquivos e pode ser apagado e reconstruído a partir deles.

## Alternativas consideradas

- **Banco de dados como fonte principal:** consultas mais fáceis desde o início, mas exige instalar e manter um banco, esconde o histórico de quem estuda e prende o Prumo a um formato que não se lê sem ferramenta.

## Consequências

- Dá pra ler, versionar com git, comparar e revisar tudo com ferramentas comuns. Quem estuda vê exatamente o que cada etapa produziu.
- O C1 não precisa de banco.
- Escrita concorrente na mesma tarefa precisa de cuidado (uma etapa por vez por tarefa no C1).
- Mudança no formato dos arquivos é mudança de contrato: exige ADR e migração.
- Dados sensíveis não podem cair nesses arquivos (ADR-0003).

## Fontes

- Local-first software, Kleppmann et al. (Ink & Switch, 2019) · https://www.inkandswitch.com/local-first/
- Hunt e Thomas, *The Pragmatic Programmer*, tópico "The Power of Plain Text"
- JSON Lines · https://jsonlines.org
