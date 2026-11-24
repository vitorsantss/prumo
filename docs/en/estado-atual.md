# Current state

> **Translated by AI, waiting for a human review.**

> **In study.** Prumo does not work as a tool yet. This page says what exists, what is missing and how we will know whether the idea is worth it.

## What exists

A command-line tool in Go that takes a sample task through two stages, specification and delivery, with a gate in between. It records what each stage read and wrote, the tokens per model, the agent turns, the AI time apart from the human time and the gate decisions ([ADR-0006](../adr/0006-execucao-local-first-e-stack.md), [ADR-0007](../adr/0007-modelo-de-execucao.md)). By default it runs without AI, with simulated answers; with Claude Code, it has already taken a real Rinha de Backend task end to end ([report of the runs](../../exemplos/rinha-2024q1/execucoes/)).

## What does not exist yet

- The research, plan, code and review stages.
- A measure of quality and cost across several real tasks.
- An interface beyond the terminal.

## What we want to prove

Each idea becomes a hypothesis with a measure. None has been tested yet: two runs on one task show that the path works, not that the idea is worth it.

| | Hypothesis | Measure | Status |
| --- | --- | --- | --- |
| H1 | Stages with declared files reduce rework. | How often the task goes back, per stage. | not tested yet |
| H2 | Cost per approved stage drops with a human gate. | Tokens and time until approval. | not tested yet |
| H3 | The best flow changes from person to person. | The same task in different flows. | not tested yet |

The first study case comes from a public challenge, the Brazilian [Rinha de Backend (2024/Q1 edition)](https://github.com/zanfranceschi/rinha-de-backend-2024-q1): a banking API that records transactions and shows a statement.

## Timeline

| When | Milestone |
| --- | --- |
| Nov 11, 2026 | The idea: a personal way of working with AI becomes the gated-stages hypothesis. |
| Nov 14, 2026 | Born open: name, free license ([ADR-0004](../adr/0004-licenca-agpl-3.md)) and the first recorded decisions. |
| Nov 15, 2026 | First skeleton running, without AI, for study. |
| Nov 19, 2026 | Foundational decisions accepted: checkpoint C0 closes with every ADR decided. |
| Nov 24, 2026 | First real task with AI, end to end: checkpoint C1 closes. |
| now | Three real tasks, with research, plan and a risk gate (checkpoint C2). |
| next | Use it on real tasks, measure cost and quality, and only then say whether it works. |
