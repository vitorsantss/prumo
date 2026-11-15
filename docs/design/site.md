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
8. **Editar é um caminho, não um botão.** O rodapé de cada página mostra de qual arquivo ela vem e leva pra "Como mudar o site e a documentação", que explica o fluxo e aponta pra issue.

## Componentes (em `site/src/components/`)
| Componente | O que faz |
| --- | --- |
| `Busca` | Modal no meio da tela, abre com Ctrl K (⌘K no Mac). Usa o índice do Pagefind gerado no build, sem serviço externo. |
| `SeletorIdioma` | Popover com o nome de cada idioma escrito nele mesmo, sem bandeira, e o estado da tradução: original, traduzida por IA, revisada ou ainda sem tradução. Sem tradução, a opção leva pra "Ajude a traduzir". |
| `ChaveTema` | Botão de ícone; o tema novo se espalha num círculo a partir do botão. Guarda a escolha em `starlight-theme`, a mesma chave da documentação. |
| `starlight/*` | Substituem moldura da página, colunas, cabeçalho, rodapé da página, busca, idioma, tema e aviso de página sem tradução do Starlight. |

Tokens: `site/src/estilos/tokens.css` (valores de `docs/design/sistema/tokens.json`).

## Movimento
| O quê | Tempo | Como |
| --- | --- | --- |
| Tema | 420 ms | revelação circular a partir do botão (View Transitions API) |
| Idioma e troca de página | 180 + 240 ms | o texto some e volta com leve desfoque |
| Popover e busca | 240 ms | surge do ponto que abriu |

Tudo desligado com `prefers-reduced-motion`.

## Fontes
- Diátaxis, organização da documentação pela necessidade de quem lê · https://diataxis.fr
- Meng, Steinhardt e Schubert, How developers use API documentation (ACM SIGDOC) · https://sigdoc.acm.org/cdq/how-developers-use-api-documentation-an-observation-study/
- Seletor de idioma sem bandeira · https://www.smartling.com/blog/language-selector-best-practices
- View Transitions API · https://developer.mozilla.org/docs/Web/API/View_Transition_API
