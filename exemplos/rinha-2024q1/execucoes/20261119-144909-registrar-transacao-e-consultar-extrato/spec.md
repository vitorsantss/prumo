# Spec: Registrar transação e consultar extrato de um cliente

> Fonte: tarefa.md, a especificação original da Rinha de Backend 2024/Q1 (https://github.com/zanfranceschi/rinha-de-backend-2024-q1) e as respostas da revisão anterior (P1, P2 e P4). As rotas, o JSON, os códigos HTTP e os números de infraestrutura vêm da especificação original. Nesta etapa eles não foram conferidos de novo contra o repositório. Releia o README original antes de implementar.

## Objetivo

Entregar uma API HTTP de um banco fictício com duas operações:

1. Registrar uma transação de crédito ou de débito para um cliente já cadastrado.
2. Consultar o extrato desse cliente: saldo atual, limite e últimas transações.

O saldo nunca pode ficar abaixo do limite negativo do cliente, mesmo com muitas requisições ao mesmo tempo. A entrega já vem empacotada na infraestrutura exigida pela Rinha: load balancer, duas instâncias da API e um banco de dados, tudo em docker-compose e dentro dos limites de recursos.

## Regras de negócio

- **RN1. Clientes pré-cadastrados.** Existem 5 clientes. Todos começam com saldo inicial `0`:

  | id | limite   |
  |----|----------|
  | 1  | 100000   |
  | 2  | 80000    |
  | 3  | 1000000  |
  | 4  | 10000000 |
  | 5  | 500000   |

  Nenhum outro id existe. Por exemplo, o id `6` não pode existir.
- **RN2. Valores em centavos.** `valor`, `saldo` e `limite` são inteiros em centavos.
- **RN3. Crédito.** Uma transação de tipo `c` soma `valor` ao saldo. Crédito nunca é recusado por causa do limite.
- **RN4. Débito.** Uma transação de tipo `d` subtrai `valor` do saldo. O débito só é aceito se o novo saldo continuar maior ou igual a `-limite`.
- **RN5. Débito recusado.** Se o débito deixaria o saldo abaixo de `-limite`, a API recusa a transação com `422`. Nada é gravado e o saldo não muda.
- **RN6. Cliente inexistente.** Qualquer operação com um id que não existe responde `404`. Nada é gravado.
- **RN7. Consistência sob concorrência.** Requisições simultâneas para o mesmo cliente, inclusive quando chegam por instâncias diferentes da API, não podem violar a RN4. Também não podem perder atualizações de saldo. O saldo final tem que ser igual à soma das transações aceitas. Para isso, a atualização do saldo e a gravação da transação acontecem numa única transação atômica do banco, com o saldo do cliente travado durante a operação (ver RN11).
- **RN8. Extrato.** O extrato mostra o saldo atual, o limite, a data e hora da consulta e as **10 transações mais recentes** do cliente, da mais nova para a mais antiga.
- **RN9. Validação de entrada.**
  - `valor`: inteiro positivo, maior que 0. Números com casas decimais são inválidos.
  - `tipo`: só `"c"` ou `"d"`.
  - `descricao`: string de 1 a 10 caracteres. Não pode ser nula nem vazia.

  Um payload inválido é recusado **sempre com `422`**, nunca com `400`. Nada é gravado.
- **RN10. Infraestrutura (requisito da Rinha, dentro do escopo).**
  - Um load balancer (nginx ou equivalente) distribui as requisições entre **2 instâncias** da API.
  - O load balancer expõe a API na **porta 9999**.
  - Toda a solução sobe com um único `docker-compose.yml`.
  - A soma dos limites de todos os serviços declarados no compose (load balancer, as 2 APIs e o banco) não pode passar de **1,5 CPU** e **550 MB** de memória.
  - As 2 instâncias da API não guardam estado entre si. Todo o estado fica no banco compartilhado.
  - As demais restrições do README original também valem (modo de rede, imagens públicas, arquitetura das imagens etc.). Confira a lista exata lá.
- **RN11. Banco de dados: PostgreSQL.** É a escolha que melhor atende à tarefa, pelos motivos abaixo:
  - Transações ACID com trava de linha (`SELECT ... FOR UPDATE` ou `UPDATE ... WHERE` condicional). Isso garante a RN4 e a RN7 mesmo com 2 instâncias escrevendo ao mesmo tempo.
  - É fácil de configurar para pouca memória, o que ajuda a caber nos limites da RN10.
  - Os 5 clientes da RN1 são criados por script de inicialização do banco (seed) quando o container sobe.

## Contratos (entradas e saídas)

### C1. Registrar transação

`POST /clientes/{id}/transacoes`

Corpo da requisição:
```json
{
  "valor": 1000,
  "tipo": "c",
  "descricao": "descricao"
}
```

Respostas:

| Situação | Status | Corpo |
|---|---|---|
| Transação aceita | `200 OK` | `{ "limite": 100000, "saldo": -9098 }`. `limite` é o limite do cliente. `saldo` é o saldo depois da transação. |
| Cliente inexistente | `404 Not Found` | A definir (ver P3) |
| Débito acima do limite (RN5) | `422 Unprocessable Entity` | A definir (ver P3) |
| Payload inválido (RN9) | `422 Unprocessable Entity` | A definir (ver P3) |

### C2. Consultar extrato

`GET /clientes/{id}/extrato`

Resposta `200 OK`:
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

- `saldo.total`: saldo atual.
- `saldo.data_extrato`: data e hora da consulta, em ISO-8601 UTC.
- `saldo.limite`: limite do cliente.
- `ultimas_transacoes`: no máximo 10 itens, da transação mais recente para a mais antiga. É uma lista vazia se o cliente não tem transações.
- `realizada_em`: data e hora em que a transação foi gravada, em ISO-8601 UTC.

Cliente inexistente: `404 Not Found`.

### C3. Implantação

- Entrada: o comando `docker compose up` na raiz do projeto.
- Saída: a API responde a C1 e C2 em `http://localhost:9999`, passando pelo load balancer, com os 5 clientes da RN1 já cadastrados e saldo `0`.

## Critérios de aceite verificáveis

1. **CA1.** `POST /clientes/1/transacoes` com `{"valor":1000,"tipo":"c","descricao":"teste"}` sobre o saldo inicial responde `200` com `{"limite":100000,"saldo":1000}`.
2. **CA2.** Com o cliente 1 em saldo `0`, um débito de `100000` responde `200` com `saldo: -100000`. Logo depois, um débito de `1` responde `422` e o saldo continua `-100000`.
3. **CA3.** `POST /clientes/6/transacoes` com payload válido responde `404`. `GET /clientes/6/extrato` também responde `404`.
4. **CA4.** Payloads inválidos respondem exatamente `422` (nunca `400`) e não mudam o saldo:
   - `valor` igual a `0`, negativo, decimal (`1.2`) ou ausente;
   - `tipo` diferente de `c` e `d`, ou ausente;
   - `descricao` nula, vazia ou com mais de 10 caracteres.
5. **CA5.** Depois de 12 transações aceitas, `GET /clientes/{id}/extrato` devolve exatamente 10 itens em `ultimas_transacoes`, da mais recente para a mais antiga. O primeiro item é a última transação feita.
6. **CA6.** Um cliente sem transações tem extrato com `saldo.total = 0`, o `limite` da tabela da RN1 e `ultimas_transacoes = []`.
7. **CA7.** O `saldo.total` do extrato é igual ao `saldo` devolvido pela última transação aceita.
8. **CA8. Concorrência.** Disparar N débitos simultâneos para o mesmo cliente pela porta 9999, de forma que caiam nas 2 instâncias e que juntos passem do limite. Resultado esperado:
   - o saldo final nunca fica abaixo de `-limite`;
   - o número de respostas `200` bate com o saldo final;
   - nenhuma transação recusada aparece no extrato.
9. **CA9.** Transações recusadas (`404` ou `422`) nunca aparecem no extrato.
10. **CA10. Subida.** Num ambiente limpo, `docker compose up` sobe o load balancer, as 2 instâncias da API e o PostgreSQL. Depois disso, `GET http://localhost:9999/clientes/1/extrato` responde `200` com saldo `0`.
11. **CA11. Recursos.** No `docker-compose.yml`, a soma dos limites de CPU de todos os serviços é ≤ `1.5` e a soma dos limites de memória é ≤ `550MB`.
12. **CA12. Balanceamento.** Com o compose no ar, requisições pela porta 9999 são atendidas pelas 2 instâncias. Dá para verificar pelos logs de cada container. Parar uma das duas não corrompe os dados já gravados.
13. **CA13. Teste de carga oficial.** O script Gatling do repositório original roda contra a solução até o fim e as validações de consistência dele (saldo, limite, extrato) passam sem falhas.

## Fora do escopo

- Cadastrar, alterar ou remover clientes, e mudar limites.
- Autenticação e autorização.
- Paginação do extrato além das 10 últimas transações.
- Outras rotas que não sejam as duas acima.
- Metas de desempenho além de passar no teste de carga oficial (CA13), como percentis de latência ou ranking.
- Publicar a submissão no repositório da Rinha.

## Perguntas em aberto

- **P3.** O corpo das respostas de erro (`404` e `422`) segue algum formato, ou pode ser vazio?
- **P5.** Um id que não é numérico (por exemplo `/clientes/abc/extrato`) deve responder `404` ou `422`?
- **P6.** Os dados precisam voltar ao estado inicial (saldo `0`, sem transações) a cada `docker compose up`, ou podem persistir entre execuções (volume do PostgreSQL)?
- **P7.** Existe meta de desempenho além de passar no teste de carga oficial (por exemplo, latência p99 máxima ou zero respostas com erro/timeout no Gatling)?