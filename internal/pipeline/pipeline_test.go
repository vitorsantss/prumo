package pipeline

import (
	"errors"
	"strings"
	"testing"
)

const valido = `
nome: padrao
etapas:
  - id: spec
    nome: SPEC
    tipo: ia
    instrucoes: Escreva a especificação.
    entradas: [tarefa.md]
    saida: spec.md
    gate: manual
  - id: entrega
    nome: Entrega
    tipo: modelo
    entradas: [tarefa.md, spec.md]
    saida: entrega.md
    gate: automatico
`

func TestInterpretarPipelineValido(t *testing.T) {
	p, err := Interpretar([]byte(valido))
	if err != nil {
		t.Fatalf("esperava pipeline válido: %v", err)
	}
	if len(p.Etapas) != 2 || p.Etapas[1].Saida != "entrega.md" {
		t.Fatalf("etapas lidas errado: %+v", p.Etapas)
	}
}

func TestValidacaoApontaOCampo(t *testing.T) {
	casos := map[string]struct {
		troca, por, campo string
	}{
		"entrada inexistente": {"entradas: [tarefa.md, spec.md]", "entradas: [plano.md]", "etapas[1].entradas[0]"},
		"id repetido":         {"id: entrega", "id: spec", "etapas[1].id"},
		"gate desconhecido":   {"gate: manual", "gate: talvez", "etapas[0].gate"},
		"tipo desconhecido":   {"tipo: modelo", "tipo: magia", "etapas[1].tipo"},
		"saida com caminho":   {"saida: spec.md", "saida: ../spec.md", "etapas[0].saida"},
	}
	for nome, c := range casos {
		t.Run(nome, func(t *testing.T) {
			_, err := Interpretar([]byte(strings.Replace(valido, c.troca, c.por, 1)))
			var ev ErroValidacao
			if !errors.As(err, &ev) {
				t.Fatalf("esperava ErroValidacao, veio %v", err)
			}
			if ev.Campo != c.campo {
				t.Fatalf("campo %q, esperava %q", ev.Campo, c.campo)
			}
		})
	}
}

func TestCampoDesconhecidoMostraALinha(t *testing.T) {
	_, err := Interpretar([]byte(strings.Replace(valido, "gate: manual", "gate: manual\n    portao: x", 1)))
	if err == nil || !strings.Contains(err.Error(), "portao") {
		t.Fatalf("esperava erro citando o campo desconhecido, veio %v", err)
	}
}
