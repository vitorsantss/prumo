# Como contribuir com o Prumo

O Prumo é feito pra ser estudado e construído junto. Estudante, professor, dev experiente ou alguém começando: tem lugar pra você.

## Por onde começar

1. Leia o [README](README.md) e os [ADRs](docs/adr/) pra entender o que já foi decidido e por quê.
2. Procure uma issue aberta:
   - `boa-primeira-issue`: pequena e bem explicada, boa pra começar.
   - `caminho-nativo`: trocar uma dependência paga ou proprietária por uma livre (ADR-0005).
   - `ajuda-bem-vinda`: precisa de alguém, não importa a experiência.
3. Comente na issue que vai pegar ela, pra ninguém fazer o mesmo trabalho em dobro.

Não achou uma issue que sirva? Abra uma com o modelo **Ideia** e conte o que pensou.

## Do seu jeito, com ou sem IA

Você pode contribuir escrevendo tudo na mão, usando a IA que já tem (paga ou gratuita) ou rodando o próprio Prumo. O que vale:

- **Diga como fez.** O modelo de PR pergunta se você usou IA e qual. Não é julgamento: é pra todo mundo aprender com o processo.
- **Você responde pelo que enviou.** Entenda cada linha do seu PR e saiba explicar.
- **Nada de dado sensível.** Nunca cole em PR, issue ou prompt dado de cliente, segredo ou código que não é seu.

## Fluxo

1. Faça um fork e crie um branch: `tipo/numero-descricao` (ex.: `feat/12-ler-pipeline-yml`).
2. Faça commits pequenos, com mensagem que diga o porquê.
3. Abra o PR ligando a issue (`Fecha #12`) e preencha o modelo.
4. Um mantenedor revisa. Pode haver pedido de ajuste, faz parte.

## Decisões grandes

Mudou algo de arquitetura (nova dependência, novo formato de arquivo, troca de padrão)? Escreva um ADR em `docs/adr/` usando o [modelo](docs/adr/0000-modelo.md) junto com o PR.

## Convivência

Todo mundo aqui segue o [Código de Conduta](CODE_OF_CONDUCT.md) (Contributor Covenant 2.1).

## Licença

Ao contribuir, você concorda que sua contribuição fica sob a [AGPL-3.0](LICENSE), a mesma do projeto.
