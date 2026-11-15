package executor

import (
	"context"
	"fmt"
	"strings"
)

type Modelo struct{}

func (Modelo) Executar(_ context.Context, p Pedido) (Resultado, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", p.Etapa.Nome)
	for _, e := range p.Entradas {
		fmt.Fprintf(&b, "## %s\n\n%s\n\n", e.Nome, strings.TrimSpace(e.Conteudo))
	}
	return Resultado{Conteudo: b.String(), Uso: Uso{Provedor: "prumo", Modelo: "modelo"}}, nil
}
