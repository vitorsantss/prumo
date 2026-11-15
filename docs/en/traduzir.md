# Help translate

> **Translated by AI, waiting for a human review.**

Why some pages have no English yet, how to check each page, and how to pick a task.

## Why not everything is in English

Prumo was born in Portuguese. The entry pages (start here, the idea, current state, ways to contribute and first contribution) were translated by AI so newcomers can understand the idea. The decisions (ADRs) were not, because they change more often and need careful review: a term translated wrong in a decision changes what it says.

## How to check a page

In English, a page without translation shows a yellow **PT** sticker in the sidebar and a notice at the top, and its content stays in Portuguese. The language selector at the top of each page shows one of these states for English:

| State | What it means |
| --- | --- |
| reviewed translation | A person read and fixed it. |
| translated by AI, awaiting review | Readable, but nobody checked it yet. |
| not translated yet | The page shows the PT sticker and its content in Portuguese. |

## How to help

- **Review an AI translation.** About twenty minutes, no coding needed: read `docs/en/<page>.md` next to `docs/<page>.md`, fix what sounds off and remove the line "Translated by AI, waiting for a human review." at the top. That changes the state to "reviewed translation".
- **Translate a page or an ADR.** One task per file: comment on the task which one you will take, create `docs/en/<same-name>.md` (or `docs/en/adr/<same-name>.md`) and open a PR. If you use AI, start the page with the same notice line so someone else reviews it.

Open tasks have the [`traducao`](https://github.com/vitorsantss/prumo/issues?q=is%3Aopen+label%3Atraducao) label. The step by step for a first PR is in [First contribution](primeira-contribuicao.md).
