# Ajude a traduzir

Por que algumas páginas ainda não têm inglês, como saber o estado de cada uma e como pegar uma tarefa.

## Por que nem tudo está em inglês

O Prumo nasceu em português. As páginas de entrada (comece aqui, a ideia, estado atual, formas de contribuir e primeira contribuição) foram traduzidas por IA pra quem chega de fora conseguir entender a ideia. As decisões (ADRs) ainda não foram, porque mudam com mais frequência e precisam de revisão cuidadosa: um termo traduzido errado numa decisão muda o que ela diz.

## Como saber o estado de uma página

O seletor de idioma, no topo de cada página, mostra um destes estados pro inglês:

| Estado | O que quer dizer |
| --- | --- |
| tradução revisada | Uma pessoa leu e corrigiu. |
| traduzida por IA, aguardando revisão | Dá pra ler, mas ninguém conferiu ainda. |
| ainda sem tradução | Em inglês, a página fica no grupo "Only in Portuguese" da barra lateral, com o selo PT, e o link abre a versão em português. |

## Como ajudar

- **Revisar uma tradução feita por IA.** Leva uns vinte minutos e não exige programar: leia `docs/en/<página>.md` ao lado de `docs/<página>.md`, corrija o que estiver estranho e remova a linha "Translated by AI, waiting for a human review." do topo. É isso que muda o estado pra "tradução revisada".
- **Traduzir um ADR.** Uma tarefa por arquivo: comente na tarefa qual vai pegar, crie `docs/en/adr/<mesmo-nome>.md` e abra um PR. Se usar IA, comece a página com a mesma linha de aviso, pra outra pessoa revisar.

As tarefas abertas têm a etiqueta [`traducao`](https://github.com/vitorsantss/prumo/issues?q=is%3Aopen+label%3Atraducao). O passo a passo do primeiro PR está em [Primeira contribuição](primeira-contribuicao.md).
