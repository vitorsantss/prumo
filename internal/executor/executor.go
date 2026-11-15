package executor

import (
	"context"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

type Entrada struct {
	Nome     string
	Conteudo string
}

type Pedido struct {
	Etapa    pipeline.Etapa
	Entradas []Entrada
	Ajuste   string
	Pasta    string
}

type Uso struct {
	Provedor      string
	Modelo        string
	TokensEntrada int
	TokensSaida   int
	CustoUSD      float64
}

type Resultado struct {
	Conteudo string
	Uso      Uso
}

type Executor interface {
	Executar(ctx context.Context, p Pedido) (Resultado, error)
}
