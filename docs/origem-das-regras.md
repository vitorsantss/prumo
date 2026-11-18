# Origem das regras

Cada regra de pesquisa e de trabalho do [AGENTS.md](../AGENTS.md) nasceu de um erro real no próprio Prumo. Aqui fica a história de cada uma: o que aconteceu, o que mudou e onde conferir. Número de issue ou PR aparece como atalho pro GitHub; a prova é o que está neste repositório.

## Fonte antes da tese
- **O que aconteceu:** na discussão do ADR-0014, a IA propôs "registrar o uso por sessão" como decisão, sem referência. A fonte só apareceu depois que o mantenedor cobrou.
- **O que mudou:** a decisão ganhou fontes de modelagem dimensional e de event sourcing, e a regra entrou no AGENTS.md.
- **Onde conferir:** notas de rodapé do [ADR-0014](adr/0014-registro-de-uso-de-ia.md).

## Link só depois de aberto
- **O que aconteceu:** dois links foram escritos de memória na conversa: o da sentença do LG München I sobre Google Fonts, que apontava pra um site intermediário, e um da especificação W3C PROV. Os dois foram pegos antes de entrar no repositório.
- **O que mudou:** a sentença passou a ser citada pelo portal oficial de legislação da Baviera, aberto antes de citar. A W3C PROV não foi citada.
- **Onde conferir:** nota `lg-munique` em [tipografia](design/fundamentos/tipografia.md), que entrou na `main` no commit `9c60980`.

## Origem aponta pra evidência que mora no repositório
- **O que aconteceu:** na revisão do PR #110, as regras diziam a origem com data solta ("em 16/11") e citavam "um link de decisão judicial" sem dizer qual, o que fazia parecer que o projeto estava num processo. A primeira correção trocou as datas por números de PR. A revisão seguinte notou que isso prende o Prumo ao GitHub: dentro de um arquivo, o próprio GitHub nem transforma número de issue em link[^autolink], e só o script do site faz essa conversão.
- **O que mudou:** a regra e este arquivo.
- **Onde conferir:** este arquivo.

## Ideia ou relato vira issue, não código
- **O que aconteceu:** três relatos do site (links do índice desalinhados #99, posição do ícone do GitHub #100, menu do celular que não abria #101) viraram código sem ninguém pedir, quando o combinado era só anotar e seguir a fila.
- **O que mudou:** o PR desses relatos (#103) espera a vez dele na fila, e a regra entrou no AGENTS.md.
- **Onde conferir:** a fila em [processo](processo.md).

## Procure antes de abrir
- **O que aconteceu:** uma issue sobre métrica de tokens (#107) repetiu outra já aberta sobre métrica de esforço por modelo (#40).
- **O que mudou:** a repetida foi fechada, e a regra entrou no AGENTS.md.
- **Onde conferir:** [ADR-0014](adr/0014-registro-de-uso-de-ia.md), que cita só a #40.

## Celular conferido com toque
- **O que aconteceu:** a troca de página com transição deu um nome de transição à barra lateral. Isso criou uma camada nova na tela e deixou o menu do celular sem resposta ao toque (#101). A mudança só tinha sido conferida com mouse.
- **O que mudou:** o nome de transição passou a valer só em telas largas (PR #103), e a conferência com toque entrou na regra 7 da [camada Site](design/site.md).
- **Onde conferir:** commit `9c60980` na `main` (a mudança da transição) e [movimento](design/fundamentos/movimento.md).

## Estilo global com seletor de escopo
- **O que aconteceu:** o cabeçalho da página inicial usava a mesma classe (`.topo`) que a moldura da documentação e herdou a transição dela. Ao navegar da página inicial pra documentação, o cabeçalho embaralhava.
- **O que mudou:** os nomes de transição passaram a valer só dentro da moldura da documentação, e a regra entrou como regra 11 da [camada Site](design/site.md).
- **Onde conferir:** commit `9c60980` na `main` e [movimento](design/fundamentos/movimento.md).

[^autolink]: GitHub Docs, *Autolinked references and URLs* · https://docs.github.com/en/enterprise-server@3.16/get-started/writing-on-github/working-with-advanced-formatting/autolinked-references-and-urls
