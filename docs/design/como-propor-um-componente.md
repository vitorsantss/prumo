# Como propor um componente

Componente novo ou mudança visual só entra com a página do componente atualizada (ADR-0011). É o mesmo fluxo de [Como mudar o site](../como-mudar-o-site.md), com uma página a mais.

## Passo a passo
1. **Abra uma issue** com o modelo "Mudança no site ou na documentação": o problema, a proposta e quem é afetado.
2. **Mostre como fica**, claro e escuro, em celular, notebook e tela larga. Vale print, desenho ou protótipo.
3. **Espere a aprovação na issue** antes de programar.
4. **Abra o PR** com:
   - o código, usando só tokens (`var(--p-…)`); token novo entra em [`tokens.json`](tokens.json);
   - a página do componente em `docs/design/componentes/`, seguindo o molde abaixo;
   - `npm run auditar` sem nota abaixo de 90.

## O molde
Copie esta estrutura. As seções são sempre as mesmas, na mesma ordem, pra pessoa e IA acharem tudo no mesmo lugar.

```markdown
# Nome do componente

**Produto:** site | app | os dois · **Estado:** em estudo | estável

Pra que serve, em uma ou duas frases.

## Prévia
(HTML com o componente dentro de .p-tema-claro e .p-tema-escuro)

## Anatomia
## Estados
## Em tela pequena
## Acessibilidade
## Tokens que usa
## Código
<!-- codigo:caminho/do/arquivo -->
## Implementações
```

O marcador `<!-- codigo:... -->` é trocado no build pelo trecho do arquivo, então o código mostrado é sempre o que roda.

## Implementação em outro framework
Quer fazer a versão em React, Angular ou outro? Siga o molde, use os tokens gerados de `tokens.json` e acrescente uma linha em "Implementações" do componente com o caminho. Os trechos de código das páginas de componente são CC0; o resto do repositório é AGPL-3.0 (ADR-0004).
