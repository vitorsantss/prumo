# Pruminhos

Um pruminho é um mini-projeto de estudo que nasce dentro do Prumo: uma biblioteca, um gate, uma ferramenta ou uma pesquisa curta. O Prumo serve de cobaia. Se o pruminho for testado, medido e documentado, ele pode ser usado por outros projetos com confiança.

A ideia é que o Prumo vire um ecossistema: o projeto principal e, em volta dele, várias frentes abertas que qualquer pessoa pode continuar.

## O que todo pruminho tem
- **Uma pergunta.** O que queremos descobrir, numa frase. Ex.: "um gate automático de acessibilidade pega problema antes do merge sem atrapalhar?"
- **Fontes.** O que a literatura e o mercado já dizem, com link, seguindo [Como escrever](como-mudar-o-site.md).
- **Um experimento no Prumo.** Onde e como foi testado aqui, e o que foi medido.
- **Resultado.** Funcionou ou não, com os números. Resultado negativo também vale: evita que alguém repita o erro.
- **Documentação que ensina.** Como reproduzir, como usar e por que cada escolha foi feita, pra servir de material de estudo.
- **Onde dá pra inovar.** O que as fontes não resolvem e o pruminho tenta resolver.

## Ciclo de vida
1. **Ideia:** uma issue com a label `pruminho` e a pergunta.
2. **Estudo:** fontes reunidas e comparadas na própria issue.
3. **Experimento:** o pruminho roda no Prumo, num PR como qualquer outro. Se ele coleta dado de quem usa, antes vem um ADR (ADR-0012).
4. **Resultado:** os números e a conclusão entram na documentação.
5. **Confiável** ou **arquivado.** Confiável é o pruminho testado em pelo menos dois projetos, com documentação completa e licença livre (ADR-0004). Arquivado é o que não funcionou: fica documentado o porquê.

## Como um pruminho vira parte do Prumo
- **Gate:** uma verificação que uma etapa pode exigir. Ex.: o gate de qualidade web na etapa de teste.
- **Etapa ou ferramenta de etapa:** algo que roda dentro do fluxo.
- **Ferramenta separada:** algo que vive fora do Prumo, mas no mesmo ecossistema. Ex.: um monitor de padrões.

## Frentes abertas
| Pruminho | Pergunta | Issue |
| --- | --- | --- |
| Observar o uso | Dá pra medir uso, erros e bugs sem tocar dado pessoal? | #66 |
| Gate de qualidade web | Um gate de desempenho, acessibilidade e SEO ajuda sem atrapalhar? | #70 |
| Monitor de padrões | Padrões com fonte viram verificações que geram tarefas úteis? | #71 |

Quer começar um? Abra uma issue com a label `pruminho` e a pergunta. O resto vem na conversa.
