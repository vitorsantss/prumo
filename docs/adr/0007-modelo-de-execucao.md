# 0007 · Modelo de execução: workspace, worktree por tarefa e rastro por etapa

- **Status:** Aceito (revisão ao fim do C2)
- **Data:** 2026-11-15

## Contexto

O desenho do mantenedor (15/11) propôs: um **workspace** por projeto, numa máquina; um **contexto de arquivos compartilhado**; cada **etapa** rodando num terminal limpo, sem herdar contexto, como um runner; e **entrada e saída de arquivos** por etapa e por tarefa, pra saber o que cada etapa leu, criou, quanto tempo levou e quanto consumiu.

Restrições: guardar tudo pesa; o repositório de quem usa não pode virar depósito de rastro; e o modelo precisa funcionar hoje no computador da pessoa e amanhã numa VPS ou em runners (ADR-0006).

## Decisão

### Workspace e tarefas
1. **Cada projeto é um workspace**: um repositório git registrado no Prumo.
2. **Cada tarefa roda num `git worktree` próprio**, num branch próprio. Tarefas em paralelo não pisam umas nas outras, e a saída de código de uma tarefa vira um PR.
3. O que é do projeto inteiro (documentação, convenções, mapa de risco) é **contexto compartilhado, só leitura** pras etapas; mudar isso é PR como qualquer outro.
4. No C1, **uma tarefa por vez**, com o limite de concorrência já configurável.

### Etapas
5. **O isolamento é de contexto, não de máquina.** Cada etapa abre uma sessão nova do agente (um processo novo), com diretório e variáveis de ambiente definidos pelo núcleo. Nada de histórico de conversa passa de uma etapa pra outra.
6. **Etapas só se comunicam por artefatos declarados** no `pipeline.yml`: o que uma etapa entrega é o que a próxima lê (estilo pipes and filters). No C1 isso é convenção verificada (o núcleo confere se as saídas declaradas existem); a trava (expor à etapa só o declarado) vem depois.

### Artefato × rastro
7. **Artefatos** são o que a etapa produz de propósito (spec.md, plan.md, código). Código fica no worktree/git; documentos ficam na pasta da tarefa.
8. **Rastro** é o que aconteceu: arquivos lidos e escritos (caminho, tamanho e hash, **nunca cópia**), duração, tokens, custo, status e o caminho do log. Cada etapa grava um único `etapa.json`; a tarefa grava um `tarefa.json` que agrega as etapas.
9. **Escritos** vêm do `git status` do worktree antes e depois da etapa. **Lidos** vêm do adapter do executor, quando o agente permite (ex.: hooks de uso de ferramenta); quando não permite, o rastro registra que a informação não está disponível. Isso entra no contrato da porta Executor.
10. Os nomes dos campos de IA no rastro seguem as **convenções semânticas de GenAI do OpenTelemetry** (`gen_ai.request.model`, `gen_ai.usage.input_tokens`, `gen_ai.usage.output_tokens`…), com a tarefa pensada como *trace* e a etapa como *span*, pra permitir exportar pra ferramentas de observabilidade no futuro sem converter nomes. Exportar de fato não é escopo do C1.

### Onde mora e por quanto tempo
11. **Nada disso fica no repositório do projeto nem é versionado.** Seguindo a XDG Base Directory Specification:
    - rastro e artefatos das tarefas: `$XDG_DATA_HOME/prumo/<workspace>/tasks/<id>/` (padrão `~/.local/share/prumo/…`);
    - logs completos: `$XDG_STATE_HOME/prumo/<workspace>/logs/` (padrão `~/.local/state/prumo/…`), que é onde a especificação manda guardar histórico de ações e logs;
    - a variável `PRUMO_HOME` sobrepõe tudo, pra quem quiser uma pasta só; no Windows e no macOS, usam-se os diretórios equivalentes do sistema.
12. **Retenção:** durante a tarefa, guarda tudo; na entrega, ficam a `entrega.md`, os `etapa.json` e o `tarefa.json` (poucos KB); **logs completos são apagados após 30 dias** (configurável).

## Consequências

- Depurar uma etapa é abrir um arquivo: o que leu, o que escreveu, quanto custou e o log.
- O armazenamento cresce devagar: referências em vez de cópias, e logs com prazo.
- Uma etapa com entradas e saídas declaradas pode, no futuro, rodar em outra máquina (runner) sem mudar o modelo: é o que torna as fases 2 e 3 da infraestrutura possíveis.
- O núcleo em Go separa desde já, por interface, o **coordenador** (decide o que roda, guarda o estado) do **executor de etapa** (roda uma etapa e devolve o `etapa.json`). No C1 os dois ficam no mesmo processo.
- Atualiza o ADR-0001: os arquivos continuam sendo a fonte da verdade, mas ficam na pasta de dados do Prumo, legíveis e não versionados.
- As convenções GenAI do OpenTelemetry ainda estão em desenvolvimento; nomes podem mudar e o rastro precisa acompanhar.

## Fontes

- XDG Base Directory Specification 0.8 · https://specifications.freedesktop.org/basedir-spec/latest/
- OpenTelemetry, convenções semânticas de GenAI (spans) · https://opentelemetry.io/docs/specs/semconv/gen-ai/gen-ai-spans/
- git-worktree · https://git-scm.com/docs/git-worktree
- Shaw e Garlan, *Software Architecture: Perspectives on an Emerging Discipline* (1996), estilo pipes and filters
- Dex Horthy, *Advanced Context Engineering for Coding Agents* · https://humanlayer.dev/blog/advanced-context-engineering
