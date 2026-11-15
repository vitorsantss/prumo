---
description: Me guia pela revisão de um PR
argument-hint: <número do PR>
---
Me ajude a revisar o PR #$ARGUMENTS do Prumo. Eu aprovo só o que entendo.

1. Leia o PR com `gh pr view` e `gh pr diff`, e as issues que ele fecha.
2. Me diga em 3 linhas o que ele muda e por quê.
3. Proponha uma ordem de leitura dos arquivos (do que define pro que usa) e me leve por ela, um arquivo por vez: o que o arquivo faz, os trechos que merecem atenção e o porquê das escolhas, ligando aos ADRs.
4. A cada arquivo, pergunte se ficou alguma dúvida. Dúvida minha que apontar um problema vira comentário no PR, na linha certa, com `gh`.
5. No fim, resuma: o que está bom, o que precisa mudar e se você recomenda aprovar. A decisão de merge é minha.
