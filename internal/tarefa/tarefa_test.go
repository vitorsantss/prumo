package tarefa

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCriarCopiaATarefaEMontaOID(t *testing.T) {
	dir := t.TempDir()
	origem := filepath.Join(dir, "origem.md")
	os.WriteFile(origem, []byte("# Extrato do cliente\n\nDetalhes."), 0o644)

	tf, err := Criar(filepath.Join(dir, "dados"), "Rinha Backend", origem, time.Date(2026, 11, 15, 9, 30, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if tf.ID != "20261115-093000-extrato-do-cliente" {
		t.Fatalf("id %q", tf.ID)
	}
	if filepath.Base(filepath.Dir(filepath.Dir(tf.Pasta))) != "rinha-backend" {
		t.Fatalf("workspace na pasta errada: %s", tf.Pasta)
	}
	if _, err := os.Stat(tf.Caminho("tarefa.md")); err != nil {
		t.Fatalf("tarefa.md não copiado: %v", err)
	}
}
