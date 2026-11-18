# 0005 · Integrações terceirizadas, sempre com caminho nativo

- **Status:** Aceito
- **Data:** 2026-11-14

## Contexto

O Prumo conversa com muitos sistemas: rastreadores de tarefa, forjas de código (GitHub, GitLab, Forgejo), provedores de IA, comunicação. Vários são proprietários ou pagos. Ao mesmo tempo, o Prumo quer ser estudado, usado e melhorado por qualquer pessoa, inclusive sem dinheiro, e servir ao setor público ([ADR-0004](0004-licenca-agpl-3.md)).

Construir tudo em casa desde o início travaria a entrega. Depender de algo proprietário pra sempre trairia o propósito.

## Decisão

1. **Toda integração fica atrás de uma interface (porta)**, com implementações trocáveis (adapters[^cockburn], padrão Strategy[^gof]).
2. **Um adapter proprietário pode entrar primeiro**, como terceirização de uma etapa, pra acelerar a entrega. Ex.: executar etapas com o Claude Code no C1.
3. **Toda porta com adapter proprietário ganha, no mesmo dia, uma issue aberta para o caminho nativo** (livre e gratuito), com a label `caminho-nativo`. Essa issue é trabalho de primeira classe, não "um dia".
4. **Dependência obrigatória é sempre livre.** "Livre" aqui é a definição da FSF: a primeira liberdade é executar o programa pra qualquer propósito, e isso não tem a ver com preço[^fsf]. Então licença "grátis só pra uso não comercial" não é livre e não pode ser dependência obrigatória. Além de livre, a licença precisa ser compatível com a AGPL-3.0, conferida na lista de licenças da FSF[^fsf-lista]. O Prumo precisa rodar de ponta a ponta sem nada proprietário quando todas as issues `caminho-nativo` de uma etapa estiverem fechadas.
5. **Toda dependência nova passa por checagem de licença** antes de entrar (no CI, assim que houver CI).

## Consequências

- O C1 pode usar o Claude Code e chegar mais rápido; o caminho com agente e modelo livres (ex.: OpenCode ou Aider com modelo local) vira issue aberta, boa pra quem quer contribuir.
- A lista de issues `caminho-nativo` mostra, a qualquer momento, o quanto o Prumo ainda depende de terceiros.
- Cada interface precisa de testes que rodem com o adapter nativo ou com um dublê, pra não acoplar no proprietário sem perceber. É a intenção original do padrão: a aplicação pode ser usada por pessoas, programas, testes automatizados ou scripts, e ser desenvolvida e testada isolada dos dispositivos e bancos de dados de produção[^cockburn-intencao].
- Exige disciplina: adapter proprietário sem issue nativa não passa no review.

## Fontes
[^cockburn]: Alistair Cockburn, *Hexagonal Architecture* (Ports and Adapters), 2005 · https://alistair.cockburn.us/hexagonal-architecture/
[^cockburn-intencao]: Alistair Cockburn, *Hexagonal Architecture for Programmers*, que reproduz a frase de intenção do artigo original · https://alistaircockburn.com/Publications/Hexagonal-Architecture-for-Programmers
[^gof]: Gamma, Helm, Johnson e Vlissides, *Design Patterns: Elements of Reusable Object-Oriented Software* (1994), padrão Strategy · https://hillside.net/elements-of-reusable-object-oriented-software-book
[^fsf]: Free Software Foundation, *O que é o software livre?* · https://www.gnu.org/p/free-sw
[^fsf-lista]: Free Software Foundation, *Various Licenses and Comments about Them* · https://www.gnu.org/l/license-list.html
