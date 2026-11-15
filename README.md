# Prumo

Prumo é uma plataforma pessoal pra conduzir tarefas de software com IA em etapas, cada uma com um portão (gate) antes da próxima:

```
Issue → Research → SPEC → PLAN → Código → Review (risco) → Entrega
```

A ideia: decidir sem IA sempre que der (regras, heurísticas), usar IA só onde o trabalho é aberto, medir os tokens de cada etapa e nunca deixar código arriscado passar sem revisão humana.

## Por que "Prumo"

> História do nome em breve (ver issue "Contar a história do nome").

## Estado

Esqueleto do C1 rodando: uma CLI em Go que lê o `pipeline.yml`, executa as etapas de uma tarefa com um gate manual e grava o rastro de cada etapa (ADR-0007).

## Como rodar

Precisa do [Go](https://go.dev/dl/) 1.24 ou mais novo.

```sh
go build -o prumo ./cmd/prumo
./prumo validar exemplos/pipeline.yml
./prumo executar --pipeline exemplos/pipeline.yml --tarefa exemplos/rinha-2024q1/tarefa.md
```

Por padrão as etapas de IA usam o **executor dublê**, que responde sem IA, pra testar o fluxo de graça. Pra usar o Claude Code (pago, ADR-0005), passe `--executor claude`. O caminho livre, com agente e modelo abertos, é a issue #13.

No gate de cada etapa manual você escolhe: aprovar, pedir ajuste (a etapa roda de novo com o seu comentário) ou refazer. Tudo fica em `~/.local/share/prumo/<workspace>/tasks/<id>/` (ou em `$PRUMO_HOME`): a tarefa, a saída de cada etapa, o `etapa.json` com o que foi lido, escrito, tokens e tempo, e o `tarefa.json` com o resumo.

Testes: `go test ./...`

| Checkpoint | Até | O que entrega |
| --- | --- | --- |
| C0 · Decidir | 19/11/2026 | ADRs das decisões que destravam o código |
| C1 · Esqueleto andando | 03/12/2026 | CLI que roda 3 etapas numa tarefa real, com 1 gate manual |
| C2 · Primeiro valor | 17/12/2026 | Research + PLAN + risk gate, usado em 3 tarefas |

As decisões de arquitetura ficam em [`docs/adr/`](docs/adr/) e o jeito de trabalhar em [`docs/processo.md`](docs/processo.md).

## Como contribuir

Issues abertas, ideias e correções são bem-vindas, com ou sem IA. Comece pelo [CONTRIBUTING.md](CONTRIBUTING.md).

## Propósito

O Prumo foi feito pra educação e pro setor público: pra quem quer estudar, contribuir e praticar como se constrói software, e pra universidades, institutos federais e órgãos públicos que criam software pra sociedade.

## Licença

[AGPL-3.0](LICENSE). Quem modificar o Prumo e oferecer a versão modificada, inclusive pela rede, precisa abrir o código das mudanças. O porquê está no [ADR-0004](docs/adr/0004-licenca-agpl-3.md).

Copyright (C) 2026 Vitor Santos
