# Passar no teste de carga oficial da Rinha

Tarefa 3 do caso de exemplo: a Rinha de Backend, edição 2024/Q1, um desafio público brasileiro com licença MIT.
Especificação original: https://github.com/zanfranceschi/rinha-de-backend-2024-q1

Depende da tarefa 2: a solução já sobe no formato da Rinha.

## Contexto (resumo próprio)

O teste oficial dispara muitas requisições ao mesmo tempo e, além de medir desempenho, confere se as regras de negócio
continuam valendo sob carga. É a prova de que a concorrência foi resolvida de verdade.

## O que entregar nesta tarefa

- O teste de carga oficial rodando contra a solução até o fim.
- Todas as validações de negócio do teste passando.
- O resultado do teste guardado junto da entrega.

## Dados da especificação original

- Ferramenta: Gatling, com a simulação `RinhaBackendCrebitosSimulation.scala`, rodada pelo `executar-teste-local.sh` (ou `.ps1` no Windows).
- Requisitos para rodar: JDK de 64 bits (11, 17 ou 21) e a variável `GATLING_HOME`.
- Resultados: gravados em `./load-test/user-files/results`.
- O que o teste confere além do desempenho: débito nunca leva o saldo abaixo do limite (senão 422), transação válida responde 200, corpo inválido responde 422 e cliente inexistente (id 6) responde 404.
