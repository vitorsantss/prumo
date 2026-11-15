package pipeline

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
)

type Gate string

const (
	GateManual     Gate = "manual"
	GateAutomatico Gate = "automatico"
)

type Tipo string

const (
	TipoIA     Tipo = "ia"
	TipoModelo Tipo = "modelo"
)

const ArquivoTarefa = "tarefa.md"

type Etapa struct {
	ID         string   `yaml:"id"`
	Nome       string   `yaml:"nome"`
	Tipo       Tipo     `yaml:"tipo"`
	Instrucoes string   `yaml:"instrucoes"`
	Entradas   []string `yaml:"entradas"`
	Saida      string   `yaml:"saida"`
	Gate       Gate     `yaml:"gate"`
}

type Pipeline struct {
	Nome   string  `yaml:"nome"`
	Etapas []Etapa `yaml:"etapas"`
}

type ErroValidacao struct {
	Campo  string
	Motivo string
}

func (e ErroValidacao) Error() string {
	return fmt.Sprintf("%s: %s", e.Campo, e.Motivo)
}

var formatoID = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

func Carregar(caminho string) (Pipeline, error) {
	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		return Pipeline{}, fmt.Errorf("lendo pipeline: %w", err)
	}
	return Interpretar(conteudo)
}

func Interpretar(conteudo []byte) (Pipeline, error) {
	var p Pipeline
	if err := yaml.UnmarshalWithOptions(conteudo, &p, yaml.Strict()); err != nil {
		return Pipeline{}, fmt.Errorf("pipeline inválido:\n%s", yaml.FormatError(err, false, true))
	}
	if err := p.Validar(); err != nil {
		return Pipeline{}, err
	}
	return p, nil
}

func (p Pipeline) Validar() error {
	if strings.TrimSpace(p.Nome) == "" {
		return ErroValidacao{"nome", "obrigatório"}
	}
	if len(p.Etapas) == 0 {
		return ErroValidacao{"etapas", "precisa de pelo menos uma etapa"}
	}
	disponiveis := map[string]bool{ArquivoTarefa: true}
	ids := map[string]bool{}
	for i, e := range p.Etapas {
		if err := e.validar(fmt.Sprintf("etapas[%d]", i), ids, disponiveis); err != nil {
			return err
		}
		ids[e.ID] = true
		disponiveis[e.Saida] = true
	}
	return nil
}

func (e Etapa) validar(campo string, ids, disponiveis map[string]bool) error {
	switch {
	case !formatoID.MatchString(e.ID):
		return ErroValidacao{campo + ".id", "use letras minúsculas, números e hífen, começando por letra"}
	case ids[e.ID]:
		return ErroValidacao{campo + ".id", fmt.Sprintf("%q repetido", e.ID)}
	case strings.TrimSpace(e.Nome) == "":
		return ErroValidacao{campo + ".nome", "obrigatório"}
	case e.Tipo != TipoIA && e.Tipo != TipoModelo:
		return ErroValidacao{campo + ".tipo", fmt.Sprintf("use %q ou %q", TipoIA, TipoModelo)}
	case e.Tipo == TipoIA && strings.TrimSpace(e.Instrucoes) == "":
		return ErroValidacao{campo + ".instrucoes", "obrigatório pra etapa de IA"}
	case e.Gate != GateManual && e.Gate != GateAutomatico:
		return ErroValidacao{campo + ".gate", fmt.Sprintf("use %q ou %q", GateManual, GateAutomatico)}
	case !nomeDeArquivoValido(e.Saida):
		return ErroValidacao{campo + ".saida", "informe um nome de arquivo simples, como spec.md"}
	case disponiveis[e.Saida]:
		return ErroValidacao{campo + ".saida", fmt.Sprintf("%q já é produzido antes", e.Saida)}
	}
	for j, entrada := range e.Entradas {
		if !disponiveis[entrada] {
			return ErroValidacao{
				fmt.Sprintf("%s.entradas[%d]", campo, j),
				fmt.Sprintf("%q não existe: use %s ou a saída de uma etapa anterior", entrada, ArquivoTarefa),
			}
		}
	}
	return nil
}

func nomeDeArquivoValido(nome string) bool {
	return nome != "" && !strings.ContainsAny(nome, `/\`) && nome != "." && nome != ".."
}
