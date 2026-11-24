package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecutarExemploComDuble(t *testing.T) {
	t.Setenv("PRUMO_HOME", t.TempDir())
	var saida bytes.Buffer
	err := rodar([]string{
		"executar",
		"--pipeline", filepath.Join("..", "..", "exemplos", "pipeline.yml"),
		"--tarefa", filepath.Join("..", "..", "exemplos", "rinha-2024q1", "1-transacao-e-extrato.md"),
		"--workspace", "rinha",
	}, strings.NewReader("a\n"), &saida)
	if err != nil {
		t.Fatalf("%v\n%s", err, saida.String())
	}
	if !strings.Contains(saida.String(), "concluida") {
		t.Fatalf("saída inesperada:\n%s", saida.String())
	}
}

func TestValidarExemplo(t *testing.T) {
	var saida bytes.Buffer
	if err := rodar([]string{"validar", filepath.Join("..", "..", "exemplos", "pipeline.yml")}, nil, &saida); err != nil {
		t.Fatal(err)
	}
}

func TestComandoDesconhecido(t *testing.T) {
	if err := rodar([]string{"voar"}, nil, os.Stdout); err == nil {
		t.Fatal("esperava erro")
	}
}
