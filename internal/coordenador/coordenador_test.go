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

	if etapa.Atributos[rastro.AtributoVoltas].(float64) != 2 {
		t.Fatalf("voltas das duas tentativas não somaram: %+v", etapa.Atributos)
	}
	if n := contarLinhas(t, tf.Caminho("usage.jsonl")); n != 2 {
		t.Fatalf("usage.jsonl com %d linhas, esperava 2 (só as tentativas de IA)", n)
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

type doisModelos struct{}

func (doisModelos) Executar(context.Context, executor.Pedido) (executor.Resultado, error) {
	return executor.Resultado{
		Conteudo: "# SPEC",
		Usos: []executor.Uso{
			{Provedor: "anthropic", Modelo: "modelo-a", TokensEntrada: 100, TokensLeituraCache: 80, TokensSaida: 10, CustoUSD: 0.1, BaseCusto: "list"},
			{Provedor: "anthropic", Modelo: "modelo-b", TokensEntrada: 20, TokensSaida: 5, CustoUSD: 0.02, BaseCusto: "list"},
		},
		Voltas:    3,
		DuracaoIA: 1500 * time.Millisecond,
	}, nil
}

func TestDoisModelosViramDuasLinhasESomamNaEtapa(t *testing.T) {
	p, tf := preparar(t)
	c := Coordenador{IA: doisModelos{}, Modelo: executor.Modelo{}, Decisor: gate.Automatico{}, Saida: io.Discard}
	if _, err := c.Executar(context.Background(), p, tf); err != nil {
		t.Fatal(err)
	}
	if n := contarLinhas(t, tf.Caminho("usage.jsonl")); n != 2 {
		t.Fatalf("usage.jsonl com %d linhas, esperava uma por modelo", n)
	}

	var etapa rastro.Etapa
	conteudo, _ := os.ReadFile(filepath.Join(tf.PastaEtapa(1, "spec"), "etapa.json"))
	if err := json.Unmarshal(conteudo, &etapa); err != nil {
		t.Fatal(err)
	}
	a := etapa.Atributos
	if a[rastro.AtributoTokensEntrada].(float64) != 120 || a[rastro.AtributoTokensLeituraCache].(float64) != 80 {
		t.Fatalf("tokens da etapa não somaram os modelos: %+v", a)
	}
	if a[rastro.AtributoVoltas].(float64) != 3 || a[rastro.AtributoDuracaoIAMS].(float64) != 1500 || a[rastro.AtributoBaseCusto] != "list" {
		t.Fatalf("voltas, duração ou base do custo erradas: %+v", a)
	}
	if _, ok := a[rastro.AtributoModelo]; ok {
		t.Fatalf("com dois modelos não pode haver um modelo principal: %+v", a)
	}
	if modelos, _ := a[rastro.AtributoModelos].([]any); len(modelos) != 2 || modelos[0] != "modelo-a" {
		t.Fatalf("lista de modelos errada: %+v", a[rastro.AtributoModelos])
	}

	var entrega rastro.Etapa
	conteudo, _ = os.ReadFile(filepath.Join(tf.PastaEtapa(2, "entrega"), "etapa.json"))
	if err := json.Unmarshal(conteudo, &entrega); err != nil {
		t.Fatal(err)
	}
	if len(entrega.Atributos) != 0 {
		t.Fatalf("etapa sem IA registrou uso: %+v", entrega.Atributos)
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
