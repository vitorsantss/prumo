# Como mudar o site e a documentação

Qualquer pessoa pode propor mudança na página inicial, na documentação ou no visual. O caminho é o mesmo que o Prumo defende pra software: a ideia é escrita e aprovada antes do código.

## Só texto da documentação
Corrigiu um erro, deixou algo mais claro, acrescentou uma fonte? O rodapé de cada página diz de qual arquivo ela vem (por exemplo, `docs/a-ideia.md`). Abra esse arquivo [no repositório](https://github.com/vitorsantss/prumo/tree/main/docs), edite pelo próprio GitHub e abra o PR direto, sem issue.

Prefere só avisar? [Abra uma issue](https://github.com/vitorsantss/prumo/issues/new?template=site.yml) dizendo a página e o que está errado.

## Conteúdo da página inicial
Hipóteses, linha do tempo e textos da página inicial ficam em `site/src/data/landing.json`, em português e inglês. Mudou o estado do projeto (uma hipótese foi testada, um marco foi cumprido)? Atualize esse arquivo no mesmo PR da mudança. Conferir isso automaticamente é a issue #50.

## Visual, componente ou navegação
1. **Abra uma issue** com o modelo ["Mudança no site ou na documentação"](https://github.com/vitorsantss/prumo/issues/new?template=site.yml): o problema que você viu, a proposta e quem é afetado.
2. **Mostre como fica**, nos temas claro e escuro. Vale print, desenho no Excalidraw ou protótipo; o importante é dar pra comparar antes e depois.
3. **Confira as regras** da [camada Site do design system](https://github.com/vitorsantss/prumo/blob/main/docs/design/site.md). Se a proposta quebra uma regra, diga qual e por quê: às vezes a regra é que precisa mudar.
4. **Espere a aprovação na issue** antes de programar. É o portão humano: evita PR grande recusado por causa do desenho.
5. **Abra o PR** com:
   - o código;
   - `docs/design/site.md` atualizado, se surgiu regra ou componente novo;
   - um ADR, se entrou dependência ou padrão novo;
   - prints nos temas claro e escuro, em celular, notebook e tela larga;
   - `cd site && npm run build` passando.
6. **Revisão e merge por squash**, como todo PR (ADR-0008).

## Testar no seu computador
```
cd site
npm install
npm run dev
```
A busca só funciona depois de `npm run build && npm run preview`, porque o índice é gerado no build.
