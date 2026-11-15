package tarefa

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/vitorsantss/prumo/internal/pipeline"
)

type Tarefa struct {
	ID        string
	Workspace string
	Pasta     string
	Titulo    string
}

func (t Tarefa) Caminho(arquivo string) string {
	return filepath.Join(t.Pasta, arquivo)
}

func (t Tarefa) PastaEtapa(ordem int, etapaID string) string {
	return filepath.Join(t.Pasta, "etapas", fmt.Sprintf("%02d-%s", ordem, etapaID))
}

func Criar(dados, workspace, origem string, agora time.Time) (Tarefa, error) {
	conteudo, err := os.ReadFile(origem)
	if err != nil {
		return Tarefa{}, fmt.Errorf("lendo a tarefa de origem: %w", err)
	}
	titulo := tituloDe(conteudo, origem)
	id := agora.Format("20060102-150405") + "-" + slug(titulo)
	t := Tarefa{
		ID:        id,
		Workspace: workspace,
		Pasta:     filepath.Join(dados, slug(workspace), "tasks", id),
		Titulo:    titulo,
	}
	if err := os.MkdirAll(t.Pasta, 0o755); err != nil {
		return Tarefa{}, fmt.Errorf("criando a pasta da tarefa: %w", err)
	}
	if err := os.WriteFile(t.Caminho(pipeline.ArquivoTarefa), conteudo, 0o644); err != nil {
		return Tarefa{}, fmt.Errorf("copiando a tarefa: %w", err)
	}
	return t, nil
}

func tituloDe(conteudo []byte, origem string) string {
	for _, linha := range strings.Split(string(conteudo), "\n") {
		if titulo, ok := strings.CutPrefix(strings.TrimSpace(linha), "# "); ok && titulo != "" {
			return strings.TrimSpace(titulo)
		}
	}
	return strings.TrimSuffix(filepath.Base(origem), filepath.Ext(origem))
}

var naoPermitido = regexp.MustCompile(`[^a-z0-9]+`)

func slug(texto string) string {
	s := strings.ToLower(removerAcentos(texto))
	s = strings.Trim(naoPermitido.ReplaceAllString(s, "-"), "-")
	if len(s) > 40 {
		s = strings.TrimRight(s[:40], "-")
	}
	if s == "" {
		return "tarefa"
	}
	return s
}

var acentos = strings.NewReplacer(
	"á", "a", "à", "a", "â", "a", "ã", "a", "é", "e", "ê", "e", "í", "i",
	"ó", "o", "ô", "o", "õ", "o", "ú", "u", "ç", "c",
	"Á", "a", "À", "a", "Â", "a", "Ã", "a", "É", "e", "Ê", "e", "Í", "i",
	"Ó", "o", "Ô", "o", "Õ", "o", "Ú", "u", "Ç", "c",
)

func removerAcentos(s string) string {
	return acentos.Replace(s)
}
