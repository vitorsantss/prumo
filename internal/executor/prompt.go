package executor

import "strings"

func montarPrompt(p Pedido) string {
	var b strings.Builder
	b.WriteString("Você está executando a etapa \"" + p.Etapa.Nome + "\" de uma tarefa de software no Prumo.\n\n")
	b.WriteString("## Instruções\n" + strings.TrimSpace(p.Etapa.Instrucoes) + "\n\n")
	for _, e := range p.Entradas {
		b.WriteString("## Arquivo de entrada: " + e.Nome + "\n" + strings.TrimSpace(e.Conteudo) + "\n\n")
	}
	if p.Ajuste != "" {
		b.WriteString("## Ajuste pedido na revisão anterior\n" + strings.TrimSpace(p.Ajuste) + "\n\n")
	}
	b.WriteString("Responda só com o conteúdo final do arquivo " + p.Etapa.Saida + ", em Markdown, sem comentários antes ou depois.\n")
	return b.String()
}
