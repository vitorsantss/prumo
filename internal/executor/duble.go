package executor

import (
	"context"
	"fmt"
	"strings"
)

type Duble struct{}

func (Duble) Executar(_ context.Context, p Pedido) (Resultado, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s (resposta simulada)\n\n", p.Etapa.Nome)
	b.WriteString("Gerado pelo executor dublê, sem IA, pra testar o fluxo de ponta a ponta.\n\n")
	b.WriteString("## Entradas lidas\n")
	for _, e := range p.Entradas {
		fmt.Fprintf(&b, "- %s (%d caracteres)\n", e.Nome, len(e.Conteudo))
	}
	if p.Ajuste != "" {
		fmt.Fprintf(&b, "\n## Ajuste aplicado\n%s\n", p.Ajuste)
	}
	prompt := montarPrompt(p)
	return Resultado{
		Conteudo: b.String(),
		Usos: []Uso{{
			Provedor:      "duble",
			Modelo:        "duble",
			TokensEntrada: len(prompt) / 4,
			TokensSaida:   b.Len() / 4,
		}},
		Voltas: 1,
	}, nil
}
