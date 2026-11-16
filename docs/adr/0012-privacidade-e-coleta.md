# 0012 · O que o Prumo coleta de quem usa

- **Status:** Proposto
- **Data:** 2026-11-15

## Contexto
A frente de estudo de observabilidade (#66) quer medir uso, erros e bugs do próprio Prumo. Essas ferramentas vão de uma contagem anônima de visitas até gravar a sessão inteira da pessoa, e quanto mais perto da sessão, mais dado pessoal elas tocam. O ADR-0003 já proíbe dado pessoal no código, nos prompts e no que a IA vê. A LGPD exige finalidade, necessidade e transparência pra qualquer tratamento de dado pessoal[^lgpd].

Precisamos de uma regra antes da primeira coleta, não depois.

## Decisão
1. **Hoje o site coleta só contagem anônima de visitas**, com o Cloudflare Web Analytics, que não usa cookie nem armazenamento no navegador pra identificar pessoas[^cwa]. Liga no painel do Pages (Metrics, Web Analytics), sem mudar código.
2. **Qualquer coleta além disso precisa de um ADR próprio** antes de entrar: erros com contexto, relato de bug com dados técnicos, replay de sessão. O ADR diz o que é coletado, pra quê, por quanto tempo, quem acessa e como a pessoa recusa.
3. **Por padrão, nada que identifique a pessoa.** Quando uma coleta tocar dado pessoal, o padrão é mascarar tudo e liberar campo por campo, nunca o contrário.
4. **Tudo à vista:** a página `docs/privacidade.md` lista o que é coletado hoje e muda no mesmo PR que mudar a coleta.
5. **Caminho nativo:** o Cloudflare Web Analytics é serviço proprietário; a alternativa livre está na #68 (ADR-0005).

## Consequências
- O site ganha números de visita sem banner de consentimento, porque não identifica ninguém.
- Cada experimento de observabilidade da #66 começa por um ADR, o que deixa o estudo mais lento e mais ensinável: a decisão de privacidade fica documentada junto do código.
- O Lighthouse continua sendo a régua de desempenho: um script de coleta não pode derrubar as notas abaixo de 90.

## Fontes
[^lgpd]: Lei nº 13.709/2018 (LGPD), art. 6º · https://www.planalto.gov.br/ccivil_03/_ato2015-2018/2018/lei/l13709.htm
[^cwa]: Cloudflare, *Enable Web Analytics* · https://developers.cloudflare.com/pages/how-to/web-analytics
