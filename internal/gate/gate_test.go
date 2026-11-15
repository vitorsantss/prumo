package gate

import (
	"io"
	"strings"
	"testing"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

func TestTerminalPedeComentarioNoAjuste(t *testing.T) {
	term := NovoTerminal(strings.NewReader("x\nj\n\ncobrir limite\n"), io.Discard)
	d, err := term.Decidir(pipeline.Etapa{Nome: "SPEC"}, "spec.md", "# SPEC")
	if err != nil {
		t.Fatal(err)
	}
	if d.Tipo != Ajustar || d.Comentario != "cobrir limite" {
		t.Fatalf("decisão %+v", d)
	}
}

func TestTerminalSemEntradaDaErro(t *testing.T) {
	if _, err := NovoTerminal(strings.NewReader(""), io.Discard).Decidir(pipeline.Etapa{}, "", ""); err == nil {
		t.Fatal("esperava erro")
	}
}
