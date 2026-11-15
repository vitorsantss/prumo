package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type ClaudeCode struct {
	Comando string
	Modelo  string
}

type respostaClaude struct {
	Result       string  `json:"result"`
	IsError      bool    `json:"is_error"`
	TotalCostUSD float64 `json:"total_cost_usd"`
	Usage        struct {
		InputTokens              int `json:"input_tokens"`
		OutputTokens             int `json:"output_tokens"`
		CacheReadInputTokens     int `json:"cache_read_input_tokens"`
		CacheCreationInputTokens int `json:"cache_creation_input_tokens"`
	} `json:"usage"`
	ModelUsage map[string]json.RawMessage `json:"modelUsage"`
}

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
	uso := Uso{
		Provedor:      "anthropic",
		Modelo:        modeloPrincipal(r.ModelUsage),
		TokensEntrada: r.Usage.InputTokens + r.Usage.CacheReadInputTokens + r.Usage.CacheCreationInputTokens,
		TokensSaida:   r.Usage.OutputTokens,
		CustoUSD:      r.TotalCostUSD,
	}
	return Resultado{Conteudo: r.Result, Uso: uso}, nil
}

func modeloPrincipal(uso map[string]json.RawMessage) string {
	for modelo := range uso {
		return modelo
	}
	return ""
}
