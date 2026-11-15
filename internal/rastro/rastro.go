package rastro

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	AtributoProvedor      = "gen_ai.provider.name"
	AtributoModelo        = "gen_ai.request.model"
	AtributoTokensEntrada = "gen_ai.usage.input_tokens"
	AtributoTokensSaida   = "gen_ai.usage.output_tokens"
	AtributoCustoUSD      = "prumo.custo_usd"
)

type Arquivo struct {
	Caminho string `json:"caminho"`
	Bytes   int64  `json:"bytes"`
	SHA256  string `json:"sha256"`
}

type Gate struct {
	Tentativa  int       `json:"tentativa"`
	Decisao    string    `json:"decisao"`
	Comentario string    `json:"comentario,omitempty"`
	Em         time.Time `json:"em"`
}

type Etapa struct {
	TraceID        string         `json:"trace_id"`
	SpanID         string         `json:"span_id"`
	Ordem          int            `json:"ordem"`
	EtapaID        string         `json:"etapa_id"`
	Nome           string         `json:"nome"`
	Inicio         time.Time      `json:"inicio"`
	Fim            time.Time      `json:"fim"`
	DuracaoMS      int64          `json:"duracao_ms"`
	Status         string         `json:"status"`
	Tentativas     int            `json:"tentativas"`
	OrigemLeituras string         `json:"origem_leituras"`
	Lidos          []Arquivo      `json:"lidos"`
	Escritos       []Arquivo      `json:"escritos"`
	Atributos      map[string]any `json:"atributos"`
	Gates          []Gate         `json:"gates"`
	Erro           string         `json:"erro,omitempty"`
}

type ResumoEtapa struct {
	Ordem     int    `json:"ordem"`
	EtapaID   string `json:"etapa_id"`
	Nome      string `json:"nome"`
	Status    string `json:"status"`
	DuracaoMS int64  `json:"duracao_ms"`
	Rastro    string `json:"rastro"`
}

type Tarefa struct {
	TraceID       string        `json:"trace_id"`
	ID            string        `json:"id"`
	Titulo        string        `json:"titulo"`
	Workspace     string        `json:"workspace"`
	Pipeline      string        `json:"pipeline"`
	Inicio        time.Time     `json:"inicio"`
	Fim           time.Time     `json:"fim"`
	Status        string        `json:"status"`
	TokensEntrada int           `json:"gen_ai.usage.input_tokens"`
	TokensSaida   int           `json:"gen_ai.usage.output_tokens"`
	CustoUSD      float64       `json:"prumo.custo_usd"`
	Etapas        []ResumoEtapa `json:"etapas"`
}

func TraceID(tarefaID string) string {
	soma := sha256.Sum256([]byte(tarefaID))
	return hex.EncodeToString(soma[:16])
}

func NovoSpanID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("sem fonte de aleatoriedade: %v", err))
	}
	return hex.EncodeToString(b)
}

func Referencia(caminho, relativoA string) (Arquivo, error) {
	f, err := os.Open(caminho)
	if err != nil {
		return Arquivo{}, err
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return Arquivo{}, err
	}
	rel, err := filepath.Rel(relativoA, caminho)
	if err != nil {
		rel = caminho
	}
	return Arquivo{Caminho: filepath.ToSlash(rel), Bytes: n, SHA256: hex.EncodeToString(h.Sum(nil))}, nil
}

func GravarJSON(caminho string, valor any) error {
	if err := os.MkdirAll(filepath.Dir(caminho), 0o755); err != nil {
		return err
	}
	conteudo, err := json.MarshalIndent(valor, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(caminho, append(conteudo, '\n'), 0o644)
}

func AcrescentarLinha(caminho string, valor any) error {
	linha, err := json.Marshal(valor)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(caminho, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(append(linha, '\n'))
	return err
}
