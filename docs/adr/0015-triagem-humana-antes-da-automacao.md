# 0015 · Triagem humana antes de qualquer trabalho em issue, e automação só depois dela

- **Status:** Proposto
- **Data:** 2026-11-16

## Contexto
Relatos e sugestões chegam sempre (#75 a #89, #99, #100), e corrigi-los à mão estava tirando o mantenedor da entrega principal. A ideia é deixar uma IA corrigir os de baixo risco e manter o humano só onde ele decide.

Dois riscos aparecem logo:
- **Issue pública é texto de qualquer pessoa.** A injeção indireta de prompt acontece quando um agente processa texto de um atacante, por exemplo numa issue do GitHub, e segue instruções escondidas nele[^injecao]. A OWASP põe injeção de prompt em primeiro lugar na lista de riscos de aplicações com LLM[^owasp] e afirma que hoje não existe mecanismo confiável de prevenção[^owasp-2026]. Um agente que lê relato público antes de um humano é exatamente essa porta.
- **Revisão sem base vira carimbo.** Com o tempo, quem revisa tende a confiar na recomendação da IA quando ela costuma acertar, e a supervisão vira aprovação automática[^vies].

A literatura de supervisão separa dois papéis. Humano **no** laço: o sistema espera a pessoa aprovar[^hitl]. Humano **sobre** o laço: a pessoa supervisiona e pode intervir sem que o sistema espere por ela[^hotl]. Decisão de alto impacto ou irreversível pede humano no laço; tarefa de baixo risco e alto volume pode ter só supervisão[^hitl].

## Decisão
1. **Toda issue nova entra com a label `precisa-triagem`.** Ninguém trabalha nela, nem humano nem IA, antes da triagem.
2. **A triagem é humana.** A pessoa lê a issue e decide:
   - `triagem/aceita`: faz sentido e está pronta pra ser trabalhada;
   - `triagem/precisa-informacao`: falta informação, então comenta pedindo ou fecha;
   - fechar: não faz sentido, é duplicada ou está no lugar errado.

   É o fluxo do Kubernetes: issue nova recebe `needs-triage` automaticamente e, depois de avaliada, `triage/accepted`, que indica que está pronta pra ser trabalhada[^k8s-guia][^k8s-kep].
3. **Aceita, a issue fica disponível pra humano ou IA.** A label `triagem/aceita` também mostra a quem quer contribuir o que está pronto pra pegar.
4. **A IA só trabalha sozinha em issue com `pode-ser-automatica`,** label que a pessoa põe na triagem, junto com `triagem/aceita`, quando o risco é baixo: ajuste visual, texto, link quebrado. Assim, nenhuma IA processa texto público que um humano ainda não leu.
5. **A IA entrega um PR em rascunho, nunca o merge.** O PR traz a prévia, capturas do que mudou (claro, escuro e celular) e os testes passando. Se houver mais de um jeito de corrigir, as opções vêm lado a lado. A IA não aprova nem mescla o próprio trabalho, o mesmo desenho do agente de código do GitHub: o trabalho vai pra uma branch separada e um PR em rascunho, sem publicar nada sem revisão humana[^copilot-limites][^copilot-issue].
6. **Merge é humano no laço,** porque muda o projeto. Triagem automática de segundo nível, reprodução e capturas são humano sobre o laço.
7. **O agente que trabalha em issue roda com o mínimo:** sem segredos, só com permissão pra criar branch e abrir PR. O modelo de ameaças (#93) detalha isso antes da rotina ligar.

## Consequências
- O tempo do mantenedor vai pra decisões curtas: aceitar na triagem e conferir na validação. Investigar e corrigir saem da mão dele nas issues de baixo risco.
- A rotina só liga depois dos testes do site no CI (#104). Sem eles, a validação vira carimbo.
- Toda issue nova ganha um passo a mais (a triagem), mas é o passo que fecha a porta da injeção de prompt.
- Cada correção feita por IA registra o uso de tokens (ADR-0014).
- As quatro labels já existem no repositório. A aplicação automática de `precisa-triagem` e a rotina entram pela #111.

## Fontes
[^injecao]: Context7, *Indirect prompt injection examples: how coding agents get hijacked through docs, issues and search* · https://context7.com/guides/prompt-injection-coding-agents
[^owasp]: OWASP, *LLM01:2025 Prompt Injection* · https://genai.owasp.org/llmrisk/llm01-prompt-injection/
[^owasp-2026]: Techplained, *AI Agent Security Risks (2026): 10 Risks Mapped to OWASP* · https://www.techplained.com/ai-agent-security-risks
[^vies]: Forkast, *Human-in-the-loop (HITL)* · https://forkast.news/glossary/human-in-the-loop-hitl/
[^hitl]: CASRAI, *Human in the loop* · https://casrai.org/ai-governance/human-in-the-loop
[^hotl]: Arize, *What Is Human-On-The-Loop?* · https://arize.com/glossary/human-on-the-loop/
[^k8s-guia]: Kubernetes, *Issue Triage Guidelines* · https://kubernetes.dev/docs/guide/issue-triage
[^k8s-kep]: Kubernetes, *KEP-1553: Issue Triage Workflow and Automation* · https://www.k8s.dev/resources/keps/1553/
[^copilot-limites]: TechRepublic, *GitHub Copilot's New AI Coding Agent Saves Developers Time – And Requires Their Oversight* · https://www.techrepublic.com/article/news-github-copilot-ai-coding-agent/
[^copilot-issue]: GitHub Docs, *Using GitHub Copilot to work on an issue* · https://docs.github.com/en/copilot/using-github-copilot/coding-agent/using-copilot-to-work-on-an-issue
