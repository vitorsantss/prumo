package local

import (
	"path/filepath"
	"testing"
)

func TestPrumoHomeTemPrioridade(t *testing.T) {
	t.Setenv("PRUMO_HOME", "/tmp/prumo")
	t.Setenv("XDG_DATA_HOME", "/xdg/data")
	dados, _ := Dados()
	estado, _ := Estado()
	if dados != filepath.FromSlash("/tmp/prumo/dados") || estado != filepath.FromSlash("/tmp/prumo/estado") {
		t.Fatalf("dados=%s estado=%s", dados, estado)
	}
}

func TestSegueXDG(t *testing.T) {
	t.Setenv("PRUMO_HOME", "")
	t.Setenv("XDG_DATA_HOME", "/xdg/data")
	t.Setenv("XDG_STATE_HOME", "/xdg/state")
	dados, _ := Dados()
	estado, _ := Estado()
	if dados != filepath.FromSlash("/xdg/data/prumo") || estado != filepath.FromSlash("/xdg/state/prumo") {
		t.Fatalf("dados=%s estado=%s", dados, estado)
	}
}

func TestPadraoNaPastaDoUsuario(t *testing.T) {
	t.Setenv("PRUMO_HOME", "")
	t.Setenv("XDG_DATA_HOME", "relativo-nao-vale")
	t.Setenv("HOME", "/casa")
	dados, _ := Dados()
	if dados != filepath.FromSlash("/casa/.local/share/prumo") {
		t.Fatalf("dados=%s", dados)
	}
}
