# Entrega

## tarefa.md

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

## spec.md

# Spec: Registrar transação e consultar extrato de um cliente

> **Nota sobre a fonte:** não consegui abrir a especificação original (https://github.com/zanfranceschi/rinha-de-backend-2024-q1) nesta execução, porque a ferramenta de leitura web não tinha permissão. Rotas, campos, códigos HTTP e limites abaixo vêm do que conheço da especificação pública da Rinha 2024/Q1. Ainda precisam ser conferidos na fonte (ver "Perguntas em aberto").

## Objetivo

Expor uma API HTTP de um banco fictício com cinco clientes já cadastrados. A API deve:

1. registrar transações de crédito e débito de um cliente;
2. devolver o extrato do cliente: saldo atual, limite e últimas transações.

Sob muitas requisições ao mesmo tempo, o saldo de um cliente nunca pode ficar abaixo do limite negativo dele.

## Regras de negócio

1. **Clientes pré-cadastrados.** A aplicação começa com estes cinco clientes, todos com saldo inicial `0`:

   | id | limite (centavos) |
   |----|-------------------|
   | 1  | 100000            |
   | 2  | 80000             |
   | 3  | 1000000           |
   | 4  | 10000000          |
   | 5  | 500000            |

   Não há cadastro de clientes pela API. Qualquer outro `id`, como o `6`, é "cliente inexistente".
2. **Valores em centavos.** `valor`, `saldo` e `limite` são inteiros em centavos.
3. **Crédito (`tipo = "c"`).** Soma `valor` ao saldo. Um crédito válido nunca é recusado por causa de saldo ou limite.
4. **Débito (`tipo = "d"`).** Subtrai `valor` do saldo. O débito é recusado se o novo saldo ficar abaixo de `-limite`, ou seja, se `saldo - valor < -limite`.
5. **Transação recusada não tem efeito.** Ela não altera o saldo e não aparece no extrato.
6. **Consistência sob concorrência.** O saldo final de um cliente é sempre a soma exata das transações aceitas, mesmo com requisições simultâneas para o mesmo cliente. Em nenhum momento o saldo gravado fica abaixo de `-limite`.
7. **Consistência entre os endpoints.** O saldo e o limite devolvidos ao registrar uma transação e os devolvidos pelo extrato vêm da mesma fonte de verdade.
8. **Extrato.** Lista no máximo as **10** transações mais recentes do cliente, da mais nova para a mais antiga.

## Contratos (entradas e saídas)

### 1. Registrar transação

`POST /clientes/{id}/transacoes`

**Corpo da requisição (JSON):**

```json
{
  "valor": 1000,
  "tipo": "c",
  "descricao": "descricao"
}
```

| Campo       | Tipo    | Validação                                         |
|-------------|---------|---------------------------------------------------|
| `valor`     | inteiro | obrigatório, positivo; sem casas decimais         |
| `tipo`      | string  | obrigatório; exatamente `"c"` ou `"d"`            |
| `descricao` | string  | obrigatório; de 1 a 10 caracteres                 |

**Respostas:**

| Situação | HTTP | Corpo |
|---|---|---|
| Transação aceita | `200` | `{"limite": <inteiro>, "saldo": <inteiro>}`, com o saldo **após** a transação |
| Cliente inexistente | `404` | qualquer corpo, ou nenhum |
| Débito que ultrapassaria o limite | `422` | qualquer corpo, ou nenhum |
| Payload inválido (campo ausente ou nulo, tipo errado, `valor` não inteiro ou ≤ 0, `tipo` diferente de `c`/`d`, `descricao` vazia ou com mais de 10 caracteres, JSON malformado) | `422` | qualquer corpo, ou nenhum |

Exemplo de resposta `200`:

```json
{ "limite": 100000, "saldo": -9098 }
```

### 2. Consultar extrato

`GET /clientes/{id}/extrato`

**Respostas:**

| Situação | HTTP | Corpo |
|---|---|---|
| Cliente existe | `200` | JSON abaixo |
| Cliente inexistente | `404` | qualquer corpo, ou nenhum |

Corpo `200`:

```json
{
  "saldo": {
    "total": -9098,
    "data_extrato": "2024-01-17T02:34:41.217753Z",
    "limite": 100000
  },
  "ultimas_transacoes": [
    {
      "valor": 10,
      "tipo": "c",
      "descricao": "descricao",
      "realizada_em": "2024-01-17T02:34:38.543030Z"
    }
  ]
}
```

| Campo | Significado |
|---|---|
| `saldo.total` | saldo atual do cliente, em centavos |
| `saldo.data_extrato` | data e hora da consulta, em ISO 8601 / RFC 3339 |
| `saldo.limite` | limite do cliente, em centavos |
| `ultimas_transacoes` | até 10 transações aceitas, ordenadas por `realizada_em` decrescente; lista vazia se não houver nenhuma |
| `ultimas_transacoes[].valor` | valor da transação, inteiro positivo em centavos |
| `ultimas_transacoes[].tipo` | `"c"` ou `"d"` |
| `ultimas_transacoes[].descricao` | descrição enviada na criação |
| `ultimas_transacoes[].realizada_em` | data e hora do registro, em ISO 8601 / RFC 3339 |

## Critérios de aceite verificáveis

1. Com o banco recém-iniciado, `GET /clientes/{1..5}/extrato` devolve `200`, com `saldo.total = 0`, `saldo.limite` igual ao da tabela de clientes e `ultimas_transacoes = []`.
2. `GET /clientes/6/extrato` devolve `404`.
3. `POST /clientes/6/transacoes` com um payload válido devolve `404`.
4. `POST /clientes/1/transacoes` com `{"valor": 1000, "tipo": "c", "descricao": "deposito"}` devolve `200` e `{"limite": 100000, "saldo": 1000}`.
5. A partir do saldo `1000` do cliente 1, um débito de `101000` devolve `200` com `saldo = -100000` (exatamente no limite). Um débito seguinte de `1` devolve `422`, e o saldo continua `-100000` no extrato.
6. Um débito recusado com `422` não aparece em `ultimas_transacoes`.
7. Cada um destes payloads devolve `422` e não altera o saldo:
   - `valor` ausente, nulo, `0`, negativo, decimal (`1.2`) ou string (`"10"`);
   - `tipo` ausente, nulo ou diferente de `"c"`/`"d"` (por exemplo `"x"`);
   - `descricao` ausente, nula, `""` ou com 11 ou mais caracteres;
   - JSON malformado.
8. Uma `descricao` com exatamente 10 caracteres é aceita (`200`).
9. Depois de 12 transações aceitas para um cliente, o extrato devolve exatamente 10, da mais recente para a mais antiga, e a primeira é a última registrada.
10. **Concorrência.** Disparar N débitos simultâneos de valor `v` para um cliente, em que a soma ultrapassa o limite. Ao final:
    - nenhuma resposta `200` traz `saldo < -limite`;
    - `saldo.total` no extrato é igual a `-(quantidade de respostas 200 × v)`;
    - a quantidade de respostas `422` é igual a `N − quantidade de respostas 200`.
11. **Concorrência mista.** Depois de créditos e débitos simultâneos, `saldo.total` é igual à soma algébrica dos valores das transações que receberam `200`.
12. Em toda resposta `200` de transação, o par `limite`/`saldo` é igual ao que um extrato feito logo depois devolve, desde que não haja outras escritas no meio.
13. `data_extrato` e `realizada_em` são strings de data e hora válidas em ISO 8601.

## Fora do escopo

- Criar, editar ou remover clientes pela API.
- Autenticação e autorização.
- Paginação ou filtro do extrato além das 10 transações mais recentes.
- Estorno ou cancelamento de transações.
- Infraestrutura exigida pela competição (docker-compose, balanceador de carga, duas instâncias da API, porta 9999, limites de CPU e memória) e o teste de carga oficial com Gatling. Ficam para tarefas futuras do caso de exemplo, se forem previstas.
- Escolha de linguagem, framework e banco de dados.

## Perguntas em aberto

1. **Conferir a fonte.** Rotas, nomes de campos, o limite de 10 caracteres da `descricao`, o máximo de 10 transações no extrato, os limites dos clientes e os códigos HTTP vieram do que conheço da especificação pública, sem leitura da fonte nesta execução. Alguém precisa conferir no repositório original antes da implementação.
2. **Payload inválido: `422` ou `400`?** Esta spec usa `422`. Pelo que conheço, os testes oficiais também aceitam `400`. Vale padronizar em `422` ou aceitar os dois?
3. **ID de cliente não numérico** (por exemplo `/clientes/abc/extrato`). Deve responder `404` ou `400`/`422`?
4. **Unicode na `descricao`.** O limite de 10 conta caracteres (code points) ou bytes?
5. **Corpo das respostas de erro** (`404` e `422`). Deve ter um formato padronizado ou pode ser vazio?
6. **Precisão e fuso das datas.** Quantas casas de fração de segundo, e é obrigatório usar UTC com sufixo `Z`?
7. **Desempate na ordenação.** Como ordenar transações com o mesmo `realizada_em`, por exemplo pela ordem de inserção?
8. **Limite superior de `valor`.** Existe um teto, ou um risco de overflow a tratar, para `valor` e para o saldo acumulado por créditos?

