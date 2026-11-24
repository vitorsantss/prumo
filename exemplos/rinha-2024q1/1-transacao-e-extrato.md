# Registrar transação e consultar extrato de um cliente

Tarefa 1 do caso de exemplo: a Rinha de Backend, edição 2024/Q1, um desafio público brasileiro com licença MIT.
Especificação original: https://github.com/zanfranceschi/rinha-de-backend-2024-q1

## Contexto (resumo próprio)

Uma API de um banco fictício com cinco clientes já cadastrados, cada um com um limite de crédito e saldo inicial zero.
A API precisa aceitar transações de crédito e débito e mostrar o extrato, aguentando muitas requisições ao mesmo tempo
sem deixar o saldo de nenhum cliente passar do limite negativo dele.

## O que entregar nesta tarefa

- Endpoint para registrar uma transação de um cliente (crédito ou débito, com valor e descrição curta).
- Endpoint para consultar o extrato de um cliente (saldo atual, limite e últimas transações).
- Recusar débito que deixaria o saldo abaixo do limite negativo do cliente.
- Responder "não encontrado" para cliente inexistente.

## Dados da especificação original

Clientes cadastrados (valores em centavos):

| id | limite | saldo inicial |
| --- | --- | --- |
| 1 | 100000 | 0 |
| 2 | 80000 | 0 |
| 3 | 1000000 | 0 |
| 4 | 10000000 | 0 |
| 5 | 500000 | 0 |

O id 6 não pode existir: o teste oficial confere que ele responde 404.

`POST /clientes/{id}/transacoes`

- Corpo: `valor` (inteiro positivo, em centavos), `tipo` (`c` para crédito ou `d` para débito) e `descricao` (texto de 1 a 10 caracteres). Os três são obrigatórios.
- Sucesso: 200, com `limite` (o limite do cliente) e `saldo` (o saldo depois da transação).
- 422: débito que deixaria o saldo abaixo de `-limite`, ou corpo inválido. Para `valor` não inteiro, 422 ou 400.
- 404: cliente inexistente.

`GET /clientes/{id}/extrato`

- Sucesso: 200, com `saldo` (`total`, que é o saldo atual inteiro, `data_extrato` e `limite`) e `ultimas_transacoes` (até 10, da mais recente para a mais antiga, cada uma com `valor`, `tipo`, `descricao` e `realizada_em`).
- 404: cliente inexistente.

Infraestrutura e teste de carga ficam para as tarefas 2 e 3.
