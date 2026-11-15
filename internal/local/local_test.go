package local

import (
	"path/filepath"
	"testing"
)

func TestPrumoHomeTemPrioridade(t *testing.T) {
	raiz := t.TempDir()
	t.Setenv("PRUMO_HOME", raiz)
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	dados, _ := Dados()
	estado, _ := Estado()
	if dados != filepath.Join(raiz, "dados") || estado != filepath.Join(raiz, "estado") {
		t.Fatalf("dados=%s estado=%s", dados, estado)
	}
}

func TestSegueXDG(t *testing.T) {
	xdgDados, xdgEstado := t.TempDir(), t.TempDir()
	t.Setenv("PRUMO_HOME", "")
	t.Setenv("XDG_DATA_HOME", xdgDados)
	t.Setenv("XDG_STATE_HOME", xdgEstado)
	dados, _ := Dados()
	estado, _ := Estado()
	if dados != filepath.Join(xdgDados, "prumo") || estado != filepath.Join(xdgEstado, "prumo") {
		t.Fatalf("dados=%s estado=%s", dados, estado)
	}
}

func TestIgnoraXDGRelativo(t *testing.T) {
	casa := t.TempDir()
	t.Setenv("PRUMO_HOME", "")
	t.Setenv("XDG_DATA_HOME", "relativo-nao-vale")
	t.Setenv("HOME", casa)
	t.Setenv("USERPROFILE", casa)
	dados, _ := Dados()
	if dados != filepath.Join(casa, ".local", "share", "prumo") {
		t.Fatalf("dados=%s", dados)
	}
}
