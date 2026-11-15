# Current state

> **Translated by AI, waiting for a human review.**

> **In study.** Prumo does not work as a tool yet. This page says what exists, what is missing and how we will know whether the idea is worth it.

## What exists

A command-line tool in Go that takes a sample task through two stages, specification and delivery, with a gate in between. It records what each stage read and wrote, the time, the tokens and the gate decisions. By default it runs without AI, with simulated answers, for study (ADR-0006, ADR-0007).

## What does not exist yet

- The research, plan, code and review stages.
- Use on real tasks and any measure of quality or cost.
- An interface beyond the terminal.

## What we want to prove

Each idea becomes a hypothesis with a measure. None has been tested yet.

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
| Nov 14, 2026 | Born open: name, free license (ADR-0004) and the first recorded decisions. |
| Nov 15, 2026 | First skeleton running, without AI, for study. |
| now | Closing the foundational decisions and documenting the process. |
| next | Use it on real tasks, measure cost and quality, and only then say whether it works. |
