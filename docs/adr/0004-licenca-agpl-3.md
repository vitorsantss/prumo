# 0004 · Licença do projeto: AGPL-3.0

- **Status:** Aceito
- **Data:** 2026-11-14

## Contexto

O Prumo existe pra três coisas: servir de referência pra quem quer estudar e praticar como se constrói software, receber contribuição de profissionais de qualquer área da tecnologia e ser útil pro setor público (universidades, institutos federais, órgãos) na criação de software que contribua pra sociedade e pro Brasil.

A vontade inicial era proibir o uso por empresas privadas. Duas restrições tornam isso incompatível com os objetivos acima:

1. **Uma licença que discrimina uso ou pessoas não é open source.** A Open Source Definition (itens 5 e 6) proíbe restringir o uso "em um negócio" ou por grupos de pessoas, e o FAQ da OSI confirma que todo software open source pode ter uso comercial.
2. **O setor público brasileiro exige licença livre.** A Instrução Normativa nº 01/2011 do Software Público Brasileiro define que o SPB adota licença livre, cuja primeira liberdade é executar o programa para qualquer propósito. Uma licença que bloqueia empresas deixaria o Prumo fora do Portal do Software Público.

Alternativas consideradas:

- **PolyForm Noncommercial 1.0.0:** libera explicitamente instituições de ensino, de pesquisa pública e governo, e bloqueia uso comercial. Rejeitada porque não é aprovada pela OSI, afasta contribuidores e impede a entrada no Portal do Software Público.
- **MIT / Apache-2.0:** permissivas. Rejeitadas porque permitem que alguém feche uma versão melhorada e não devolva nada à comunidade.
- **GPL-3.0:** obriga abrir o código só quando há distribuição. Rejeitada porque o Prumo tende a rodar como serviço web, e aí a GPL não obriga quem modifica a abrir o código.

## Decisão

O Prumo é licenciado sob a **GNU Affero General Public License v3.0 (AGPL-3.0-only)**.

A intenção do projeto (educação, estudo e uso público) fica declarada no README como propósito, sem força de cláusula de licença.

## Consequências

- O Prumo é software livre e open source, apto a entrar no Portal do Software Público Brasileiro e seguro juridicamente pra quem quiser estudar e contribuir.
- Quem modificar o Prumo e oferecer a versão modificada, inclusive como serviço pela rede, é obrigado a disponibilizar o código das mudanças sob a mesma licença. Melhorias voltam pra comunidade.
- Empresas podem usar o Prumo, inclusive internamente e sem custo. Essa é a parte do desejo inicial que a licença não cobre, aceita conscientemente em troca de ser livre.
- A licença vale pro código do Prumo, não pros softwares criados com ele: cada projeto gerado escolhe a própria licença.
- Contribuições externas entram sob a mesma licença (inbound = outbound). Se no futuro for preciso relicenciar, será necessário o acordo dos contribuidores.

## Fontes

- Open Source Definition (anotada), OSI · https://opensource.org/osd-annotated
- FAQ da Open Source Initiative · https://opensource.org/faq
- Instrução Normativa SLTI nº 01/2011 (Software Público Brasileiro) · https://softwarepublico.gov.br/social/spb/documentos/estudos/licencas-fase-1/in-01-spb-2011.pdf
- PolyForm Noncommercial 1.0.0 · https://polyformproject.org/licenses/noncommercial/1.0.0
- GNU AGPL v3.0 · https://www.gnu.org/licenses/agpl-3.0.html
