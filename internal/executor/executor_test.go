package executor

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

func pedido() Pedido {
	return Pedido{
		Etapa:    pipeline.Etapa{ID: "spec", Nome: "SPEC", Tipo: pipeline.TipoIA, Instrucoes: "Especifique.", Saida: "spec.md"},
		Entradas: []Entrada{{Nome: "tarefa.md", Conteudo: "# Extrato"}},
		Ajuste:   "Inclua o limite.",
	}
}

func TestPromptLevaInstrucoesEntradasEAjuste(t *testing.T) {
	prompt := montarPrompt(pedido())
	for _, esperado := range []string{"Especifique.", "tarefa.md", "# Extrato", "Inclua o limite.", "spec.md"} {
		if !strings.Contains(prompt, esperado) {
			t.Fatalf("prompt sem %q:\n%s", esperado, prompt)
		}
	}
}

func TestDubleNaoChamaIAERegistraUso(t *testing.T) {
	r, err := Duble{}.Executar(context.Background(), pedido())
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Usos) != 1 || r.Usos[0].Provedor != "duble" || r.Usos[0].TokensEntrada == 0 || r.Voltas != 1 {
		t.Fatalf("uso do dublê inesperado: %+v", r)
	}
	if !strings.Contains(r.Conteudo, "Inclua o limite.") {
		t.Fatalf("dublê ignorou o ajuste:\n%s", r.Conteudo)
	}
}

func TestModeloNaoRegistraUsoDeIA(t *testing.T) {
	r, err := Modelo{}.Executar(context.Background(), pedido())
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Usos) != 0 || r.Voltas != 0 {
		t.Fatalf("etapa sem IA devolveu uso: %+v", r)
	}
}

func TestRespostaRealDoClaudeViraUmUsoPorModelo(t *testing.T) {
	bruto, err := os.ReadFile("testdata/claude-oi.json")
	if err != nil {
		t.Fatal(err)
	}
	r, err := interpretarRespostaClaude(bruto)
	if err != nil {
		t.Fatal(err)
	}
	esperado := Uso{
		Provedor:           "anthropic",
		Modelo:             "claude-opus-5-5",
		TokensEntrada:      28520,
		TokensLeituraCache: 14801,
		TokensCriacaoCache: 13717,
		TokensSaida:        4,
		CustoUSD:           0.1127842,
		BaseCusto:          "list",
	}
	if r.Conteudo != "oi" || len(r.Usos) != 1 || r.Usos[0] != esperado {
		t.Fatalf("uso lido errado: %+v", r)
	}
	if r.Voltas != 1 || r.DuracaoIA != 4277*time.Millisecond {
		t.Fatalf("voltas ou duração erradas: %d, %s", r.Voltas, r.DuracaoIA)
	}
}

func TestDoisModelosSaemSempreNaMesmaOrdem(t *testing.T) {
	bruto := []byte(`{"result":"ok","num_turns":3,"modelUsage":{
		"modelo-b":{"inputTokens":1,"outputTokens":2,"costUSD":0.2,"costBasis":"list"},
		"modelo-a":{"inputTokens":3,"outputTokens":4,"cacheReadInputTokens":10,"costUSD":0.1,"costBasis":"list"}}}`)
	for i := 0; i < 20; i++ {
		r, err := interpretarRespostaClaude(bruto)
		if err != nil {
			t.Fatal(err)
		}
		if len(r.Usos) != 2 || r.Usos[0].Modelo != "modelo-a" || r.Usos[1].Modelo != "modelo-b" {
			t.Fatalf("ordem dos modelos instável: %+v", r.Usos)
		}
		if r.Usos[0].TokensEntrada != 13 {
			t.Fatalf("entrada do modelo-a sem o cache: %+v", r.Usos[0])
		}
	}
}

func TestRespostaSemModelUsageUsaOTotalGeral(t *testing.T) {
	bruto := []byte(`{"result":"ok","total_cost_usd":0.5,
		"usage":{"input_tokens":10,"output_tokens":50,"cache_read_input_tokens":90,"cache_creation_input_tokens":5}}`)
	r, err := interpretarRespostaClaude(bruto)
	if err != nil {
		t.Fatal(err)
	}
	u := r.Usos[0]
	if len(r.Usos) != 1 || u.Modelo != "" || u.TokensEntrada != 105 || u.TokensLeituraCache != 90 || u.CustoUSD != 0.5 {
		t.Fatalf("plano B lido errado: %+v", r.Usos)
	}
}

func TestRespostaClaudeComErro(t *testing.T) {
	if _, err := interpretarRespostaClaude([]byte(`{"result":"sem crédito","is_error":true}`)); err == nil {
		t.Fatal("esperava erro")
	}
}
