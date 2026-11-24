# Rinha de Backend 2024/Q1

Primeiro caso de estudo do Prumo: a [Rinha de Backend 2024/Q1](https://github.com/zanfranceschi/rinha-de-backend-2024-q1), um desafio público brasileiro, com licença MIT.
Uma API de banco fictício que registra créditos e débitos e mostra o extrato, sem deixar nenhum saldo passar do limite, mesmo com muita requisição ao mesmo tempo.
Foi escolhido por ser pequeno, conhecido e ter teste oficial: se algo der errado, o problema está no Prumo, não no projeto.

## Tarefas

| | Tarefa | O que entrega |
| --- | --- | --- |
| 1 | [Registrar transação e consultar extrato](1-transacao-e-extrato.md) | As regras do domínio e as duas rotas, sem infraestrutura |
| 2 | [Subir no formato da Rinha](2-formato-da-rinha.md) | Load balancer, 2 instâncias da API e banco, num `docker-compose.yml`, dentro do limite de recursos |
| 3 | [Passar no teste de carga oficial](3-teste-de-carga.md) | O teste roda até o fim e as validações de negócio passam |

Cada tarefa traz os números da especificação original, com o link, pra etapa SPEC não precisar completar de memória.
