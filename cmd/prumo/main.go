package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/vitorsantss/prumo/internal/coordenador"
	"github.com/vitorsantss/prumo/internal/executor"
	"github.com/vitorsantss/prumo/internal/gate"
	"github.com/vitorsantss/prumo/internal/local"
	"github.com/vitorsantss/prumo/internal/pipeline"
	"github.com/vitorsantss/prumo/internal/tarefa"
)

const uso = `prumo: conduz uma tarefa de software em etapas com portões.

Uso:
  prumo validar <pipeline.yml>
  prumo executar --pipeline <pipeline.yml> --tarefa <tarefa.md> [opções]

Opções de executar:
  --executor duble|claude   quem executa as etapas de IA (padrão: duble, sem IA)
  --modelo <nome>           modelo repassado ao executor claude
  --workspace <nome>        nome do workspace (padrão: pasta atual)
  --aprovar-tudo            aprova todos os gates sem perguntar
`

func main() {
	if err := rodar(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}

func rodar(args []string, entrada io.Reader, saida io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(saida, uso)
		return nil
	}
	switch args[0] {
	case "validar":
		return validar(args[1:], saida)
	case "executar":
		return executar(args[1:], entrada, saida)
	case "ajuda", "-h", "--help":
		fmt.Fprint(saida, uso)
		return nil
	default:
		return fmt.Errorf("comando desconhecido %q\n\n%s", args[0], uso)
	}
}

func validar(args []string, saida io.Writer) error {
	if len(args) != 1 {
		return errors.New("informe o caminho do pipeline: prumo validar <pipeline.yml>")
	}
	p, err := pipeline.Carregar(args[0])
	if err != nil {
		return err
	}
	fmt.Fprintf(saida, "pipeline %q válido, %d etapa(s)\n", p.Nome, len(p.Etapas))
	return nil
}

func executar(args []string, entrada io.Reader, saida io.Writer) error {
	fs := flag.NewFlagSet("executar", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	caminhoPipeline := fs.String("pipeline", "", "")
	caminhoTarefa := fs.String("tarefa", "", "")
	nomeExecutor := fs.String("executor", "duble", "")
	modelo := fs.String("modelo", "", "")
	workspace := fs.String("workspace", "", "")
	aprovarTudo := fs.Bool("aprovar-tudo", false, "")
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%w\n\n%s", err, uso)
	}
	if *caminhoPipeline == "" || *caminhoTarefa == "" {
		return fmt.Errorf("--pipeline e --tarefa são obrigatórios\n\n%s", uso)
	}

	p, err := pipeline.Carregar(*caminhoPipeline)
	if err != nil {
		return err
	}
	ia, err := escolherExecutor(*nomeExecutor, *modelo)
	if err != nil {
		return err
	}
	dados, err := local.Dados()
	if err != nil {
		return err
	}
	nomeWorkspace, err := escolherWorkspace(*workspace)
	if err != nil {
		return err
	}
	t, err := tarefa.Criar(dados, nomeWorkspace, *caminhoTarefa, time.Now())
	if err != nil {
		return err
	}

	var decisor gate.Decisor = gate.NovoTerminal(entrada, saida)
	if *aprovarTudo {
		decisor = gate.Automatico{}
	}
	c := coordenador.Coordenador{IA: ia, Modelo: executor.Modelo{}, Decisor: decisor, Saida: saida}

	ctx, parar := signal.NotifyContext(context.Background(), os.Interrupt)
	defer parar()

	fmt.Fprintf(saida, "tarefa %s\npasta %s\n\n", t.ID, t.Pasta)
	resumo, err := c.Executar(ctx, p, t)
	fmt.Fprintf(saida, "\n%s · %d tokens de entrada · %d de saída · US$ %.4f\nrastro: %s\n",
		resumo.Status, resumo.TokensEntrada, resumo.TokensSaida, resumo.CustoUSD, t.Caminho("tarefa.json"))
	return err
}

func escolherExecutor(nome, modelo string) (executor.Executor, error) {
	switch nome {
	case "duble":
		return executor.Duble{}, nil
	case "claude":
		return executor.ClaudeCode{Modelo: modelo}, nil
	default:
		return nil, fmt.Errorf("executor %q desconhecido: use duble ou claude", nome)
	}
}

func escolherWorkspace(nome string) (string, error) {
	if nome != "" {
		return nome, nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(dir), nil
}
