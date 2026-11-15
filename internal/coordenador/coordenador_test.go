package coordenador

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/vitorsantss/prumo/internal/executor"
	"github.com/vitorsantss/prumo/internal/gate"
	"github.com/vitorsantss/prumo/internal/pipeline"
	"github.com/vitorsantss/prumo/internal/rastro"
	"github.com/vitorsantss/prumo/internal/tarefa"
)

const pipelineTeste = `
nome: teste
etapas:
  - id: spec
    nome: SPEC
    tipo: ia
    instrucoes: Especifique.
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

func preparar(t *testing.T) (pipeline.Pipeline, tarefa.Tarefa) {
	t.Helper()
	dir := t.TempDir()
	origem := filepath.Join(dir, "tarefa.md")
	os.WriteFile(origem, []byte("# Extrato\n\nMostrar saldo."), 0o644)
	p, err := pipeline.Interpretar([]byte(pipelineTeste))
	if err != nil {
		t.Fatal(err)
	}
	tf, err := tarefa.Criar(filepath.Join(dir, "dados"), "ws", origem, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	return p, tf
}

func TestFluxoComAjusteNoGate(t *testing.T) {
	p, tf := preparar(t)
	c := Coordenador{
		IA:      executor.Duble{},
		Modelo:  executor.Modelo{},
		Decisor: gate.NovoTerminal(strings.NewReader("j\nincluir limite\na\n"), io.Discard),
		Saida:   io.Discard,
	}
	resumo, err := c.Executar(context.Background(), p, tf)
	if err != nil {
		t.Fatal(err)
	}
	if resumo.Status != "concluida" || len(resumo.Etapas) != 2 {
		t.Fatalf("resumo %+v", resumo)
	}

	spec, _ := os.ReadFile(tf.Caminho("spec.md"))
	if !strings.Contains(string(spec), "incluir limite") {
		t.Fatal("o ajuste não chegou na segunda tentativa")
	}
	entrega, _ := os.ReadFile(tf.Caminho("entrega.md"))
	if !strings.Contains(string(entrega), "## spec.md") {
		t.Fatal("entrega sem a spec")
	}

	var etapa rastro.Etapa
	conteudo, _ := os.ReadFile(filepath.Join(tf.PastaEtapa(1, "spec"), "etapa.json"))
	if err := json.Unmarshal(conteudo, &etapa); err != nil {
		t.Fatal(err)
	}
	if etapa.Tentativas != 2 || len(etapa.Gates) != 2 || etapa.Lidos[0].Caminho != "tarefa.md" || etapa.Escritos[0].SHA256 == "" {
		t.Fatalf("rastro da etapa incompleto: %+v", etapa)
	}
	if etapa.Atributos[rastro.AtributoTokensEntrada].(float64) <= 0 {
		t.Fatal("tokens não registrados")
	}

	if n := contarLinhas(t, tf.Caminho("usage.jsonl")); n != 3 {
		t.Fatalf("usage.jsonl com %d linhas, esperava 3", n)
	}
	if n := contarLinhas(t, tf.Caminho("gates.jsonl")); n != 3 {
		t.Fatalf("gates.jsonl com %d linhas, esperava 3", n)
	}
}

func TestParaQuandoAsTentativasAcabam(t *testing.T) {
	p, tf := preparar(t)
	c := Coordenador{
		IA:            executor.Duble{},
		Modelo:        executor.Modelo{},
		Decisor:       gate.NovoTerminal(strings.NewReader("r\nr\n"), io.Discard),
		Saida:         io.Discard,
		MaxTentativas: 2,
	}
	resumo, err := c.Executar(context.Background(), p, tf)
	if err == nil || resumo.Status != statusEsgotada {
		t.Fatalf("esperava parar com tentativas esgotadas: status=%s err=%v", resumo.Status, err)
	}
	if _, err := os.Stat(tf.Caminho("tarefa.json")); err != nil {
		t.Fatal("tarefa.json precisa existir mesmo com falha")
	}
}

func contarLinhas(t *testing.T, caminho string) int {
	t.Helper()
	f, err := os.Open(caminho)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	n := 0
	for s := bufio.NewScanner(f); s.Scan(); {
		n++
	}
	return n
}
