# Privacidade

O que o Prumo coleta de quem visita o site, hoje. A regra está no [ADR-0012](adr/0012-privacidade-e-coleta.md): qualquer coleta nova precisa de decisão própria e muda esta página no mesmo PR.

## O que é coletado
| O quê | Pra quê | Como |
| --- | --- | --- |
| Contagem de visitas por página, origem do acesso e desempenho da página | Saber se a documentação é lida e se o site está rápido | Cloudflare Web Analytics, que segundo a Cloudflare não usa cookie nem identifica a pessoa[^cwa] |

## O que não é coletado
- Nome, e-mail ou qualquer dado que identifique você, inclusive usar o seu endereço IP pra te reconhecer.
- Cookies de rastreamento ou identificador salvo no seu navegador.
- Gravação da sua sessão, cliques ou o que você digita.

A busca do site roda no seu navegador: o que você procura não sai da página. As fontes também vêm do próprio site, então abrir uma página não manda o seu endereço IP pra nenhum serviço de fontes.

## Preferências salvas no seu navegador
O tema (claro ou escuro) fica guardado no seu navegador, só pra lembrar a sua escolha. Não é enviado pra lugar nenhum.

## Dúvidas
Abra uma conversa em [Discussions](https://github.com/vitorsantss/prumo/discussions).

[^cwa]: Cloudflare, *Free Website Analytics* · https://cloudflare.com/web-analytics. É afirmação do fornecedor; a política de privacidade da Cloudflare diz que ela pode registrar o IP de quem acessa. Os detalhes e o caminho pra trocar por uma contagem própria estão no [ADR-0012](adr/0012-privacidade-e-coleta.md).
