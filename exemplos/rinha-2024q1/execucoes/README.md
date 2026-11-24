# Execuções reais do caso Rinha

Relato das execuções do Prumo com IA de verdade neste caso. É a evidência do checkpoint C1: uma tarefa real passou por tarefa → SPEC → entrega, com o gate humano no meio.

**Arquivos.** As saídas de cada execução estão nas pastas ao lado, copiadas como estavam na pasta da tarefa: a tarefa usada, a SPEC aprovada, a entrega e o rastro (`tarefa.json`, `usage.jsonl`, `gates.jsonl` e um `etapa.json` por etapa). É uma cópia de arquivo, não a pasta de trabalho do pipeline: pelo [ADR-0001](../../../docs/adr/0001-arquivos-como-fonte-da-verdade.md), a pasta da tarefa continua fora do repositório até um ADR novo decidir onde moram os resultados. Antes de publicar, os arquivos foram conferidos: o rastro só guarda caminhos relativos, e não há identificador de sessão nem dado pessoal.

- [Execução 1, 19/11/2026](20261119-144909-registrar-transacao-e-consultar-extrato/)
- [Execução 2, 24/11/2026](20261124-115556-registrar-transacao-e-consultar-extrato/)

**Instrumento.** Os números vêm do rastro do próprio Prumo (`usage.jsonl` e `etapa.json` da tarefa), gerado na máquina do mantenedor, com o Claude Code na conta pessoal dele e o `pipeline.yml` deste diretório de exemplos. Nenhum identificador de sessão foi guardado. O custo é o equivalente em preço de tabela da API (`prumo.custo_base: list`): na assinatura, o custo extra foi zero.

## Referência: uma sessão vazia

Antes das execuções, uma pergunta de uma palavra ("Responda só: oi") ao Claude Code mostrou o piso de uma sessão nova.

| Entrada | Leitura de cache | Criação de cache | Saída | Voltas | Custo de tabela |
| --- | --- | --- | --- | --- | --- |
| 28.520 | 14.801 | 13.717 | 4 | 1 | US$ 0,11 |

Toda etapa de IA abre uma sessão nova ([ADR-0007](../../../docs/adr/0007-modelo-de-execucao.md), decisão 5), então toda etapa paga esse piso. A resposta bruta, sem identificador de sessão nem credencial, está em `internal/executor/testdata/claude-oi.json`.

## Execução 1: SPEC com um ajuste no gate (19/11/2026)

Tarefa 1, na versão anterior do arquivo, que só apontava pra especificação original.

- **Gate, tentativa 1:** pedido de ajuste, respondendo três perguntas abertas da própria SPEC: a infraestrutura entra na tarefa, payload inválido responde 422 e "use o banco que faz mais sentido".
- **Gate, tentativa 2:** aprovada. As perguntas respondidas viraram regra; as não respondidas continuaram abertas.
- **A SPEC da tentativa 1 se perdeu.** O Prumo grava a saída de cada tentativa por cima da anterior, então só a versão aprovada sobreviveu. O rastro guarda os números das duas tentativas, mas não o texto que voltou no gate.

| Tentativa | Modelo | Entrada | Saída | Custo de tabela |
| --- | --- | --- | --- | --- |
| 1 | claude-opus-5-5 | 30.521 | 3.679 | US$ 0,20 |
| 2 | claude-opus-5-5 | 90.209 | 4.700 | US$ 0,25 |

A etapa durou 823 segundos, somando a IA e o tempo humano no gate; essa versão do Prumo ainda não separava os dois nem registrava voltas e cache.

## Execução 2: SPEC aprovada direto (24/11/2026)

Mesma tarefa, com a versão do executor que grava uso por modelo.

| Modelo | Entrada | Leitura de cache | Criação de cache | Saída | Custo de tabela |
| --- | --- | --- | --- | --- | --- |
| claude-haiku-5-5 | 9.395 | 3.375 | 6.016 | 970 | US$ 0,001 |
| claude-opus-5-5 | 58.051 | 40.667 | 17.380 | 4.909 | US$ 0,25 |

- 1 tentativa, 2 voltas do agente.
- 51 segundos de IA, dentro de 93 segundos de etapa: o resto foi o tempo humano no gate.
- O agente tentou abrir a especificação original na web e foi barrado por falta de permissão.

## O que aprendemos

- **Cada volta do agente relê o contexto.** A execução 2 registrou 2 voltas e gastou pouco mais de 2 vezes o piso. Na execução 1, a primeira tentativa ficou perto do piso e a segunda chegou a cerca de 3 vezes, compatível com mais voltas; aquela versão não registrava voltas, então fica como hipótese.
- **Uma etapa pode usar mais de um modelo.** O Claude Code chamou um modelo menor pra uma tarefa auxiliar por conta própria. O Prumo não controla essa escolha; registra uma linha por modelo. A escolha de modelo por etapa ficou para o [#146](https://github.com/vitorsantss/prumo/issues/146).
- **Sem a fonte, a SPEC completa de memória.** As rotas e os limites vieram da memória do modelo, com aviso. As tarefas passaram a trazer os números da fonte, e a regra de fontes e permissões do agente ficou para o [#147](https://github.com/vitorsantss/prumo/issues/147).
- **Sem etapa de PLAN, o "como" entra na SPEC.** E uma decisão delegada à IA no gate não fica marcada como delegada: [#148](https://github.com/vitorsantss/prumo/issues/148).

- **O rastro não guarda o texto de cada tentativa.** Pra medir retrabalho (H1), é preciso ver o que mudou entre a versão devolvida e a aprovada, não só quantas vezes voltou.

Duas execuções numa tarefa não testam as hipóteses do projeto. Elas mostram que o caminho funciona e o que o rastro precisa medir.
