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

### Entregas em fila

Trabalho que não é do produto (site, qualidade, automação, métricas) anda em **entregas**, também como milestones, numa fila com ordem. A entrega ganha prazo quando começa, e então segue a mesma regra dos checkpoints: tempo fixo, escopo variável.

**Uma entrega em andamento por vez,** além do checkpoint do produto. A próxima só começa quando a atual fecha. O motivo é matemático: pela Lei de Little, L = λW, o trabalho em andamento (L) é a vazão (λ) vezes o tempo de ciclo (W); então o tempo de ciclo é o trabalho em andamento dividido pela vazão[^little]. Com a mesma vazão, cada frente aberta a mais atrasa todas. A pesquisa do DORA recomenda o mesmo na prática: limitar o que está em andamento e focar em terminar poucas tarefas de alta prioridade[^dora-wip].

Dentro de cada entrega, os PRs continuam pequenos: trabalhar em lotes pequenos permite testar rápido se uma mudança tem o efeito esperado e corrigir o rumo[^dora-lotes].

[^little]: John D. C. Little, *Little's Law as Viewed on Its 50th Anniversary*, Operations Research 59(3), 2011, revendo a prova original de 1961 e o uso em gestão de operações · https://projectproduction.org/journal/reprint-littles-law-as-viewed-on-its-50th-anniversary/
[^dora-wip]: DORA, *Work in process limits* · https://dora.dev/capabilities/wip-limits/
[^dora-lotes]: DORA, *Working in small batches* · https://dora.dev/capabilities/working-in-small-batches/

## Issues

- **Toda issue nova passa pela triagem humana** antes de alguém trabalhar nela, humano ou IA. Ela entra com `precisa-triagem` e sai com `triagem/aceita`, com `triagem/precisa-informacao` ou fechada. Se o risco for baixo, a triagem pode marcar também `pode-ser-automatica`, e aí uma IA abre o PR e um humano valida ([ADR-0015](adr/0015-triagem-humana-contra-injecao-de-prompt.md)).
- Cada checkpoint tem uma **issue-mãe** com as tarefas como **sub-issues**.
- Quando uma tarefa depende de outra, ela fica marcada como **bloqueada por** (`is:blocked` mostra todas). Não comece uma tarefa bloqueada.
- Toda tarefa diz **o que fazer** e **pronto quando**, com critérios verificáveis.
- Labels de área (`area:*`) dizem de que parte do software a tarefa é; `caminho-nativo` marca a troca de algo pago ou proprietário por algo livre ([ADR-0005](adr/0005-integracoes-terceirizadas-caminho-nativo.md)).

### Labels

| Grupo | Labels | Pra quê |
| --- | --- | --- |
| Tipo | `tarefa`, `ideia`, `bug`, `feedback`, `adr`, `documentation`, `traducao`, `pruminho` | O que a issue é. `ideia` é proposta pra discutir; `tarefa` já diz o que fazer e quando está pronta; `feedback` é relato de quem usa |
| Triagem | `precisa-triagem`, `triagem/aceita`, `triagem/precisa-informacao`, `pode-ser-automatica` | Onde a issue está na triagem humana ([ADR-0015](adr/0015-triagem-humana-contra-injecao-de-prompt.md)) |
| Área | `area:produto`, `area:arquitetura`, `area:ux`, `area:qualidade`, `area:infra`, `area:integracoes`, `area:ia-custos`, `area:gestao` | De que parte do software a tarefa é |
| Especiais | `caminho-nativo`, `automacao`, `accessibility`, `good first issue`, `help wanted` | Troca por algo livre, issue aberta pela rotina automática, barreira de acessibilidade e convites pra quem quer contribuir |

Duplicada ou "não vai ser feita" não tem label: a issue é fechada com o motivo do próprio GitHub. Label nova entra nesta tabela no mesmo PR que a cria.

## Decisões (ADR)

Decisão que muda arquitetura, dependência, formato de arquivo ou forma de trabalhar vira um ADR em `docs/adr/`, no formato de Michael Nygard (contexto, decisão, consequências), sempre com fontes. Um ADR começa como **Proposto** e só vale quando muda pra **Aceito**. Decisão revista não é apagada: o ADR novo marca o antigo como substituído.

Referência: Michael Nygard, *Documenting Architecture Decisions* (2011) · https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions

## Desenhos

Diagramas são bem-vindos quando ajudam e opcionais sempre. Qualquer ferramenta vale se o desenho aparecer no GitHub e a fonte editável ficar no repositório ([ADR-0009](adr/0009-diagramas.md)).

## Revisão

- Nada entra na `main` sem PR e sem aprovação de um mantenedor.
- No Review, quem reprova classifica a causa e diz pra onde a tarefa volta: regra de negócio → SPEC, implementação → PLAN, ajuste visual → fluxo curto ([ADR-0002](adr/0002-review-volta-pela-causa.md)). Por enquanto, quem revisa é o mantenedor.
- Todo PR diz se usou IA e qual. Quem envia responde por cada linha.
- Cada PR sai direto da `main`. Evite empilhar PRs: depois do squash de um, o seguinte entra em conflito.

## Verificações automáticas

Antes da revisão humana, todo PR passa por verificações que rodam sozinhas. Elas existem porque erro de citação passava na leitura: o mesmo tipo de erro apareceu quatro vezes no C0 ([origem das regras](origem-das-regras.md#origem-aponta-pra-evidência-que-mora-no-repositório)).

| Verificação | O que olha | Efeito |
| --- | --- | --- |
| Testes do núcleo | `gofmt`, `go vet` e `go test` em Linux, Windows e macOS | Trava o PR |
| Referências da documentação | `ADR-0000` sem link, nota de rodapé citada e não definida (ou o contrário), link interno quebrado em `docs/` e no `AGENTS.md` | Trava o PR |
| Números de issue | `#123` sem link | Só avisa, até o padrão de citação ([#121](https://github.com/vitorsantss/prumo/issues/121)) |
| Links externos | Se os links de fora ainda abrem | Roda toda segunda e sob demanda, sem travar PR, porque site de terceiro fora do ar não é erro do PR |

Pra rodar antes de abrir o PR:

```sh
node site/scripts/conferir-referencias.mjs             # referências
node site/scripts/conferir-referencias.mjs --externos  # também abre os links de fora
```

**O comando é do Prumo; só o gatilho é do GitHub.** Os arquivos em `.github/workflows/` dizem quando rodar. Numa troca de forja (GitLab, Forgejo), muda o gatilho e o comando continua o mesmo.

## Onde as coisas moram

| O quê | Onde |
| --- | --- |
| Código, ADRs e documentação | Este repositório |
| Tarefas e checkpoints | Issues e milestones |
| Perguntas e conversas abertas | Discussions |
| Rascunhos e ideias cruas | Onde cada pessoa preferir. Só passam a existir pro projeto quando viram issue ou ADR |
