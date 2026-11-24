# Registrar transação e consultar extrato de um cliente

Primeira tarefa do caso de exemplo: a Rinha de Backend, edição 2024/Q1, um desafio público brasileiro.
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

Os detalhes exatos (rotas, formato do JSON, códigos HTTP) devem ser conferidos na especificação original.
