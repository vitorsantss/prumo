# Subir a API no formato da Rinha

Tarefa 2 do caso de exemplo: a Rinha de Backend, edição 2024/Q1, um desafio público brasileiro com licença MIT.
Especificação original: https://github.com/zanfranceschi/rinha-de-backend-2024-q1

Depende da tarefa 1: a API já registra transações e mostra o extrato.

## Contexto (resumo próprio)

A Rinha não avalia só a API: ela exige uma montagem fixa, com duas cópias da aplicação atrás de um balanceador,
um banco compartilhado e um teto apertado de CPU e memória. As cópias não guardam estado entre si.

## O que entregar nesta tarefa

- Um `docker-compose.yml` que sobe tudo com um comando.
- Um load balancer que distribui as requisições em rodízio entre as instâncias. Não precisa ser Nginx.
- A API acessível pelo load balancer.

## Dados da especificação original

- Porta: o load balancer aceita requisições na **9999**.
- Instâncias: **2** servidores da API atrás do load balancer.
- Banco: relacional ou não relacional. Armazenamento só em memória, como Redis, não vale como banco principal.
- Recursos: todos os serviços somados ficam dentro de **1,5 CPU** e **550 MB** de memória, declarados em `deploy.resources.limits` no `docker-compose.yml`, com a memória em MB.
- Imagens: publicadas num registro público (como o Docker Hub), não só construídas na máquina local.
- Rede: o exemplo usa `bridge`; `host` também é permitido, desde que as portas não conflitem.
- Subida: antes do teste, `GET /clientes/1/extrato` precisa responder em até 40 segundos (consultado a cada 2 segundos).
