# 0015 · Triagem humana antes da IA: a primeira defesa do Prumo contra injeção de prompt

- **Status:** Aceito
- **Data:** 2026-11-16

## Contexto
Relatos e sugestões chegam sempre: dezessete issues com a label `feedback` em poucos dias[^relatos]. Corrigi-los à mão estava tirando o mantenedor da entrega principal. A ideia é deixar uma IA corrigir os de baixo risco e manter o humano só onde ele decide.

Dois riscos aparecem logo:
- **Issue pública é texto de qualquer pessoa.** Greshake e colegas mostraram que, quando uma aplicação com LLM lê dados de fora, um atacante pode injetar instruções nesses dados e controlar a aplicação remotamente, sem acesso direto a ela; a recuperação de conteúdo apaga a fronteira entre dado e instrução[^greshake]. A OWASP põe a injeção de prompt em primeiro lugar entre os riscos de aplicações com LLM e diz que não está claro se existe prevenção infalível, dada a natureza estocástica dos modelos[^owasp]. Um agente que lê relato público antes de um humano é exatamente essa porta.
- **Revisão sem base vira carimbo.** O viés de automação aparece em leigos e em especialistas, não se evita com treino ou instrução, e afeta decisões individuais e de equipe[^vies]. Quem valida trabalho de IA sem um critério objetivo tende a aprovar o que a IA trouxe.

As diretrizes de IA confiável da União Europeia separam os graus de supervisão. Humano **no** laço é a capacidade de intervir em cada ciclo de decisão do sistema. Humano **sobre** o laço é a capacidade de intervir no desenho do sistema e acompanhar a operação dele[^hleg]. E, quanto menos supervisão humana, mais testes e governança mais rígida são necessários[^hleg].

## Decisão
1. **Toda issue nova entra com a label `precisa-triagem`.** Ninguém trabalha nela, nem humano nem IA, antes da triagem.
2. **A triagem é humana.** A pessoa lê a issue e decide:
   - `triagem/aceita`: faz sentido e está pronta pra ser trabalhada;
   - `triagem/precisa-informacao`: falta informação, então comenta pedindo ou fecha;
   - fechar: não faz sentido, é duplicada ou está no lugar errado.

   É o fluxo do Kubernetes: issue nova recebe `needs-triage` automaticamente e, depois de avaliada, `triage/accepted`, que indica que está pronta pra ser trabalhada[^k8s-guia][^k8s-kep].
3. **Aceita, a issue fica disponível pra humano ou IA.** A label `triagem/aceita` também mostra a quem quer contribuir o que está pronto pra pegar.
4. **A IA só trabalha sozinha em issue com `pode-ser-automatica`,** label que a pessoa põe na triagem, junto com `triagem/aceita`, quando o risco é baixo: ajuste visual, texto, link quebrado. Assim, nenhuma IA processa texto público que um humano ainda não leu.

   ```mermaid
   flowchart LR
     subgraph F["Texto não lido por humano: nenhuma IA toca"]
       N["Issue nova<br/>precisa-triagem"]
     end
     N --> H{"Humano lê<br/>(triagem)"}
     H -->|falta informação| I["triagem/precisa-informacao"]
     H -->|não faz sentido| X["Fechada"]
     H -->|aceita| A["triagem/aceita<br/>humano ou IA podem pegar"]
     H -->|aceita, baixo risco| B["+ pode-ser-automatica"]
     B --> IA["IA com privilégio mínimo<br/>sem segredos, só branch e PR"]
     IA --> R["PR em rascunho<br/>prévia, capturas, testes"]
     R --> V{"Humano valida"}
     V -->|aprova| M["Merge, humano"]
     V -->|recusa| IA
   ```

5. **A IA entrega um PR em rascunho, nunca o merge.** O PR traz a prévia, capturas do que mudou (claro, escuro e celular) e os testes passando. Se houver mais de um jeito de corrigir, as opções vêm lado a lado. A IA não aprova nem mescla o próprio trabalho. É o mesmo desenho do agente de código do GitHub: o PR em rascunho dele precisa ser revisado e mesclado por um humano, e o agente não consegue marcá-lo como pronto, aprovar nem mesclar[^copilot].
6. **Cada merge passa por uma pessoa; a rotina da IA, não passo a passo.**
   - **Merge:** uma pessoa aprova cada um, porque o merge muda o projeto. É o que a OWASP recomenda pra ação de alto risco[^owasp], e o que as diretrizes da UE chamam de humano *no* laço: intervir em cada decisão[^hleg].
   - **Rotina da IA** (reproduzir o problema, tirar capturas, rodar testes): a pessoa define as regras e as permissões da rotina e confere o resultado, sem aprovar cada passo. É o humano *sobre* o laço: intervir no desenho e acompanhar a operação[^hleg].
7. **O agente que trabalha em issue roda com o mínimo:** sem segredos, só com permissão pra criar branch e abrir PR, como recomenda a OWASP: privilégios do modelo restritos ao mínimo necessário[^owasp]. O modelo de ameaças (#93) detalha isso antes da rotina ligar.

## Como isso protege contra injeção de prompt
Este ADR abre o tema da injeção de prompt no Prumo. Cada decisão acima corresponde a uma das mitigações da OWASP pra LLM01[^owasp]:

| Camada | Mitigação da OWASP | Onde está neste ADR |
| --- | --- | --- |
| Separar o texto não confiável | Segregar e identificar conteúdo externo | Decisões 1 e 4: nenhuma IA lê issue antes de um humano; a triagem é a fronteira |
| Privilégio mínimo | Controle de privilégio e acesso mínimo | Decisão 7: sem segredos, só criar branch e abrir PR |
| Humano no que é de alto risco | Aprovação humana pra ação de alto risco | Decisões 5 e 6: PR em rascunho, merge humano |
| Ataque simulado | Testes adversariais | Ainda não existe; entra pelo modelo de ameaças (#93) |

## O que este ADR ainda não cobre
A triagem reduz o risco, mas não zera:
- **O humano pode não ver a instrução.** A OWASP lembra que a entrada pode afetar o modelo mesmo sendo imperceptível pra uma pessoa: basta o modelo ler o conteúdo[^owasp]. A triagem pega o óbvio, não o escondido. Por isso o privilégio mínimo e o merge humano continuam valendo depois que a issue é aceita.
- **Issue não é a única porta.** Comentário em PR, arquivos de um PR de outra pessoa (incluindo um `AGENTS.md` alterado, que um agente revisor carregaria como regra), páginas lidas na pesquisa e documentação de dependências também são texto de fora. A recuperação de conteúdo apaga a fronteira entre dado e instrução[^greshake]. O Claude Code, por exemplo, pede aprovação antes de importar arquivos de fora da pasta do projeto, pra proteger de arquivos que outras pessoas fazem commit num projeto compartilhado[^cc-imports].

Cada porta nova, quando for tratada, vira ADR próprio citando este como base. O modelo de ameaças (#93) é onde elas são listadas.

## Consequências
- O tempo do mantenedor vai pra decisões curtas: aceitar na triagem e conferir na validação. Investigar e corrigir saem da mão dele nas issues de baixo risco.
- A rotina só liga depois dos testes do site no CI (#104). Menos supervisão pede mais teste[^hleg]; sem eles, a validação vira carimbo.
- Toda issue nova ganha um passo a mais (a triagem), mas é o passo que fecha a porta da injeção de prompt.
- Cada correção feita por IA registra o uso de tokens (ADR-0014).
- As quatro labels já existem no repositório. A aplicação automática de `precisa-triagem` e a rotina entram pela #111.

## Fontes
[^relatos]: Issues com a label `feedback` no repositório (quinze relatos do design system e dois do site) · https://github.com/vitorsantss/prumo/issues?q=label%3Afeedback
[^greshake]: Kai Greshake, Sahar Abdelnabi, Shailesh Mishra, Christoph Endres, Thorsten Holz e Mario Fritz, *Not what you've signed up for: Compromising Real-World LLM-Integrated Applications with Indirect Prompt Injection*, AISec '23 · https://arxiv.org/abs/2302.12173v2
[^owasp]: OWASP GenAI Security Project, *LLM01:2025 Prompt Injection* · https://genai.owasp.org/llmrisk/llm01-prompt-injection/
[^vies]: Raja Parasuraman e Dietrich Manzey, *Complacency and Bias in Human Use of Automation: An Attentional Integration*, Human Factors 52(3), 2010 · https://pubmed.ncbi.nlm.nih.gov/21077562/
[^hleg]: Grupo de Especialistas de Alto Nível em IA da Comissão Europeia, *Ethics Guidelines for Trustworthy AI* (2019), requisito 1, Human agency and oversight · https://ec.europa.eu/futurium/en/ai-alliance-consultation/guidelines/1.html
[^k8s-guia]: Kubernetes, *Issue Triage Guidelines* · https://kubernetes.dev/docs/guide/issue-triage
[^k8s-kep]: Kubernetes, *KEP-1553: Issue Triage Workflow and Automation* · https://www.k8s.dev/resources/keps/1553/
[^copilot]: GitHub Docs, *Risks and mitigations for GitHub Copilot cloud agent* · https://docs.github.com/en/enterprise-cloud@latest/copilot/concepts/agents/cloud-agent/risks-and-mitigations
[^cc-imports]: Anthropic, *Claude Code: Manage memory* (imports externos) · https://code.claude.com/docs/en/memory?r=0
