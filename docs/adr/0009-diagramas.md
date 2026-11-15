# 0009 · Diagramas: um critério, não uma ferramenta

- **Status:** Aceito
- **Data:** 2026-11-15

## Contexto
Desenhar ajuda a decidir e a explicar: o próprio modelo de execução (ADR-0007) nasceu de um desenho. O mantenedor usa Excalidraw, com e sem IA, mas quem contribui pode usar outra ferramenta, ou preferir não desenhar. Impor uma ferramenta exclui pessoas; não ter regra nenhuma gera imagens soltas que ninguém consegue editar depois.

## Decisão
1. **Diagrama é opcional.** Entra quando o texto sozinho não explica bem: em ADR, issue, PR ou qualquer documento.
2. **Qualquer ferramenta vale, desde que cumpra duas condições:**
   - aparece direto no GitHub, sem instalar nada;
   - a fonte editável fica no repositório, num formato aberto.
3. **Recomendação de uso:**

| Ferramenta | Quando | Custo |
| --- | --- | --- |
| Mermaid (bloco ` ```mermaid ` no Markdown) | Fluxos, sequências, estados, tudo que cabe em caixas e setas | Quase zero: é texto, o diff aparece no PR, o GitHub desenha sozinho, a IA escreve bem. Layout rígido |
| Excalidraw, salvo como `.excalidraw.svg` com a cena embutida | Desenho livre: arquitetura, rascunho, ideia | O GitHub mostra como imagem e o arquivo reabre editável no Excalidraw. O diff é ilegível e editar exige abrir a ferramenta |
| draw.io (`.drawio.svg`) e afins | Quem já usa | Mesmas condições do Excalidraw |

4. Arquivos de desenho ficam ao lado do documento que os usa, ou em `docs/diagramas/` quando servem a vários.
5. Gerar diagramas com IA (por exemplo, uma skill de agente que escreve Mermaid ou Excalidraw neste padrão) é uma facilidade opcional. Quem desenha à mão segue só o item 2.

## Consequências
- Todo diagrama do repositório pode ser visto no navegador e editado por outra pessoa depois.
- Fluxos simples tendem a ir pro Mermaid, que é revisável linha a linha no PR.
- Troca de ferramenta não exige ADR novo, desde que a nova cumpra o item 2.

## Fontes
- GitHub Docs, *Creating diagrams* (Mermaid no Markdown) · https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-diagrams
- Excalidraw, exportação com a cena embutida (MIT) · https://github.com/excalidraw/excalidraw
- Simon Brown, *The C4 model*, sobre diagramas que explicam em vez de enfeitar · https://c4model.com
