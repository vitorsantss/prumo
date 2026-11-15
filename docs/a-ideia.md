# A ideia em 5 minutos

## O problema

Pedir código a uma IA é rápido. Saber se ele resolve o problema certo, do jeito certo, é difícil. Um pedido mal entendido vira uma especificação errada, que vira um plano errado, que vira muito código errado; e, sem etapas, isso só aparece no fim, quando corrigir custa mais[^1].

## A hipótese

Cada tarefa percorre **etapas**, e cada etapa entrega **um arquivo** que a próxima lê. Entre uma etapa e outra existe um **portão**: uma pessoa aprova, pede ajuste ou manda refazer.

| Etapa | Lê | Escreve | Pra quê |
| --- | --- | --- | --- |
| Research | a tarefa | `research.md` | entender o pedido, o sistema e o negócio antes de qualquer código |
| SPEC | `research.md` | `spec.md` | o que precisa ser feito, com critérios de aceite verificáveis |
| PLAN | `spec.md` | `plan.md` | como fazer, arquivo por arquivo |
| Código | `plan.md` | o diff | a implementação, numa cópia do repositório só da tarefa |
| Review | o diff | a decisão | reprovou? volta pra etapa onde o erro nasceu (ADR-0002) |

### Por que começar pelo Research

O Research nasceu do jeito de trabalhar de quem mantém o Prumo: antes de escrever código, entender o pedido, o sistema e o negócio. Depois fomos ver se a literatura sustentava essa prática, e sustenta. No design, a fase de **descoberta** vem antes da entrega: o Double Diamond, do Design Council, separa explorar o problema de construir a solução[^2]. Em produto, Marty Cagan chama isso de *dual-track*: descobrir o que vale construir corre em paralelo a entregar[^3]. No trabalho com agentes de código, o mesmo aparece no fluxo **Research → Plan → Implement**, porque um erro na pesquisa se multiplica nas etapas seguintes[^4], e no guia de boas práticas do Claude Code, que recomenda explorar e planejar antes de codar[^5]. O Prumo trata essa pesquisa como uma etapa com arquivo próprio e portão.

## O portão

A IA faz o trabalho; o portão é humano. As três decisões possíveis (aprovar, pedir ajuste, refazer) ficam registradas no rastro da etapa, junto com o que ela leu, escreveu, gastou e quanto tempo levou (ADR-0007).

## Um padrão que se adapta a você

Ninguém sabe ainda qual é o jeito certo de construir software com IA: surgem padrões novos o tempo todo, como o desenvolvimento guiado por especificação[^6], e quase ninguém mede. O Prumo quer ser o lugar onde dá pra **montar fluxos diferentes, rodar a mesma tarefa em cada um e medir** qual funciona pro jeito de cada pessoa trabalhar.

## O que o Prumo é, e o que não é

- **É** uma hipótese em teste, aberta, com cada decisão registrada.
- **Não é** uma ferramenta pronta. Veja [o estado atual](estado-atual.md).

## Fontes

[^1]: Barry Boehm, *Software Engineering Economics* (1981): o custo de corrigir um defeito cresce com a fase em que é encontrado · https://openlibrary.org/books/OL4267874M
[^2]: Design Council, *The Double Diamond* · https://www.designcouncil.org.uk/our-resources/the-double-diamond/
[^3]: Marty Cagan, *Dual-Track Agile* (2012) · https://www.svpg.com/dual-track-agile/
[^4]: Dex Horthy, *Advanced Context Engineering for Coding Agents* · https://humanlayer.dev/blog/advanced-context-engineering
[^5]: Anthropic, *Claude Code: Best practices for agentic coding* · https://www.anthropic.com/engineering/claude-code-best-practices
[^6]: GitHub, *Spec Kit* · https://github.com/github/spec-kit
