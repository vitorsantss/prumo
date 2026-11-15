# 0008 · Merge só por squash

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
O repositório aceitava três tipos de merge: merge commit, squash e rebase. Com os três ao mesmo tempo, a `main` fica com um formato diferente a cada PR. O projeto espera contribuições de gente com níveis bem diferentes de experiência, e os commits intermediários de um PR ("wip", "corrige teste") não ajudam quem lê a história depois.

## Decisão
Todo PR entra na `main` por **squash**: um PR vira um commit. O título do commit é o título do PR, e o corpo leva as issues que ele fecha. Merge commit e rebase ficam desligados nas configurações, e o ruleset `proteger-main` exige histórico linear.

## Consequências
- A `main` fica linear: cada commit corresponde a uma entrega revisada, que pode ser revertida inteira com `git revert`.
- Quem contribui pode commitar à vontade na própria branch. Só o título e a descrição do PR precisam estar bons.
- As linhas `Co-Authored-By` dos commits do PR são preservadas no commit final, então o crédito de quem ajudou (inclusive IA) continua visível.
- PRs empilhados (um em cima do outro) dão conflito depois do squash do anterior. Quem empilhou refaz a base em cima da nova `main` com `git rebase --onto main <base-antiga>`.
- O histórico detalhado de cada PR continua disponível na página do PR no GitHub.
