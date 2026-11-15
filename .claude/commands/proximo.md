---
description: Mostra a próxima decisão do projeto, uma de cada vez
---
Você vai me ajudar a tocar o Prumo decidindo uma coisa por vez.

1. Leia, com a GitHub CLI (`gh`), nesta ordem:
   - o último comentário da issue aberta "Resumo diário" (label `automacao`), se existir;
   - os PRs abertos que esperam meu review;
   - as issues abertas do milestone com prazo mais próximo que não estão bloqueadas;
   - os ADRs em `docs/adr/` com status Proposto.
2. Monte uma fila curta do que precisa de mim, do mais urgente pro menos (prazo do checkpoint primeiro).
3. Me mostre **só o primeiro item**: o que é, por que importa agora, as opções que eu tenho, a sua recomendação com o porquê, e quanto tempo leva. Se for um PR, ofereça rodar `/revisar`.
4. Espere a minha resposta. Registre a decisão onde ela mora (comentário na issue, review no PR, ADR) e mostre o próximo item.
5. Se eu falar de uma ideia que não tem a ver com o item atual, anote como issue `ideia` (igual ao `/ideia`) e volte pro item.
6. Quando o checkpoint estiver garantido ou a fila acabar, diga e sugira parar.

$ARGUMENTS
