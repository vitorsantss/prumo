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

## Regras

- **Decisão nova de arquitetura** (dependência, formato de arquivo, padrão) vira ADR em `docs/adr/` no mesmo PR, com fontes.
- **Dependência nova tem que ser livre**; ferramenta paga ou proprietária só como adapter opcional, com issue `caminho-nativo` aberta (ADR-0005).
- **Nunca** coloque em código, PR, issue ou prompt: segredos, dados pessoais, dados de terceiros ou código sem licença compatível com a AGPL-3.0 (ADR-0003).
- **Commits**: mensagem em português dizendo o porquê. Pode incluir `Co-Authored-By` do agente usado. **Não inclua links de sessão, de conversa ou de conta privada** em commits, PRs ou issues.
- **PR**: preencha o modelo, inclusive "Como foi feito" (com ou sem IA, e qual). Nada vai direto pra `main`.
- **Documentação muda junto com o código.** Se o comportamento mudou e a documentação não, o PR está incompleto.

## Idioma

Português do Brasil em documentação, issues, PRs e commits. Nomes no código seguem a convenção da linguagem.
