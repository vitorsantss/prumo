// Package executor define a porta que roda uma etapa de IA e os adaptadores
// que a implementam (ADR-0005). O núcleo só conhece a interface Executor.
package executor

import (
	"context"
	"time"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

// Entrada é um arquivo declarado no pipeline.yml, já lido pelo núcleo.
type Entrada struct {
	Nome     string
	Conteudo string
}

// Pedido é tudo o que o executor recebe para rodar uma tentativa de uma etapa.
// Ajuste traz o comentário do gate quando a tentativa anterior foi devolvida.
// Pasta é a pasta da tarefa, onde o executor roda.
type Pedido struct {
	Etapa    pipeline.Etapa
	Entradas []Entrada
	Ajuste   string
	Pasta    string
}

// Uso é o consumo de um modelo numa tentativa. É o grão do registro de uso:
// uma linha do usage.jsonl por Uso (ADR-0014, decisão 1).
type Uso struct {
	Provedor string
	Modelo   string

	// TokensEntrada é o total de entrada, já somando a leitura e a criação de
	// cache, como pede a convenção do OpenTelemetry (ADR-0014, decisão 3).
	TokensEntrada      int
	TokensLeituraCache int
	TokensCriacaoCache int
	TokensSaida        int

	// CustoUSD vem do provedor. BaseCusto diz como ele foi calculado: "list"
	// é preço de tabela da API, não o que foi pago numa assinatura.
	CustoUSD  float64
	BaseCusto string
}

// Resultado é o que uma tentativa devolve. Usos tem um item por modelo usado,
// em ordem estável; uma etapa sem IA devolve Usos vazio. Voltas e DuracaoIA
// são da tentativa inteira, não de um modelo.
type Resultado struct {
	Conteudo  string
	Usos      []Uso
	Voltas    int
	DuracaoIA time.Duration
}

// Executor roda uma tentativa de uma etapa. Cada chamada é uma sessão nova,
// sem histórico das anteriores (ADR-0007, decisão 5).
type Executor interface {
	Executar(ctx context.Context, p Pedido) (Resultado, error)
}
