# Estado atual

> **Em estudo.** O Prumo ainda não funciona como ferramenta. Esta página diz o que já existe, o que falta e como vamos saber se a ideia vale a pena.

## O que já existe

Uma linha de comando em Go que leva uma tarefa de exemplo por duas etapas, especificação e entrega, com um portão no meio. Ela grava o que cada etapa leu e escreveu, o tempo, os tokens e as decisões do portão. Por padrão roda sem IA, com respostas simuladas, pra estudo ([ADR-0006](adr/0006-execucao-local-first-e-stack.md), [ADR-0007](adr/0007-modelo-de-execucao.md)).

## O que ainda não existe

- As etapas de pesquisa, plano, código e revisão.
- Uso em tarefas reais e qualquer medida de qualidade ou custo.
- Uma interface além do terminal.

## O que queremos provar

Cada ideia vira uma hipótese com uma medida. Nenhuma foi testada ainda.

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
| agora | Fechando as decisões de base e documentando o processo. |
| depois | Usar em tarefas reais, medir custo e qualidade e só então dizer se funciona. |
