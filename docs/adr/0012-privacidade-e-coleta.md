# 0012 · O que o Prumo coleta de quem usa

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
A frente de estudo de observabilidade ([#66](https://github.com/vitorsantss/prumo/issues/66)) quer medir uso, erros e bugs do próprio Prumo. Essas ferramentas vão de uma contagem anônima de visitas até gravar a sessão inteira da pessoa, e quanto mais perto da sessão, mais dado pessoal elas tocam. O [ADR-0003](0003-dado-sensivel-fora-da-ia.md) já proíbe dado pessoal no código, nos prompts e no que a IA vê. A LGPD exige finalidade, necessidade e transparência pra qualquer tratamento de dado pessoal[^lgpd].

Precisamos de uma regra antes da primeira coleta, não depois.

## Decisão
1. **Hoje o site coleta só contagem anônima de visitas**, com o Cloudflare Web Analytics. Pela página do produto, ele não usa estado no navegador, como cookie ou localStorage, pra coletar métricas[^cwa]. É afirmação do fornecedor: uma análise independente aponta que a política de privacidade da Cloudflare diz que ela pode registrar o IP de quem acessa[^ctrl], e IP pode ser dado pessoal ([#119](https://github.com/vitorsantss/prumo/issues/119)). Por isso o caminho nativo (decisão 5) tem prioridade. Liga no painel do Pages (Metrics, Web Analytics), sem mudar código[^cwa-ativar].
2. **Qualquer coleta além disso precisa de um ADR próprio** antes de entrar: erros com contexto, relato de bug com dados técnicos, replay de sessão. O ADR diz o que é coletado, pra quê, por quanto tempo, quem acessa e como a pessoa recusa.
3. **Por padrão, nada que identifique a pessoa.** Quando uma coleta tocar dado pessoal, o padrão é mascarar tudo e liberar campo por campo, nunca o contrário.
4. **Tudo à vista:** a página `docs/privacidade.md` lista o que é coletado hoje e muda no mesmo PR que mudar a coleta.
5. **Caminho nativo:** o Cloudflare Web Analytics é serviço proprietário; a alternativa livre está na [#68](https://github.com/vitorsantss/prumo/issues/68) ([ADR-0005](0005-integracoes-terceirizadas-caminho-nativo.md)), que também cobre levar a contagem pra base do próprio Prumo.

## Consequências
- O site ganha números de visita sem banner de consentimento. No guia de cookies da ANPD, o banner existe pra pedir consentimento pra cookie não necessário[^anpd]; como esta coleta não usa cookie, essa lógica não se aplica. É leitura do projeto, não parecer jurídico.
- Sem cookie, o Prumo não sabe quantas pessoas únicas visitam nem quantas voltam. Nenhuma pergunta de pesquisa dele precisa disso hoje; se um dia precisar, entra por ADR (decisão 2), com o banner e o consentimento que vierem junto.
- Cada experimento de observabilidade da [#66](https://github.com/vitorsantss/prumo/issues/66) começa por um ADR, o que deixa o estudo mais lento e mais ensinável: a decisão de privacidade fica documentada junto do código.
- O Lighthouse continua sendo a régua de desempenho: um script de coleta não pode derrubar as notas abaixo de 90.

## Fontes
[^lgpd]: Lei nº 13.709/2018 (LGPD), art. 6º · https://www.planalto.gov.br/ccivil_03/_ato2015-2018/2018/lei/l13709.htm
[^cwa]: Cloudflare, *Free Website Analytics* (página do produto) · https://cloudflare.com/web-analytics
[^cwa-ativar]: Cloudflare, *Enable Web Analytics* · https://developers.cloudflare.com/pages/how-to/web-analytics
[^ctrl]: Ctrl blog, *A technical and privacy review of Cloudflare Web Analytics* · https://www.ctrl.blog/entry/review-cloudflare-analytics/
[^anpd]: ANPD, *Guia orientativo: Cookies e proteção de dados pessoais* (2022) · https://www.gov.br/anpd/pt-br/documentos-e-publicacoes/guia-orientativo-cookies-e-protecao-de-dados-pessoais.pdf
