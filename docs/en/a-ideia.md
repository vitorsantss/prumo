# The idea in 5 minutes

> **Translated by AI, waiting for a human review.**

## The problem

Asking an AI for code is fast. Knowing whether it solves the right problem, the right way, is hard. A misunderstood request becomes a wrong specification, which becomes a wrong plan, which becomes a lot of wrong code; and without stages this only shows up at the end, when fixing it costs more[^1].

## The hypothesis

Each task goes through **stages**, and each stage delivers **one file** that the next one reads. Between stages there is a **gate**: a person approves, asks for changes or sends it back to be redone.

| Stage | Reads | Writes | What for |
| --- | --- | --- | --- |
| Research | the task | `research.md` | understand the request, the system and the business before any code |
| SPEC | `research.md` | `spec.md` | what needs to be done, with verifiable acceptance criteria |
| PLAN | `spec.md` | `plan.md` | how to do it, file by file |
| Code | `plan.md` | the diff | the implementation, in a copy of the repository just for the task |
| Review | the diff | the decision | rejected? back to the stage where the error was born (ADR-0002) |

### Why start with Research

Research comes from the way Prumo's maintainer works: before writing code, understand the request, the system and the business. Then we checked whether the literature backs this practice, and it does. In design, **discovery** comes before delivery: the Design Council's Double Diamond separates exploring the problem from building the solution[^2]. In product, Marty Cagan calls it *dual-track*: discovering what is worth building runs in parallel with delivery[^3]. With coding agents, the same idea shows up in the **Research → Plan → Implement** flow, because an error in research multiplies in the following stages[^4], and in the Claude Code best-practices guide, which recommends exploring and planning before coding[^5]. Prumo treats this research as a stage with its own file and gate.

## The gate

The AI does the work; the gate is human. The three possible decisions (approve, ask for changes, redo) are recorded in the stage trace, along with what it read, wrote, spent and how long it took (ADR-0007).

## A pattern that adapts to you

Nobody knows the right way to build software with AI yet: new patterns keep appearing, such as spec-driven development[^6], and almost nobody measures. Prumo wants to be the place where you can **build different flows, run the same task through each and measure** which one fits the way each person works.

## What Prumo is, and what it is not

- **It is** a hypothesis under test, open, with every decision recorded.
- **It is not** a finished tool. See [the current state](estado-atual.md).

## Sources

[^1]: Barry Boehm, *Software Engineering Economics* (1981): the cost of fixing a defect grows with the phase in which it is found · https://openlibrary.org/books/OL4267874M
[^2]: Design Council, *The Double Diamond* · https://www.designcouncil.org.uk/our-resources/the-double-diamond/
[^3]: Marty Cagan, *Dual-Track Agile* (2012) · https://www.svpg.com/dual-track-agile/
[^4]: Dex Horthy, *Advanced Context Engineering for Coding Agents* · https://humanlayer.dev/blog/advanced-context-engineering
[^5]: Anthropic, *Claude Code: Best practices for agentic coding* · https://www.anthropic.com/engineering/claude-code-best-practices
[^6]: GitHub, *Spec Kit* · https://github.com/github/spec-kit
