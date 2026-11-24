# Estado atual

> **Em estudo.** O Prumo ainda não funciona como ferramenta. Esta página diz o que já existe, o que falta e como vamos saber se a ideia vale a pena.

## O que já existe

Uma linha de comando em Go que leva uma tarefa de exemplo por duas etapas, especificação e entrega, com um portão no meio. Ela grava o que cada etapa leu e escreveu, os tokens por modelo, as voltas do agente, o tempo da IA separado do tempo humano e as decisões do portão ([ADR-0006](adr/0006-execucao-local-first-e-stack.md), [ADR-0007](adr/0007-modelo-de-execucao.md)). Por padrão roda sem IA, com respostas simuladas; com o Claude Code, já levou uma tarefa real da Rinha de Backend de ponta a ponta ([relato das execuções](../exemplos/rinha-2024q1/execucoes/)).

## O que ainda não existe

- As etapas de pesquisa, plano, código e revisão.
- Medida de qualidade e de custo em várias tarefas reais.
- Uma interface além do terminal.

## O que queremos provar

Cada ideia vira uma hipótese com uma medida. Nenhuma foi testada ainda: duas execuções numa tarefa mostram que o caminho funciona, não que a ideia vale.

| | Hipótese | Medida | Situação |
| --- | --- | --- | --- |
| H1 | Etapas com arquivos declarados reduzem o retrabalho. | Quantas vezes a tarefa volta, por etapa. | ainda não testada |
| H2 | O custo por etapa aprovada cai com o portão humano. | Tokens e tempo até a aprovação. | ainda não testada |
| H3 | O melhor fluxo muda de pessoa pra pessoa. | A mesma tarefa em fluxos diferentes. | ainda não testada |

O primeiro caso de estudo vem de um desafio público, a [Rinha de Backend (edição 2024/Q1)](https://github.com/zanfranceschi/rinha-de-backend-2024-q1): uma API de banco que registra transações e mostra o extrato.

## Linha do tempo

| Quando | Marco |
| --- | --- |
| 11/11/2026 | A ideia: um jeito pessoal de trabalhar com IA vira a hipótese das etapas com portões. |
| 14/11/2026 | Nasce aberto: nome, licença livre ([ADR-0004](adr/0004-licenca-agpl-3.md)) e as primeiras decisões registradas. |
| 15/11/2026 | Primeiro esqueleto rodando, sem IA, pra estudo. |
| 19/11/2026 | Decisões de base aceitas: o checkpoint C0 fecha com todos os ADRs decididos. |
| 24/11/2026 | Primeira tarefa real com IA, de ponta a ponta: o checkpoint C1 fecha. |
| agora | Três tarefas reais, com pesquisa, plano e gate de risco (checkpoint C2). |
| depois | Usar em tarefas reais, medir custo e qualidade e só então dizer se funciona. |
