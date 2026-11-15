# Primeira contribuição, passo a passo

Antes de tudo, leia o [Comece aqui](comece-aqui.md). Depois:

1. **Escolha uma tarefa** entre as [marcadas pra começar](https://github.com/vitorsantss/prumo/issues?q=is%3Aopen+label%3A%22good+first+issue%22). Leia a tarefa inteira: ela diz o que fazer e quando está pronta.
2. **Avise que vai pegar**, comentando na tarefa. Assim duas pessoas não fazem a mesma coisa.
3. **Faça um fork** do repositório no GitHub e clone o seu fork:
   ```sh
   git clone https://github.com/<seu-usuario>/prumo
   cd prumo
   ```
4. **Crie um branch** com um nome que diga o que é:
   ```sh
   git checkout -b docs/revisa-traducao-comece-aqui
   ```
5. **Faça a mudança.** Se mexer em código, rode os testes (precisa do Go 1.24 ou mais novo):
   ```sh
   go test ./...
   ```
6. **Faça o commit** com uma mensagem em português dizendo o porquê. Se usou IA, pode creditar com uma linha `Co-Authored-By`.
7. **Abra o PR** pro repositório original. Escreva `Closes #<número da tarefa>` na descrição (só as palavras em inglês fecham a tarefa sozinhas) e preencha o modelo, inclusive "Como foi feito".
8. **Acompanhe a revisão.** Um mantenedor comenta no PR; ajuste no mesmo branch. Aprovado, o merge é por squash: o título do PR vira o commit (ADR-0008).

> **Nota.** Hoje os testes automáticos do repositório estão parados por uma pendência da conta. Até voltarem, quem revisa roda os testes à mão.

Dúvidas em qualquer passo: pergunte na própria tarefa ou nas [conversas do repositório](https://github.com/vitorsantss/prumo/discussions).
