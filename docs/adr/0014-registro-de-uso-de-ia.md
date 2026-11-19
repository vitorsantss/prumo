# 0014 · Registro de uso de IA por sessão

- **Status:** Aceito
- **Data:** 2026-11-16
- **Afeta:** `pesquisa/uso-de-ia/` e o `usage.jsonl` do pipeline

## Contexto
O Prumo é uma pesquisa sobre construir software com IA com qualidade, e uma das perguntas é quanto isso custa: em tokens, por modelo, por ferramenta e por tarefa. A outra é se o Prumo gasta menos que os outros jeitos de trabalhar conforme cresce ([#40](https://github.com/vitorsantss/prumo/issues/40)). Sem registro, essa comparação não existe.

Já temos uma parte. Dentro do pipeline do Prumo, o `usage.jsonl` grava uma linha por chamada de IA ([ADR-0001](0001-arquivos-como-fonte-da-verdade.md)), com os nomes das convenções GenAI do OpenTelemetry ([ADR-0007](0007-modelo-de-execucao.md)). Mas o próprio Prumo é construído fora do pipeline, em conversas com ferramentas como Claude Code, claude.ai, Orca ou Maestri. Esse gasto não é registrado em lugar nenhum.

O que a pesquisa feita pra este ADR encontrou (fontes nas notas de rodapé):
- **A ferramenta mede por sessão, não por tarefa.** O Claude Code anexa `session.id` a todas as métricas que exporta[^cc-otel], e o uso que o app do Claude mostra também é da conversa inteira. Uma sessão costuma tocar vários PRs, e um PR pode levar várias sessões.
- **Leitura de cache domina o gasto em trabalho com agentes.** O benchmark do Akita mostra runs com 5 a 15 milhões de tokens de leitura de cache, e o custo saía subestimado enquanto esses tokens não eram contados[^akita-preco]. O nosso `usage.jsonl` hoje só tem entrada e saída: tem o mesmo buraco.
- **O resultado depende do conjunto todo.** Modelo, harness, prompt, ferramentas, contexto, execução e auditoria influem juntos[^akita-metodo]. Registrar só o modelo não basta.

## Decisão
1. **O grão é a sessão de trabalho.** Cada registro é uma sessão, com os números como a ferramenta mediu. O custo de uma tarefa ou PR é calculado depois, somando as sessões que tocaram nele.

   *Por quê:* numa sessão que tocou n PRs, só se conhece o total T = x₁ + … + xₙ. É uma equação com n incógnitas, e pra n ≥ 2 ela tem infinitas soluções. Qualquer divisão por PR é escolha, não medida. Do detalhe sempre dá pra chegar no total; do total não dá pra voltar ao detalhe. É a regra da modelagem dimensional: declarar o que cada registro representa[^kimball-grao], começando pelo grão mais atômico, que é decidido pela realidade física da fonte do dado[^kimball-fisico].

2. **Sessão compartilhada usa pesos que somam 1.** Quando uma sessão toca mais de uma tarefa, o registro lista as tarefas com um peso cada, somando 1, e o resultado de cada tarefa sai marcado como `rateado`. Copiar o total pra cada tarefa contaria o gasto duas vezes; o peso por linha evita isso[^kimball-ponte]. Com uma tarefa só, o peso é 1 e o custo é exato. Por isso a recomendação é **uma sessão por tarefa** sempre que der.

   ```mermaid
   flowchart LR
     subgraph S["Sessões: o que a ferramenta mede"]
       A["Sessão A<br/>total T_A, medido"]
       B["Sessão B<br/>total T_B, medido"]
     end
     subgraph T["Tarefas: o que se quer saber"]
       X["Tarefa X"]
       Y["Tarefa Y"]
     end
     A -- "peso 1 · exato" --> X
     B -- "peso 0,5 · rateado" --> X
     B -- "peso 0,5 · rateado" --> Y
   ```

   custo(X) = 1·T_A + 0,5·T_B e custo(Y) = 0,5·T_B. Somando as tarefas volta T_A + T_B: nada é contado duas vezes. A sessão B só conhece o total, então o 0,5 é escolha e sai marcado `rateado`; a sessão A tem uma tarefa só, então o custo de X que veio dela é exato.

3. **Os quatro tipos de token, com os nomes do OpenTelemetry:**
   - `gen_ai.usage.input_tokens`
   - `gen_ai.usage.output_tokens`
   - `gen_ai.usage.cache_read.input_tokens`
   - `gen_ai.usage.cache_creation.input_tokens`

   Esses nomes estão no registro de atributos[^otel-registro]. Pela convenção, `input_tokens` **já inclui** os tokens de cache[^otel-registro]. Algumas APIs mandam a entrada sem o cache, e confundir as duas leituras conta o cache duas vezes[^contrato-cache]. O registro guarda a entrada no formato da convenção (entrada sem cache + leitura + escrita) e diz em `prumo.entrada_inclui_cache` qual leitura usou. A convenção ainda está em desenvolvimento e mudou de casa: os atributos foram movidos pro repositório das convenções GenAI, onde o token de escrita de cache também aparece como `cache_write.input_tokens`[^otel-genai]. Por isso o registro diz a versão de esquema que segue, e uma troca de nome é migração do esquema, não mudança desta decisão.

4. **Medição bruta é imutável; métrica é recalculável e versionada.** O registro da sessão nunca é editado. Total por modelo, custo por tarefa e custo por etapa aprovada são calculados a partir dos registros, e cada número calculado leva a versão do método que o gerou. Se o método mudar, o histórico inteiro é recalculado a partir do bruto. É o *event sourcing*: o registro de eventos permite reconstruir estados passados e ajustar a mudanças retroativas[^fowler]. Números de métodos diferentes não se comparam direto, como no benchmark do Akita, em que v1 e v2 não são comparáveis[^akita-metodo].

5. **Todo número diz de onde veio** (proveniência[^prov]):
   - `medido`: veio da ferramenta (contador de sessão, painel, API);
   - `estimado`: calculado por um método escrito no próprio registro;
   - `sem_dado`: a ferramenta não mostra uso.

   Número declarado pela IA sem vir da ferramenta é `estimado`, nunca `medido`: a IA não conta os próprios tokens com precisão.

6. **Duas datas por registro.** Uma é quando a sessão aconteceu (`inicio`, `fim`) e outra é quando foi registrada (`registrado_em`). Isso permite preencher o passado depois sem fingir que foi medido na hora[^retroativo].

7. **Formato: JSON Lines, um arquivo por sessão,** em `pesquisa/uso-de-ia/AAAA/AAAA-MM-DD-<id>.jsonl`. É uma linha por modelo usado na sessão.
   - JSON Lines é texto puro, um registro por linha, bom pra arquivo de log e pra processar um registro por vez[^jsonl], e qualquer ferramenta lê. Migrar pra um banco é importar linha por linha.
   - Um arquivo por sessão evita conflito de merge, porque cada branch cria um arquivo novo em vez de mexer no fim do mesmo arquivo.
   - Banco só quando uma medida mostrar que arquivo não dá conta ([#108](https://github.com/vitorsantss/prumo/issues/108)).

8. **Campos do registro:**

   | Campo | O que guarda |
   | --- | --- |
   | `esquema` | versão do formato, por exemplo `prumo.uso-ia/1` |
   | `sessao` | identificador opaco; **nunca** link de sessão ou de conta |
   | `inicio`, `fim`, `registrado_em` | datas em ISO 8601 |
   | `ferramenta` | onde o trabalho aconteceu: `claude-code`, `claude.ai`, `orca`… |
   | `gen_ai.provider.name`, `gen_ai.request.model` | provedor e modelo |
   | os quatro `gen_ai.usage.*` | tokens, ou `null` quando `sem_dado` |
   | `prumo.entrada_inclui_cache` | `true` quando a entrada segue a convenção |
   | `origem` | `medido`, `estimado` ou `sem_dado` |
   | `metodo_estimativa` | obrigatório quando `estimado` |
   | `tarefas` | lista de `{ "ref": "github:vitorsantss/prumo#95", "peso": 0.5 }`, com pesos somando 1. A referência diz de qual sistema vem, pra poder ser traduzida numa migração; quando existirem as tarefas do próprio Prumo, vira `prumo:<id>`, no formato que o ADR da pasta da tarefa definir |
   | `custo` | `{ "tipo": "assinatura" \| "api", "equivalente_usd", "tarifa_em" }` |
   | `periodo` | `inicial` nos registros do período de exceção |

9. **O período inicial é uma exceção declarada.** Ele vai até existirem o esquema validável e a checagem do núcleo (ver Consequências). O fim é um fato verificável, não uma data: é o PR que entrega a checagem. Tudo o que vier antes disso entra com `periodo: "inicial"`, `origem` honesta (quase sempre `estimado` ou `sem_dado`) e o método escrito. Depois disso, toda sessão de IA que gerar PR registra o uso na mesma mudança que entrega o trabalho. Quem ler os dados sabe onde começa o número medido.

10. **Custo com data, e assinatura com custo marginal zero.** Custo convertido leva a tarifa e a data em que ela foi conferida, porque a tarifa muda e um número antigo não se compara com um atual[^akita-preco]. Em plano de assinatura, o custo marginal é 0, com o equivalente em API ao lado, como no benchmark do Akita[^akita-metodo].

11. **Nada de conteúdo.** Só números, nomes de ferramenta e modelo, e referências a tarefas. Nenhum trecho de conversa, prompt, link de sessão ou dado pessoal ([ADR-0003](0003-dado-sensivel-fora-da-ia.md), [ADR-0012](0012-privacidade-e-coleta.md)).

## Consequências
- O custo por tarefa passa a ser um dado de pesquisa com prova de origem, e a comparação entre modelos e ferramentas fica possível ([#40](https://github.com/vitorsantss/prumo/issues/40)).
- Trabalhar numa sessão por tarefa vira boa prática, porque deixa o custo exato.
- O `usage.jsonl` do pipeline precisa ganhar os dois campos de cache pra seguir a mesma regra.
- É preciso um esquema validável (JSON Schema) e uma checagem do núcleo, disparada pela forja como define o [ADR-0016](0016-ciclo-de-vida-do-adr.md), que confira pesos somando 1, `metodo_estimativa` quando `estimado` e ausência de links. O PR que entregar essa checagem encerra o período inicial.
- O totalizador na página inicial e o painel público ([#92](https://github.com/vitorsantss/prumo/issues/92)) passam a ler destes arquivos.
- Comparar o Prumo com outras ferramentas só vale entre tarefas parecidas ou num caso de referência repetido ([#7](https://github.com/vitorsantss/prumo/issues/7)). Tarefas reais diferentes não se comparam direto.

## Fontes
[^cc-otel]: Anthropic, *Claude Code: Monitoring usage* · https://code.claude.com/docs/en/monitoring-usage
[^akita-preco]: Fabio Akita, *LLM Benchmarks: Grok 4.5 e GPT 5.6 Sol* (09/07/2026) · https://akitaonrails.com/2026/07/09/llm-benchmark-grok-4-5-gpt-5-6-sol/
[^akita-metodo]: Fabio Akita, *Novo LLM Benchmark: refiz todos os testes!* (30/07/2026) · https://akitaonrails.com/2026/07/30/novo-llm-benchmark-refiz-todos-os-testes/
[^kimball-grao]: Kimball Group, *Declaring the Grain* · https://www.kimballgroup.com/2003/03/declaring-the-grain/
[^kimball-fisico]: Ralph Kimball, *Keep to the Grain in Dimensional Modeling*, 2007 · https://www.kimballgroup.com/2007/07/keep-to-the-grain-in-dimensional-modeling/
[^kimball-ponte]: Margy Ross (Kimball Group), *Design Tip 166: Potential Bridge (Table) Detours*, 2014 (fator de alocação por linha da ponte contra contagem a mais) · https://www.kimballgroup.com/?p=20451
[^otel-registro]: OpenTelemetry, *Gen AI attributes registry* (a nota de `input_tokens` diz que inclui os tokens de cache; a página marca os `gen_ai.usage.*` como movidos pro repositório das convenções GenAI) · https://opentelemetry.io/docs/specs/semconv/attributes-registry/gen-ai/
[^contrato-cache]: Agenta, pull request 5720 (contrato de `input_tokens` com cache) · https://github.com/Agenta-AI/agenta/pull/5720
[^otel-genai]: OpenTelemetry, repositório *semantic-conventions-genai*, página de OpenAI (atributos com selo Development; `input_tokens` como total; escrita de cache como `cache_write.input_tokens`) · https://raw.githubusercontent.com/open-telemetry/semantic-conventions-genai/main/docs/gen-ai/openai.md
[^fowler]: Martin Fowler, *Event Sourcing* · https://martinfowler.com/eaaDev/EventSourcing.html
[^prov]: W3C, *PROV-DM: The PROV Data Model*, Recomendação de 2013 (define proveniência) · https://www.w3.org/TR/prov-dm/
[^retroativo]: InfoQ, *Retroactive and Future Events in an Event Sourced System* · https://www.infoq.com/news/2018/02/retroactive-future-event-sourced
[^jsonl]: *JSON Lines* · https://jsonlines.org
