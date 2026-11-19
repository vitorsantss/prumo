# Referências

Todas as fontes que sustentam as decisões do Prumo, agrupadas por assunto, com onde cada uma é usada. A regra: fonte primária antes de resumo; texto sem autor não entra.

## Processo e decisões

| Fonte | Onde usamos |
| --- | --- |
| Michael Nygard, *Documenting Architecture Decisions* (2011) · https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions | Formato dos ADRs |
| Ryan Singer, *Shape Up*, cap. "Set Boundaries" · https://basecamp.com/shapeup/1.2-chapter-03 | Checkpoints com tempo fixo e escopo variável |
| Alan Hevner et al., [*Design Science in Information Systems Research*](https://doi.org/10.2307/25148625), MIS Quarterly 28(1), 2004 | Método de pesquisa do projeto (em discussão) |

## O fluxo com IA

| Fonte | Onde usamos |
| --- | --- |
| Dex Horthy, *Advanced Context Engineering for Coding Agents* · https://humanlayer.dev/blog/advanced-context-engineering | Etapa de Research; etapas como sessões separadas ([ADR-0007](adr/0007-modelo-de-execucao.md)) |
| Anthropic, *Best practices for Claude Code* · https://code.claude.com/docs/en/best-practices | Explorar e planejar antes de codar |
| Design Council, *The Double Diamond* · https://www.designcouncil.org.uk/our-resources/the-double-diamond/ | Descoberta antes da entrega |
| Marty Cagan, *Dual-Track Agile* · https://www.svpg.com/dual-track-agile/ | Descoberta e entrega como trilhas separadas |
| GitHub, *Spec Kit* · https://github.com/github/spec-kit | Desenvolvimento guiado por especificação, referência de comparação |
| Lingjiao Chen, Matei Zaharia, James Zou, *FrugalGPT* (2023) · https://arxiv.org/abs/2305.05176 | Decidir sem IA primeiro, modelo maior só quando precisa |
| Barry Boehm, [*Software Engineering Economics*](https://openlibrary.org/books/OL4267874M) (1981) | Custo de corrigir defeitos por fase ([ADR-0002](adr/0002-review-volta-pela-causa.md)) |
| IEEE 1044-2009, [*Standard Classification for Software Anomalies*](https://doi.org/10.1109/IEEESTD.2010.5953441) | Classificar a origem de um defeito ([ADR-0002](adr/0002-review-volta-pela-causa.md)) |

## Arquitetura e execução

| Fonte | Onde usamos |
| --- | --- |
| Alistair Cockburn, *Hexagonal Architecture* · https://alistair.cockburn.us/hexagonal-architecture/ | Portas e adaptadores ([ADR-0005](adr/0005-integracoes-terceirizadas-caminho-nativo.md)) |
| Mary Shaw e David Garlan, [*Software Architecture: Perspectives on an Emerging Discipline*](https://openlibrary.org/works/OL2982368W) (1996) | Estilo pipes and filters ([ADR-0007](adr/0007-modelo-de-execucao.md)) |
| OpenTelemetry, convenções semânticas de GenAI · https://opentelemetry.io/docs/specs/semconv/gen-ai/ | Nomes dos campos do rastro ([ADR-0007](adr/0007-modelo-de-execucao.md)) |
| XDG Base Directory Specification · https://specifications.freedesktop.org/basedir-spec/latest/ | Onde o Prumo guarda dados ([ADR-0007](adr/0007-modelo-de-execucao.md)) |
| git-worktree · https://git-scm.com/docs/git-worktree | Uma cópia do repositório por tarefa ([ADR-0007](adr/0007-modelo-de-execucao.md)) |
| The Twelve-Factor App, *Config* · https://12factor.net/config | Variáveis de ambiente ([ADR-0003](adr/0003-dado-sensivel-fora-da-ia.md)) |
| OWASP, *Secrets Management Cheat Sheet* · https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html | Segredos fora da IA ([ADR-0003](adr/0003-dado-sensivel-fora-da-ia.md)) |
| AGENTS.md · https://agents.md | Porta de entrada pra agentes de IA |

## Licença e documentação

| Fonte | Onde usamos |
| --- | --- |
| GNU Affero General Public License v3 · https://www.gnu.org/licenses/agpl-3.0.html | Licença ([ADR-0004](adr/0004-licenca-agpl-3.md)) |
| Starlight · https://starlight.astro.build | Site de documentação ([ADR-0010](adr/0010-site-com-starlight.md)) |
| GitHub Docs, *Creating diagrams* · https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-diagrams | Diagramas em Mermaid ([ADR-0009](adr/0009-diagramas.md)) |

Achou uma fonte que falta ou que não sustenta o que dizemos? Abra uma tarefa.
