package gate

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

type Tipo string

const (
	Aprovar Tipo = "aprovar"
	Ajustar Tipo = "ajustar"
	Refazer Tipo = "refazer"
)

type Decisao struct {
	Tipo       Tipo
	Comentario string
}

type Decisor interface {
	Decidir(etapa pipeline.Etapa, caminhoSaida, conteudo string) (Decisao, error)
}

type Automatico struct{}

func (Automatico) Decidir(pipeline.Etapa, string, string) (Decisao, error) {
	return Decisao{Tipo: Aprovar}, nil
}

type Terminal struct {
	Entrada      *bufio.Reader
	Saida        io.Writer
	LinhasPrevia int
}

func NovoTerminal(entrada io.Reader, saida io.Writer) Terminal {
	return Terminal{Entrada: bufio.NewReader(entrada), Saida: saida, LinhasPrevia: 30}
}

func (t Terminal) Decidir(etapa pipeline.Etapa, caminhoSaida, conteudo string) (Decisao, error) {
	fmt.Fprintf(t.Saida, "\n── %s pronta: %s\n\n%s\n", etapa.Nome, caminhoSaida, previa(conteudo, t.LinhasPrevia))
	for {
		fmt.Fprint(t.Saida, "\n[a] aprovar · [j] pedir ajuste · [r] refazer daqui > ")
		resposta, err := lerLinha(t.Entrada)
		if err != nil {
			return Decisao{}, err
		}
		switch strings.ToLower(resposta) {
		case "a", "aprovar":
			return Decisao{Tipo: Aprovar}, nil
		case "r", "refazer":
			return Decisao{Tipo: Refazer}, nil
		case "j", "ajuste", "ajustar":
			comentario, err := t.pedirComentario()
			if err != nil {
				return Decisao{}, err
			}
			return Decisao{Tipo: Ajustar, Comentario: comentario}, nil
		default:
			fmt.Fprintln(t.Saida, "Opção inválida.")
		}
	}
}

func (t Terminal) pedirComentario() (string, error) {
	for {
		fmt.Fprint(t.Saida, "O que precisa mudar? > ")
		comentario, err := lerLinha(t.Entrada)
		if err != nil || comentario != "" {
			return comentario, err
		}
		fmt.Fprintln(t.Saida, "Ajuste sem comentário não ajuda a próxima tentativa; escreva o que mudar.")
	}
}

const marcaBOM = "\uFEFF"

func lerLinha(r *bufio.Reader) (string, error) {
	linha, err := r.ReadString('\n')
	if err != nil && linha == "" {
		return "", fmt.Errorf("entrada encerrada antes da decisão do gate: %w", err)
	}
	return strings.TrimSpace(strings.TrimPrefix(linha, marcaBOM)), nil
}

func previa(conteudo string, linhas int) string {
	partes := strings.Split(strings.TrimRight(conteudo, "\n"), "\n")
	if len(partes) <= linhas {
		return strings.Join(partes, "\n")
	}
	return strings.Join(partes[:linhas], "\n") + fmt.Sprintf("\n… (+%d linhas no arquivo)", len(partes)-linhas)
}
