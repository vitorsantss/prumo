# 0006 · Execução local-first e divisão da stack

- **Status:** Aceito (revisão obrigatória ao fim do C2)
- **Data:** 2026-11-14

## Contexto

O Prumo executa agentes de IA, roda comandos, sobe ambientes e mexe em arquivos de um projeto. Isso exige uma máquina. Hoje essa máquina é o computador de quem usa; no futuro pode ser uma VPS ou runners sob demanda. Também queremos uma versão web e, quem sabe, mobile.

Ao mesmo tempo, quem for contribuir precisa conseguir rodar o Prumo na própria máquina com pouco esforço e sem pagar nada ([ADR-0005](0005-integracoes-terceirizadas-caminho-nativo.md)).

O mantenedor domina Java/Spring e React/TypeScript. O projeto vai ter partes com necessidades diferentes, então pode ter mais de uma linguagem.

## Decisão (proposta)

1. **Local-first.** O núcleo do Prumo (o motor que lê o `pipeline.yml`, executa as etapas e grava em `tasks/<id>/`) roda na máquina de quem usa, como um processo local com uma CLI.
2. **Interface web servida localmente.** A tela é uma aplicação web que o próprio núcleo serve em `localhost` e abre no navegador. Não nasce como app desktop nativo.
3. **Uma base, três formas.** A mesma combinação núcleo + interface web serve pra:
   - rodar no computador da pessoa (agora);
   - virar app desktop depois, embrulhando a interface com uma casca leve (ex.: Tauri), se houver motivo;
   - rodar numa VPS, com o mesmo núcleo atrás de login, quando chegar a hora (fase 1 da infraestrutura).
4. **Poliglota por fronteira, não por gosto.**
   - Núcleo: **Go** (ver "Linguagem do núcleo").
   - Interface: TypeScript + React, seguindo o design system.
   - Adapters (integrações, executores): **qualquer linguagem**, desde que falem com o núcleo por processo/CLI ou por MCP. Assim alguém pode contribuir com um adapter em Python mesmo que o núcleo seja em outra linguagem.
5. **Rodar pra contribuir em um comando.** O repositório oferece um caminho único de execução local (um binário ou `compose up`) e um **modo dublê**, em que os executores devolvem respostas gravadas, pra testar tudo sem IA paga.

## Linguagem do núcleo: Go

| Opção | A favor | Contra |
| --- | --- | --- |
| Java / Spring | O mantenedor domina; ecossistema maduro | Pesado pra uma CLI; distribuição exige JVM ou compilação nativa |
| Go | Binário único, fácil de distribuir; bom com processos e concorrência | Linguagem nova pro mantenedor |
| Python | Mais fácil pra estudante contribuir; maior ecossistema de IA | Distribuição e empacotamento mais trabalhosos; tipagem opcional |
| TypeScript / Node | Mesma linguagem da interface; vários agentes livres são em TS | Runtime separado; dois mundos (servidor e tela) na mesma linguagem podem misturar responsabilidades |

Critérios usados: facilidade de distribuir pra quem usa, facilidade de quem chega pra contribuir, adequação ao problema (orquestrar processos externos, ler e gravar arquivos, rodar etapas em paralelo) e licença livre de tudo que entrar.

**Escolha: Go**, por ser a opção mais orientada ao problema:

- O núcleo é, no fundo, um orquestrador de processos (agentes de IA, testes, comandos de ambiente). A biblioteca padrão do Go cobre isso (`os/exec`, `context`, goroutines) sem dependência extra.
- Gera **um binário só** por sistema operacional, com compilação cruzada nativa. Quem usa baixa e roda; quem contribui não precisa instalar runtime.
- É a linguagem de várias ferramentas de linha de comando que o público do Prumo já usa, o que facilita achar exemplo e contribuidor.
- A linguagem é pequena e tem formatação e testes padronizados (`gofmt`, `go test`), o que ajuda estudante a entrar.
- Licença BSD-3-Clause, compatível com a AGPL-3.0.

**Custo aceito:** Go é novo pro mantenedor. Isso aumenta o risco do C1; se o prazo apertar, corta-se escopo do C1, não a data. Outras partes podem ter versões em outras linguagens no futuro, via a fronteira de processo/MCP, sem mexer no núcleo.

## Consequências

- O C1 não precisa de servidor, banco nem login: só a CLI na máquina local.
- A versão web/VPS não exige reescrever o núcleo, só autenticação e isolamento (que viram ADRs próprios).
- Adapters em linguagens diferentes obrigam a manter contratos claros (formato de entrada e saída) e testes de contrato.
- "Desktop" deixa de ser uma decisão agora e vira uma casca opcional depois.

## Fontes

- Go: tutorial e documentação oficial · https://go.dev/doc/
- Go: licença (BSD-3-Clause) · https://go.dev/LICENSE

- Local-first software, Kleppmann et al. (Ink & Switch, 2019) · https://www.inkandswitch.com/local-first/
- Ports and Adapters, Alistair Cockburn · https://alistair.cockburn.us/hexagonal-architecture
- Model Context Protocol · https://modelcontextprotocol.io
- Tauri · https://tauri.app

## Dependências do núcleo (checagem de licença, ADR-0005)

| Dependência | Uso | Licença |
| --- | --- | --- |
| github.com/goccy/go-yaml | Ler o `pipeline.yml` com erros apontando linha e campo | MIT |
