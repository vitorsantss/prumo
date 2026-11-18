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

O Prumo é uma pesquisa: decisão sem prova contamina as decisões que vierem depois. Estas regras valem pra qualquer IA e pra qualquer pessoa. Cada uma nasceu de um erro real, citado entre parênteses.

- **Fonte antes da tese.** Na etapa de Research, toda afirmação sobre o mundo e toda proposta de decisão vêm com referência. Opinião da IA não é prova. Se não achou fonte, diga "não achei fonte" e não apresente a ideia como decisão. (Origem: "registrar por sessão" foi proposto sem referência na discussão do ADR-0014, #110, e só ganhou fonte depois de cobrado.)
- **Raciocínio também precisa de prova.** Lógica ou matemática valem, mas escritas por extenso e apoiadas numa referência que mostre que o princípio é aceito. Exemplo: a equação do ADR-0014, apoiada na modelagem dimensional de Kimball.
- **Link só depois de aberto ou encontrado numa busca.** Nunca monte URL de memória. Se a página não abriu, não cite. (Origem: dois links escritos de memória durante o #110.)
- **Origem aponta pra evidência, não pra data.** Ao dizer de onde algo veio (um erro, uma medição, uma pesquisa), aponte o PR, a issue, o commit ou a nota de rodapé onde dá pra conferir. Data solta não leva a lugar nenhum, fica ambígua sem o ano, e o GitHub já guarda quando cada coisa aconteceu. Se a evidência for privada, como uma conversa numa conta pessoal, descreva o que foi e qual instrumento mediu, sem link. (Origem: datas soltas e uma citação vaga a uma sentença, apontadas na revisão do #110.)
- **Separe fonte de proposta.** Diga o que a fonte afirma e o que é conclusão sua a partir dela. Conclusão vira decisão do Prumo no ADR, com as fontes de onde partiu, nunca disfarçada de literatura.
- **ADR aceito também se revalida.** Fonte nova, ADR novo no mesmo tema ou mudança no método de validação reabrem a checagem (#109).

## Fila de trabalho

- **Ideia ou relato vira issue, não código.** Quando alguém despeja ideias ou relata um problema, anote em issues e só implemente quando pedirem. (Origem: três relatos viraram PR sem pedido, #103.)
- **Siga a ordem combinada.** Se a prioridade foi definida, não puxe outra coisa por ser rápida ou urgente sem combinar antes.
- **Procure antes de abrir.** Antes de criar issue ou ADR, busque se já existe. (A #107 duplicou a #40.)
- **Uma sessão de IA por tarefa,** sempre que der: deixa o custo da tarefa exato (ADR-0014).

## Mudança visual

- **Teste no celular com toque**, além de claro e escuro e das larguras de [como mudar o site](docs/como-mudar-o-site.md): abrir o menu, tocar num link e chegar na página. (O #95 quebrou o menu do celular, #101, porque só foi conferido com mouse.)
- **Estilo global com seletor de escopo.** Regra em CSS global não pode depender de uma classe genérica que outra página também usa. (O `.topo` da página inicial herdou a transição da documentação.)

## Idioma

Português do Brasil em documentação, issues, PRs e commits. Nomes no código seguem a convenção da linguagem.
