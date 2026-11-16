# Acessibilidade

O site é medido com o Lighthouse e não aceita nota abaixo de 90 (`npm run auditar`, em [Como mudar o site](../../como-mudar-o-site.md)). As regras abaixo são o que faz a nota ficar alta de verdade, não só no teste.

| Regra | Por quê | Base |
| --- | --- | --- |
| Contraste de texto de pelo menos 4,5:1 | Ler com pouca visão ou sol na tela | WCAG 1.4.3[^contraste] |
| Alvo de toque de 44 × 44 px | Acertar com o dedo | WCAG 2.5.5[^alvo] |
| Foco sempre visível | Navegar só pelo teclado | WCAG 2.4.7[^foco] |
| Todo botão e link com nome acessível; botão só de ícone leva `aria-label` | Leitor de tela anuncia o que faz | WCAG 4.1.2[^nome] |
| Idioma da página e de trechos em outro idioma declarado (`lang`) | Leitor de tela pronuncia certo | WCAG 3.1.1 e 3.1.2[^idioma] |
| Estado com forma, não só cor | Quem não distingue cores entende o estado | WCAG 1.4.1[^cor] |
| Movimento desligado com `prefers-reduced-motion` | Evitar enjoo e distração | WCAG 2.3.3[^animacao] |
| Elementos de verdade (`<button>`, `<a href>`, `<dialog>`) | Funciona com teclado e tecnologia assistiva sem truque | — |

Testes automatizados acham elementos pelo papel e pelo nome acessível, não por `data-testid`: se o teste acha, o leitor de tela também acha (#70).

[^contraste]: W3C, *WCAG 2.2, 1.4.3 Contrast (Minimum)* · https://www.w3.org/TR/WCAG22/#contrast-minimum
[^alvo]: W3C, *WCAG 2.2, 2.5.5 Target Size (Enhanced)* · https://www.w3.org/TR/WCAG22/#target-size-enhanced
[^foco]: W3C, *WCAG 2.2, 2.4.7 Focus Visible* · https://www.w3.org/TR/WCAG22/#focus-visible
[^nome]: W3C, *WCAG 2.2, 4.1.2 Name, Role, Value* · https://www.w3.org/TR/WCAG22/#name-role-value
[^idioma]: W3C, *WCAG 2.2, 3.1.1 Language of Page e 3.1.2 Language of Parts* · https://www.w3.org/TR/WCAG22/#language-of-page
[^cor]: W3C, *WCAG 2.2, 1.4.1 Use of Color* · https://www.w3.org/TR/WCAG22/#use-of-color
[^animacao]: W3C, *WCAG 2.2, 2.3.3 Animation from Interactions* · https://www.w3.org/TR/WCAG22/#animation-from-interactions
