package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// ClaudeCode roda a etapa com o Claude Code em modo não interativo
// (claude -p). É o adaptador terceirizado; o caminho livre é a issue #13.
type ClaudeCode struct {
	Comando string
	Modelo  string
}

// respostaClaude lê só o que o Prumo usa. Identificadores de sessão e dados de
// credencial ficam de fora de propósito: nunca são gravados.
type respostaClaude struct {
	Result       string  `json:"result"`
	IsError      bool    `json:"is_error"`
	TotalCostUSD float64 `json:"total_cost_usd"`
	NumTurns     int     `json:"num_turns"`
	DurationMS   int64   `json:"duration_ms"`
	Usage        struct {
		InputTokens              int `json:"input_tokens"`
		OutputTokens             int `json:"output_tokens"`
		CacheReadInputTokens     int `json:"cache_read_input_tokens"`
		CacheCreationInputTokens int `json:"cache_creation_input_tokens"`
	} `json:"usage"`
	ModelUsage map[string]usoModeloClaude `json:"modelUsage"`
}

type usoModeloClaude struct {
	InputTokens              int     `json:"inputTokens"`
	OutputTokens             int     `json:"outputTokens"`
	CacheReadInputTokens     int     `json:"cacheReadInputTokens"`
	CacheCreationInputTokens int     `json:"cacheCreationInputTokens"`
	CostUSD                  float64 `json:"costUSD"`
	CostBasis                string  `json:"costBasis"`
}

const provedorClaude = "anthropic"

func (c ClaudeCode) Executar(ctx context.Context, p Pedido) (Resultado, error) {
	comando := c.Comando
	if comando == "" {
		comando = "claude"
	}
	args := []string{"-p", "--output-format", "json"}
	if c.Modelo != "" {
		args = append(args, "--model", c.Modelo)
	}
	cmd := exec.CommandContext(ctx, comando, args...)
	cmd.Dir = p.Pasta
	cmd.Stdin = strings.NewReader(montarPrompt(p))
	var saida, erro bytes.Buffer
	cmd.Stdout, cmd.Stderr = &saida, &erro
	if err := cmd.Run(); err != nil {
		var naoAchou *exec.Error
		if errors.As(err, &naoAchou) {
			return Resultado{}, fmt.Errorf("comando %q não encontrado: instale o Claude Code ou use --executor duble", comando)
		}
		return Resultado{}, fmt.Errorf("claude falhou: %w: %s", err, strings.TrimSpace(erro.String()))
	}
	return interpretarRespostaClaude(saida.Bytes())
}

func interpretarRespostaClaude(bruto []byte) (Resultado, error) {
	var r respostaClaude
	if err := json.Unmarshal(bruto, &r); err != nil {
		return Resultado{}, fmt.Errorf("resposta do claude ilegível: %w", err)
	}
	if r.IsError {
		return Resultado{}, fmt.Errorf("claude devolveu erro: %s", r.Result)
	}
	return Resultado{
		Conteudo:  r.Result,
		Usos:      usosPorModelo(r),
		Voltas:    r.NumTurns,
		DuracaoIA: time.Duration(r.DurationMS) * time.Millisecond,
	}, nil
}

func usosPorModelo(r respostaClaude) []Uso {
	if len(r.ModelUsage) == 0 {
		u := r.Usage
		return []Uso{{
			Provedor:           provedorClaude,
			TokensEntrada:      u.InputTokens + u.CacheReadInputTokens + u.CacheCreationInputTokens,
			TokensLeituraCache: u.CacheReadInputTokens,
			TokensCriacaoCache: u.CacheCreationInputTokens,
			TokensSaida:        u.OutputTokens,
			CustoUSD:           r.TotalCostUSD,
		}}
	}
	modelos := make([]string, 0, len(r.ModelUsage))
	for m := range r.ModelUsage {
		modelos = append(modelos, m)
	}
	sort.Strings(modelos)
	usos := make([]Uso, 0, len(modelos))
	for _, m := range modelos {
		u := r.ModelUsage[m]
		usos = append(usos, Uso{
			Provedor:           provedorClaude,
			Modelo:             m,
			TokensEntrada:      u.InputTokens + u.CacheReadInputTokens + u.CacheCreationInputTokens,
			TokensLeituraCache: u.CacheReadInputTokens,
			TokensCriacaoCache: u.CacheCreationInputTokens,
			TokensSaida:        u.OutputTokens,
			CustoUSD:           u.CostUSD,
			BaseCusto:          u.CostBasis,
		})
	}
	return usos
}
