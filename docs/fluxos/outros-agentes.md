# Outros agentes (Codex, OpenCode, Gemini e afins)

Todo agente que segue o padrão aberto [AGENTS.md](https://agents.md) já entende o projeto: ele lê o [AGENTS.md](../../AGENTS.md) da raiz e as regras valem igual.

## Os comandos do dia a dia, como prompts

Os comandos do Claude Code em `.claude/commands/` são texto simples. Pra usar em outro agente, copie o conteúdo do arquivo como prompt, ou cadastre como comando personalizado se a ferramenta tiver esse recurso:

| Arquivo | Pra quê |
| --- | --- |
| `.claude/commands/proximo.md` | Uma decisão por vez a partir do GitHub |
| `.claude/commands/ideia.md` | Anotar ideia como issue |
| `.claude/commands/revisar.md` | Revisar um PR guiado |
| `.claude/commands/como-funciona.md` | Explicar uma parte do projeto |

Eles usam a GitHub CLI (`gh`). Se o seu agente não roda comandos, faça os passos à mão no site.

## Agente revisor

Um padrão comum é usar um agente pra escrever e outro pra revisar antes da pessoa. Se fizer isso, a decisão de merge continua sendo de um mantenedor humano (ver [processo](../processo.md)).
