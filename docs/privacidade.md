# Privacidade

O que o Prumo coleta de quem visita o site, hoje. A regra está no ADR-0012: qualquer coleta nova precisa de decisão própria e muda esta página no mesmo PR.

## O que é coletado
| O quê | Pra quê | Como |
| --- | --- | --- |
| Contagem de visitas por página, origem do acesso e desempenho da página | Saber se a documentação é lida e se o site está rápido | Cloudflare Web Analytics, sem cookie de rastreamento e sem identificar a pessoa[^cwa] |

## O que não é coletado
- Nome, e-mail ou qualquer dado que identifique você, inclusive usar o seu endereço IP pra te reconhecer.
- Cookies de rastreamento ou identificador salvo no seu navegador.
- Gravação da sua sessão, cliques ou o que você digita.

A busca do site roda no seu navegador: o que você procura não sai da página.

## Preferências salvas no seu navegador
O tema (claro ou escuro) fica guardado no seu navegador, só pra lembrar a sua escolha. Não é enviado pra lugar nenhum.

## Dúvidas
Abra uma conversa em [Discussions](https://github.com/vitorsantss/prumo/discussions).

[^cwa]: Cloudflare, *Enable Web Analytics* · https://developers.cloudflare.com/pages/how-to/web-analytics
