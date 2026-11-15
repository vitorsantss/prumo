package executor

import (
	"context"
	"strings"
	"testing"

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
	if r.Uso.Provedor != "duble" || r.Uso.TokensEntrada == 0 || !strings.Contains(r.Conteudo, "Inclua o limite.") {
		t.Fatalf("resultado inesperado: %+v", r)
	}
}

func TestInterpretarRespostaClaude(t *testing.T) {
	bruto := `{"result":"# SPEC","is_error":false,"total_cost_usd":0.012,
		"usage":{"input_tokens":10,"output_tokens":50,"cache_read_input_tokens":90,"cache_creation_input_tokens":0},
		"modelUsage":{"claude-opus-5-5":{}}}`
	r, err := interpretarRespostaClaude([]byte(bruto))
	if err != nil {
		t.Fatal(err)
	}
	if r.Conteudo != "# SPEC" || r.Uso.TokensEntrada != 100 || r.Uso.TokensSaida != 50 || r.Uso.Modelo != "claude-opus-5-5" {
		t.Fatalf("uso lido errado: %+v", r)
	}
}

func TestRespostaClaudeComErro(t *testing.T) {
	if _, err := interpretarRespostaClaude([]byte(`{"result":"sem crédito","is_error":true}`)); err == nil {
		t.Fatal("esperava erro")
	}
}
