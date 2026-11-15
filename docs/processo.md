# Como trabalhamos

O Prumo é construído do jeito que ele mesmo propõe: em etapas pequenas, com decisões registradas e revisão humana antes de seguir.

## Checkpoints

O trabalho anda em checkpoints com data fixa (milestones no GitHub). Tempo fixo, escopo variável: se o prazo aperta, corta-se escopo, não se estica a data. Todo checkpoint termina com algo que roda ou com decisões aceitas, nunca só com mais desenho.

| Checkpoint | Até | Pronto quando |
| --- | --- | --- |
| C0 · Decidir | 19/11/2026 | As decisões que destravam o código são ADRs aceitos |
| C1 · Esqueleto andando | 03/12/2026 | Uma CLI roda 3 etapas numa tarefa real, com 1 gate manual |
| C2 · Primeiro valor | 17/12/2026 | Research + PLAN + risk gate usados em 3 tarefas reais |

Ideia nova não entra no checkpoint em andamento: vira issue com a label `ideia` e é avaliada no próximo.

Referência: Ryan Singer, *Shape Up*, cap. "Set Boundaries" · https://basecamp.com/shapeup/1.2-chapter-03

## Issues

- Cada checkpoint tem uma **issue-mãe** com as tarefas como **sub-issues**.
- Quando uma tarefa depende de outra, ela fica marcada como **bloqueada por** (`is:blocked` mostra todas). Não comece uma tarefa bloqueada.
- Toda tarefa diz **o que fazer** e **pronto quando**, com critérios verificáveis.
- Labels de área (`area:*`) dizem de que parte do software a tarefa é; `caminho-nativo` marca a troca de algo pago ou proprietário por algo livre (ADR-0005).

## Decisões (ADR)

Decisão que muda arquitetura, dependência, formato de arquivo ou forma de trabalhar vira um ADR em `docs/adr/`, no formato de Michael Nygard (contexto, decisão, consequências), sempre com fontes. Um ADR começa como **Proposto** e só vale quando muda pra **Aceito**. Decisão revista não é apagada: o ADR novo marca o antigo como substituído.

Referência: Michael Nygard, *Documenting Architecture Decisions* (2011) · https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions

## Revisão

- Nada entra na `main` sem PR e sem aprovação de um mantenedor.
- No Review, quem reprova classifica a causa e diz pra onde a tarefa volta: regra de negócio → SPEC, implementação → PLAN, ajuste visual → fluxo curto (ADR-0002). Por enquanto, quem revisa é o mantenedor.
- Todo PR diz se usou IA e qual. Quem envia responde por cada linha.

## Onde as coisas moram

| O quê | Onde |
| --- | --- |
| Código, ADRs e documentação | Este repositório |
| Tarefas e checkpoints | Issues e milestones |
| Perguntas e conversas abertas | Discussions |
