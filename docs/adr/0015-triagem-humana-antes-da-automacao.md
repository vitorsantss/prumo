# 0015 · Triagem humana antes de qualquer trabalho em issue, e automação só depois dela

- **Status:** Proposto
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
5. **A IA entrega um PR em rascunho, nunca o merge.** O PR traz a prévia, capturas do que mudou (claro, escuro e celular) e os testes passando. Se houver mais de um jeito de corrigir, as opções vêm lado a lado. A IA não aprova nem mescla o próprio trabalho. É o mesmo desenho do agente de código do GitHub: o PR em rascunho dele precisa ser revisado e mesclado por um humano, e o agente não consegue marcá-lo como pronto, aprovar nem mesclar[^copilot].
6. **Merge é humano no laço,** porque muda o projeto: uma pessoa decide cada um. A rotina que reproduz o problema e tira as capturas é humano sobre o laço: a pessoa desenha a rotina e acompanha o que ela faz, sem aprovar cada passo[^hleg]. A OWASP recomenda o mesmo: aprovação humana pra ação de alto risco[^owasp].
7. **O agente que trabalha em issue roda com o mínimo:** sem segredos, só com permissão pra criar branch e abrir PR, como recomenda a OWASP: privilégios do modelo restritos ao mínimo necessário[^owasp]. O modelo de ameaças (#93) detalha isso antes da rotina ligar.

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
