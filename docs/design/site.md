# Camada Site do design system

O design system do Prumo nasceu pro canvas do app. Esta camada aplica as mesmas regras à página inicial e à documentação. Quem muda o visual do site muda este arquivo no mesmo PR.

## Regras
1. **O canvas é a tela.** Fundo com a grade de pontos. Cabeçalho, navegação e índices são painéis que flutuam (`.p-ilha`); texto longo fica numa folha opaca (`.p-folha`), porque a grade atrás do texto atrapalha a leitura.
2. **Destaque é a cor do texto.** Links, item ativo e botão primário usam `--p-ink`. Não existe cor de destaque decorativa.
3. **Cor só diz estado, e o estado muda de forma também.**

   | Estado | Forma | Token |
   | --- | --- | --- |
   | Não iniciado | ponto vazado | `--p-todo` |
   | Em andamento | anel com miolo | `--p-run` |
   | Pede sua ação | cheio com anel, único laranja da tela | `--p-wait` |
   | Feito | ponto cheio | `--p-done` |

4. **Sem pílula nem barra segmentada.** Idioma é um popover; tema é um botão de ícone.
5. **Manuscrito só em destaque.** Architects Daughter no logo, títulos e no adesivo "em estudo". Texto corrido em Inter, código em JetBrains Mono.
6. **Sempre claro e escuro.** Toda mudança visual é conferida nos dois temas.
7. **Três larguras.** Confira em celular (390 px), notebook (1440 px) e tela larga (2560 px). O conteúdo fica num contêiner centralizado de até 1400 px; nada estica até a borda.
8. **Idioma sempre declarado.** A interface segue o idioma escolhido; o conteúdo diz o dele. Página sem tradução continua no endereço em inglês, com a barra lateral e o cabeçalho em inglês, o adesivo **PT** (mesmo estilo do "em estudo": algo que dá pra fazer) e um aviso em inglês com link pra "Ajude a traduzir". O texto em português vai num bloco com `lang="pt-BR"`, pra leitor de tela pronunciar certo. Tradução por IA ganha o aviso tracejado no topo. Base: WCAG 3.1.2 e a orientação do W3C de marcar o idioma de links.
9. **Editar é um caminho, não um botão.** O rodapé de cada página mostra de qual arquivo ela vem e leva pra "Como mudar o site e a documentação", que explica o fluxo e aponta pra issue.
10. **Rolagem discreta e sem rolagem lateral.** A barra de rolagem da página é fina, com o polegar na cor `--p-line` e trilho transparente, nos dois temas; painéis internos (navegação lateral) rolam sem barra visível. Nenhuma página rola pro lado: nada de `100vw` em elemento dentro da página, porque em sistemas com barra clássica (Windows) `100vw` inclui a largura da barra e cria rolagem horizontal. Tabela larga quebra o texto nas células; no celular, empilha.

## Componentes

Ficam em `site/src/components/`.

| Componente | O que faz |
| --- | --- |
| `Busca` | Modal no meio da tela, abre com Ctrl K (⌘K no Mac). Usa o índice do Pagefind gerado no build, sem serviço externo. |
| `SeletorIdioma` | Popover com o nome de cada idioma escrito nele mesmo, sem bandeira, e o estado da tradução: original, traduzida por IA, revisada ou ainda sem tradução. Sem tradução, a opção leva pra "Ajude a traduzir". |
| `ChaveTema` | Botão de ícone; o tema novo se espalha num círculo a partir do botão. Guarda a escolha em `starlight-theme`, a mesma chave da documentação. |
| `starlight/*` | Substituem moldura da página, colunas, cabeçalho, rodapé da página, busca, idioma, tema e aviso de página sem tradução do Starlight. |

Tokens: `site/src/estilos/tokens.css` (valores de `docs/design/tokens.json`).

## Movimento
| O quê | Tempo | Como |
| --- | --- | --- |
| Tema | 420 ms | revelação circular a partir do botão (View Transitions API) |
| Idioma e troca de página | 120 + 180 ms | primeiro o conteúdo some, depois o novo aparece; cabeçalho e menus ficam parados |
| Popover e busca | 240 ms | surge do ponto que abriu |

Tudo desligado com `prefers-reduced-motion`.

A troca de página só começa quando o menu lateral, o título e o índice da página nova já chegaram (`<link rel="expect" blocking="render">` em `site/astro.config.mjs`). Sem isso, em rede lenta, a animação mostra a página nova ainda pela metade, com o menu vazio.

## Fontes

- Diátaxis, organização da documentação pela necessidade de quem lê · https://diataxis.fr
- Meng, Steinhardt e Schubert, How developers use API documentation (ACM SIGDOC) · https://sigdoc.acm.org/cdq/how-developers-use-api-documentation-an-observation-study/
- Seletor de idioma sem bandeira · https://www.smartling.com/blog/language-selector-best-practices
- View Transitions API · https://developer.mozilla.org/docs/Web/API/View_Transition_API
- WCAG 2.2, critério 3.1.2 Language of Parts · https://www.w3.org/WAI/WCAG22/Understanding/language-of-parts
- W3C, idioma em links de navegação · https://w3c.studio24.net/docs/internationalization-approach-agreed/
- Chrome for Developers, Consistent cross-document transitions · https://developer.chrome.com/docs/modern-web-guidance/guides/ui-behaviors/consistent-cross-document-transitions
