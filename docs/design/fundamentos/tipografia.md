# Tipografia

Três famílias, cada uma com um papel:

| Família | Token | Papel |
| --- | --- | --- |
| Architects Daughter | `--p-fonte-manuscrita` | Só destaque: logo, títulos, adesivo. Nunca texto corrido. |
| Inter | `--p-fonte-texto` | Texto e interface. |
| JetBrains Mono | `--p-fonte-mono` | Código, tokens, teclas. |

As três são servidas pelo próprio site (`site/src/fontes/`), sob a licença SIL Open Font License 1.1. Nada de Google Fonts: carregar de lá manda o IP de quem lê pro Google, o que fere o ADR-0012[^lg-munique], e cria uma conexão a mais antes do texto aparecer. Só o subconjunto latino e só os pesos usados: Inter 400, 500, 600 e 700, Mono 400 e Architects 400.

Pra não piscar:
- **as fontes da primeira tela (Inter 400 e 600, Architects) são pré-carregadas** no `<head>` (`site/src/fontes/primeira-tela.ts`);
- **os arquivos com hash no nome ficam em cache por um ano** (`site/public/_headers`). Sem isso, o Cloudflare manda o navegador conferir cada fonte a cada página, e o texto aparece na fonte do sistema até a resposta chegar;
- **a Architects usa `font-display: optional`:** se não chegar a tempo, a página fica na fonte reserva até a próxima, em vez de trocar na frente de quem lê[^preload-optional].

[^preload-optional]: web.dev, *Prevent layout shifting and flashes of invisible text by preloading optional fonts* · https://web.dev/articles/preload-optional-fonts

[^lg-munique]: LG München I, sentença de 20/01/2022, processo 3 O 17493/20 · https://www.gesetze-bayern.de/Content/Pdf/Y-300-Z-BECKRS-B-2022-N-612?all=False

## Escala
Gerada a partir de [`tokens.json`](../tokens.json).

<!-- tokens:tipografia -->
