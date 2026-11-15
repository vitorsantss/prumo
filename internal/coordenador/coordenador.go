package coordenador

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/vitorsantss/prumo/internal/executor"
	"github.com/vitorsantss/prumo/internal/gate"
	"github.com/vitorsantss/prumo/internal/pipeline"
	"github.com/vitorsantss/prumo/internal/rastro"
	"github.com/vitorsantss/prumo/internal/tarefa"
)

const (
	statusAprovada  = "aprovada"
	statusErro      = "erro"
	statusEsgotada  = "tentativas_esgotadas"
	statusConcluida = "concluida"
	origemDeclarada = "entradas_declaradas"
)

type Coordenador struct {
	IA            executor.Executor
	Modelo        executor.Executor
	Decisor       gate.Decisor
	Saida         io.Writer
	Agora         func() time.Time
	MaxTentativas int
}

func (c Coordenador) Executar(ctx context.Context, p pipeline.Pipeline, t tarefa.Tarefa) (rastro.Tarefa, error) {
	resumo := rastro.Tarefa{
		TraceID:   rastro.TraceID(t.ID),
		ID:        t.ID,
		Titulo:    t.Titulo,
		Workspace: t.Workspace,
		Pipeline:  p.Nome,
		Inicio:    c.agora(),
		Status:    statusConcluida,
	}
	var falha error
	for i, etapa := range p.Etapas {
		fmt.Fprintf(c.Saida, "▶ %d/%d %s\n", i+1, len(p.Etapas), etapa.Nome)
		r, err := c.executarEtapa(ctx, i+1, etapa, t, resumo.TraceID)
		caminho := filepath.Join(t.PastaEtapa(i+1, etapa.ID), "etapa.json")
		if errGravar := rastro.GravarJSON(caminho, r); errGravar != nil && err == nil {
			err = errGravar
		}
		resumo.Etapas = append(resumo.Etapas, rastro.ResumoEtapa{
			Ordem: r.Ordem, EtapaID: r.EtapaID, Nome: r.Nome, Status: r.Status,
			DuracaoMS: r.DuracaoMS, Rastro: filepath.ToSlash(mustRel(t.Pasta, caminho)),
		})
		resumo.TokensEntrada += inteiro(r.Atributos[rastro.AtributoTokensEntrada])
		resumo.TokensSaida += inteiro(r.Atributos[rastro.AtributoTokensSaida])
		resumo.CustoUSD += decimal(r.Atributos[rastro.AtributoCustoUSD])
		if err != nil {
			resumo.Status = r.Status
			falha = fmt.Errorf("etapa %s: %w", etapa.ID, err)
			break
		}
	}
	resumo.Fim = c.agora()
	if err := rastro.GravarJSON(t.Caminho("tarefa.json"), resumo); err != nil && falha == nil {
		falha = err
	}
	return resumo, falha
}

func (c Coordenador) executarEtapa(ctx context.Context, ordem int, etapa pipeline.Etapa, t tarefa.Tarefa, traceID string) (rastro.Etapa, error) {
	r := rastro.Etapa{
		TraceID: traceID, SpanID: rastro.NovoSpanID(), Ordem: ordem, EtapaID: etapa.ID, Nome: etapa.Nome,
		Inicio: c.agora(), OrigemLeituras: origemDeclarada, Atributos: map[string]any{}, Gates: []rastro.Gate{},
	}
	finalizar := func(status string, err error) (rastro.Etapa, error) {
		r.Fim = c.agora()
		r.DuracaoMS = r.Fim.Sub(r.Inicio).Milliseconds()
		r.Status = status
		if err != nil {
			r.Erro = err.Error()
		}
		return r, err
	}

	entradas, lidos, err := lerEntradas(etapa, t)
	if err != nil {
		return finalizar(statusErro, err)
	}
	r.Lidos = lidos

	exec, decisor := c.IA, c.Decisor
	if etapa.Tipo == pipeline.TipoModelo {
		exec = c.Modelo
	}
	if etapa.Gate == pipeline.GateAutomatico {
		decisor = gate.Automatico{}
	}

	ajuste := ""
	for tentativa := 1; tentativa <= c.maxTentativas(); tentativa++ {
		r.Tentativas = tentativa
		res, err := exec.Executar(ctx, executor.Pedido{Etapa: etapa, Entradas: entradas, Ajuste: ajuste, Pasta: t.Pasta})
		if err != nil {
			return finalizar(statusErro, err)
		}
		registrarUso(r.Atributos, res.Uso)
		if err := c.registrarUsoNaTarefa(t, etapa, tentativa, res.Uso); err != nil {
			return finalizar(statusErro, err)
		}
		caminhoSaida := t.Caminho(etapa.Saida)
		if err := os.WriteFile(caminhoSaida, []byte(res.Conteudo), 0o644); err != nil {
			return finalizar(statusErro, err)
		}
		escrito, err := rastro.Referencia(caminhoSaida, t.Pasta)
		if err != nil {
			return finalizar(statusErro, err)
		}
		r.Escritos = []rastro.Arquivo{escrito}

		d, err := decisor.Decidir(etapa, caminhoSaida, res.Conteudo)
		if err != nil {
			return finalizar(statusErro, err)
		}
		g := rastro.Gate{Tentativa: tentativa, Decisao: string(d.Tipo), Comentario: d.Comentario, Em: c.agora()}
		r.Gates = append(r.Gates, g)
		if err := rastro.AcrescentarLinha(t.Caminho("gates.jsonl"), struct {
			Etapa string `json:"etapa"`
			rastro.Gate
		}{etapa.ID, g}); err != nil {
			return finalizar(statusErro, err)
		}
		if d.Tipo == gate.Aprovar {
			return finalizar(statusAprovada, nil)
		}
		ajuste = d.Comentario
		fmt.Fprintf(c.Saida, "↻ %s: tentativa %d\n", etapa.Nome, tentativa+1)
	}
	return finalizar(statusEsgotada, errors.New("limite de tentativas atingido sem aprovação"))
}

func (c Coordenador) registrarUsoNaTarefa(t tarefa.Tarefa, etapa pipeline.Etapa, tentativa int, u executor.Uso) error {
	return rastro.AcrescentarLinha(t.Caminho("usage.jsonl"), map[string]any{
		"em":                         c.agora(),
		"etapa":                      etapa.ID,
		"tentativa":                  tentativa,
		rastro.AtributoProvedor:      u.Provedor,
		rastro.AtributoModelo:        u.Modelo,
		rastro.AtributoTokensEntrada: u.TokensEntrada,
		rastro.AtributoTokensSaida:   u.TokensSaida,
		rastro.AtributoCustoUSD:      u.CustoUSD,
	})
}

func lerEntradas(etapa pipeline.Etapa, t tarefa.Tarefa) ([]executor.Entrada, []rastro.Arquivo, error) {
	entradas := make([]executor.Entrada, 0, len(etapa.Entradas))
	lidos := make([]rastro.Arquivo, 0, len(etapa.Entradas))
	for _, nome := range etapa.Entradas {
		caminho := t.Caminho(nome)
		conteudo, err := os.ReadFile(caminho)
		if err != nil {
			return nil, nil, fmt.Errorf("entrada %s: %w", nome, err)
		}
		ref, err := rastro.Referencia(caminho, t.Pasta)
		if err != nil {
			return nil, nil, err
		}
		entradas = append(entradas, executor.Entrada{Nome: nome, Conteudo: string(conteudo)})
		lidos = append(lidos, ref)
	}
	return entradas, lidos, nil
}

func registrarUso(atributos map[string]any, u executor.Uso) {
	atributos[rastro.AtributoProvedor] = u.Provedor
	atributos[rastro.AtributoModelo] = u.Modelo
	atributos[rastro.AtributoTokensEntrada] = inteiro(atributos[rastro.AtributoTokensEntrada]) + u.TokensEntrada
	atributos[rastro.AtributoTokensSaida] = inteiro(atributos[rastro.AtributoTokensSaida]) + u.TokensSaida
	atributos[rastro.AtributoCustoUSD] = decimal(atributos[rastro.AtributoCustoUSD]) + u.CustoUSD
}

func (c Coordenador) agora() time.Time {
	if c.Agora != nil {
		return c.Agora()
	}
	return time.Now().UTC()
}

func (c Coordenador) maxTentativas() int {
	if c.MaxTentativas > 0 {
		return c.MaxTentativas
	}
	return 5
}

func inteiro(v any) int {
	n, _ := v.(int)
	return n
}

func decimal(v any) float64 {
	f, _ := v.(float64)
	return f
}

func mustRel(base, caminho string) string {
	rel, err := filepath.Rel(base, caminho)
	if err != nil {
		return caminho
	}
	return rel
}
