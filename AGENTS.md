# AGENTS.md

Instruções pra agentes de IA (e pra pessoas) que vão trabalhar neste repositório. Formato aberto: https://agents.md

## O que é o Prumo

Plataforma local-first que conduz tarefas de software em etapas com portões (gates): Issue → Research → SPEC → PLAN → Código → Review → Entrega. Visão, propósito e estado atual: [README.md](README.md).

## Antes de mexer em qualquer coisa

1. Leia [docs/processo.md](docs/processo.md): checkpoints, issues, ADRs e revisão.
2. Leia os ADRs em [docs/adr/](docs/adr/). Eles são as decisões em vigor; não contradiga um ADR **Aceito** sem propor um novo.
3. Pegue uma issue aberta do checkpoint atual que **não esteja bloqueada** (`is:blocked` mostra as bloqueadas). Se ela depende de algo, termine a dependência primeiro.

## Mapa do repositório

| Caminho | O que tem |
| --- | --- |
| `docs/adr/` | Decisões de arquitetura (Nygard), com fontes |
| `docs/processo.md` | Como o projeto é conduzido |
| `CONTRIBUTING.md` | Fluxo de contribuição |
| `.github/` | Modelos de issue e PR, dono do código, CI |
| `.claude/commands/` | Comandos do Claude Code pro dia a dia (ver `docs/fluxos/`) |
| `docs/fluxos/` | Jeitos de trabalhar neste projeto, um arquivo por ferramenta |
| `site/` | Site (Starlight): página inicial e documentação gerada a partir de `docs/` (ADR-0010) |
| `cmd/prumo/` | Ponto de entrada da CLI |
| `internal/pipeline/` | Leitura e validação do `pipeline.yml` |
| `internal/coordenador/` | Decide o que roda e grava o rastro (ADR-0007) |
| `internal/executor/` | Porta Executor e adapters: dublê, modelo (sem IA), Claude Code |
| `internal/gate/` | Decisão humana entre etapas |
| `internal/rastro/` | `etapa.json`, `tarefa.json`, `usage.jsonl`, `gates.jsonl` |
| `internal/tarefa/`, `internal/local/` | Pasta da tarefa e diretórios XDG |
| `exemplos/` | Pipeline e tarefa de exemplo (Rinha de Backend) |

Núcleo em Go (ADR-0006); a interface web (TypeScript/React) ainda não existe.

## Comandos

- Testes: `go test ./...` (sempre antes de abrir PR)
- Formatação: `gofmt -w .`
- Análise: `go vet ./...`
- Site: `cd site && npm install && npm run dev` (precisa de Node 22+)
- Rodar sem IA: `go run ./cmd/prumo executar --pipeline exemplos/pipeline.yml --tarefa exemplos/rinha-2024q1/tarefa.md`

## Jeitos de trabalhar

O estado do projeto mora no GitHub (issues, PRs, ADRs, docs); este arquivo é o contrato; cada ferramenta é só um adaptador. Guias por ferramenta em [docs/fluxos/](docs/fluxos/).

## Regras

- **Decisão nova de arquitetura** (dependência, formato de arquivo, padrão) vira ADR em `docs/adr/` no mesmo PR, com fontes.
- **Dependência nova tem que ser livre**; ferramenta paga ou proprietária só como adapter opcional, com issue `caminho-nativo` aberta (ADR-0005).
- **Nunca** coloque em código, PR, issue ou prompt: segredos, dados pessoais, dados de terceiros ou código sem licença compatível com a AGPL-3.0 (ADR-0003).
- **Commits**: mensagem em português dizendo o porquê. Pode incluir `Co-Authored-By` do agente usado. **Não inclua links de sessão, de conversa ou de conta privada** em commits, PRs ou issues.
- **PR**: preencha o modelo, inclusive "Como foi feito" (com ou sem IA, e qual). Nada vai direto pra `main`.
- **Citações**: toda afirmação de fora tem fonte com link (nota de rodapé na documentação); ADRs e issues citados como `ADR-0007` e `#50` viram link no site.
- **Site e visual**: siga [docs/como-mudar-o-site.md](docs/como-mudar-o-site.md) e o [design system](docs/design/README.md) (ADR-0011). Use só tokens (`var(--p-…)`); token novo entra em `docs/design/tokens.json`. Componente novo ou mudança visual só entra com a página do componente em `docs/design/componentes/` atualizada ([como propor](docs/design/como-propor-um-componente.md)).
- **Documentação muda junto com o código.** Se o comportamento mudou e a documentação não, o PR está incompleto.

## Pesquisa e prova

O Prumo é uma pesquisa: decisão sem prova contamina as decisões que vierem depois. Estas regras valem pra qualquer IA e pra qualquer pessoa. Cada uma nasceu de um erro real, contado em [origem das regras](docs/origem-das-regras.md).

- **Fonte antes da tese.** Na etapa de Research, toda afirmação sobre o mundo e toda proposta de decisão vêm com referência. Opinião da IA não é prova. Se não achou fonte, diga "não achei fonte" e não apresente a ideia como decisão. ([origem](docs/origem-das-regras.md#fonte-antes-da-tese))
- **Raciocínio também precisa de prova.** Lógica ou matemática valem, mas escritas por extenso e apoiadas numa referência que mostre que o princípio é aceito. Exemplo: a equação do ADR-0014, apoiada na modelagem dimensional de Kimball.
- **Link só depois de aberto ou encontrado numa busca.** Nunca monte URL de memória. Se a página não abriu, não cite. ([origem](docs/origem-das-regras.md#link-só-depois-de-aberto))
- **Origem aponta pra evidência que mora no repositório.** Ao dizer de onde algo veio (um erro, uma medição, uma pesquisa), aponte um arquivo do repositório, uma nota de rodapé ou um commit da `main`. Número de issue ou PR pode vir junto como atalho, mas não sozinho: só tem sentido no GitHub, e o Prumo não pode depender dele. Data solta também não serve: não leva a lugar nenhum e fica ambígua sem o ano. Se a evidência for privada, como uma conversa numa conta pessoal, descreva o que foi e qual instrumento mediu, sem link. ([origem](docs/origem-das-regras.md#origem-aponta-pra-evidência-que-mora-no-repositório))
- **Separe fonte de proposta.** Diga o que a fonte afirma e o que é conclusão sua a partir dela. Conclusão vira decisão do Prumo no ADR, com as fontes de onde partiu, nunca disfarçada de literatura.
- **ADR aceito também se revalida.** Fonte nova, ADR novo no mesmo tema ou mudança no método de validação reabrem a checagem (#109).

## Fila de trabalho

- **Ideia ou relato vira issue, não código.** Quando alguém despeja ideias ou relata um problema, anote em issues e só implemente quando pedirem. ([origem](docs/origem-das-regras.md#ideia-ou-relato-vira-issue-não-código))
- **Siga a ordem combinada.** Se a prioridade foi definida, não puxe outra coisa por ser rápida ou urgente sem combinar antes.
- **Procure antes de abrir.** Antes de criar issue ou ADR, busque se já existe. ([origem](docs/origem-das-regras.md#procure-antes-de-abrir))
- **Uma sessão de IA por tarefa,** sempre que der: deixa o custo da tarefa exato (ADR-0014).

## Mudança visual

Segue as [regras da camada Site](docs/design/site.md) e o roteiro de [como mudar o site](docs/como-mudar-o-site.md). O design system é a fonte: regra visual nova entra lá, não aqui.

## Idioma

Português do Brasil em documentação, issues, PRs e commits. Nomes no código seguem a convenção da linguagem.
