# Claude Code no seu computador

Um agente só, com acesso ao seu terminal, ao navegador e ao GitHub. Funciona no terminal puro, no VS Code ou num canvas de agentes como o Maestri (macOS e Windows).

## Instalar (uma vez)

| O quê | Windows | macOS / Linux |
| --- | --- | --- |
| Git | `winget install Git.Git` | já vem ou pelo gerenciador de pacotes |
| Go 1.24+ | `winget install GoLang.Go` | `brew install go` ou go.dev/dl |
| GitHub CLI | `winget install GitHub.cli` | `brew install gh` |
| Claude Code | instalador oficial (docs.claude.com) | instalador oficial |

Depois:

```sh
gh auth login                 # o agente usa o gh pra issues e PRs
git clone https://github.com/vitorsantss/prumo
cd prumo
go test ./...                 # confere se está tudo certo
claude --chrome               # abre o Claude Code com acesso ao Chrome
```

O `--chrome` liga a integração com a extensão Claude in Chrome: o mesmo agente consegue mexer em sites (configurações do GitHub, por exemplo) sem você trocar de ferramenta. Dá pra ligar depois com `/chrome`.

## Comandos do projeto

Ficam em `.claude/commands/` e aparecem quando você digita `/`:

| Comando | O que faz |
| --- | --- |
| `/proximo` | Lê o resumo diário, os PRs e as issues do checkpoint e te mostra **uma decisão por vez**, com opções e recomendação |
| `/ideia <texto>` | Anota a ideia como issue `ideia`, sem desviar o checkpoint |
| `/revisar <número do PR>` | Te guia pelo PR na ordem de leitura e transforma suas dúvidas em comentários |
| `/como-funciona <assunto>` | Explica uma parte do projeto a partir dos docs e ADRs |

## Suas preferências pessoais

O jeito como **você** gosta que o agente trabalhe (tom, quanto explicar, se ele deve questionar suas decisões) não é regra do projeto. Guarde no arquivo pessoal do Claude Code, fora do repositório: `~/.claude/CLAUDE.md` (no Windows, `C:\Users\<você>\.claude\CLAUDE.md`). Exemplo:

```markdown
# Como eu gosto de trabalhar
- Responda em português.
- Aja como mentor: questione minhas decisões e explique o porquê.
- Me mostre uma decisão por vez, com opções e a sua recomendação.
- Ideia solta que eu mencionar vira issue `ideia`, sem mudar o checkpoint.
- Proponha pontos de parada quando o checkpoint estiver garantido.
```

## No Maestri

O Maestri é um canvas onde cada terminal é um nó. Um arranjo que funciona bem pra este projeto:

- **Terminal "trabalho"**: Claude Code na pasta do repositório, com `--chrome`. É onde você roda `/proximo`.
- **Nota "checkpoint"**: o objetivo do checkpoint atual e o que falta, que o agente pode atualizar.
- **Terminal "revisor"** (opcional): um segundo agente conectado ao primeiro, só pra revisar PRs antes de você.
- **Navegador** apontando pro board do projeto no GitHub.

Nada disso guarda estado do projeto: se fechar o Maestri, o GitHub continua com tudo.

## Cuidados

- O agente tem acesso ao seu terminal e ao seu navegador logado. Leia o que ele pede antes de aprovar, principalmente comandos que apagam coisas ou ações em sites.
- Segredos e dados pessoais nunca entram no contexto (ADR-0003).
