package danfse

import (
	"context"

	"danfse-service/internal/app"
	"danfse-service/internal/domain/enums"
	"danfse-service/internal/service"
)

type Situacao int

const (
	Valida      Situacao = 0
	Substituida Situacao = 1
	Cancelada   Situacao = 2
)

type Gerador struct {
	servico service.ServicoGeracaoPdf
}

func NovoGerador(maxInputSize ...int) (*Gerador, error) {
	servico, err := app.NewPdfService(maxInputSize...)
	if err != nil {
		return nil, err
	}
	return &Gerador{servico: servico}, nil
}

func (g *Gerador) Gerar(ctx context.Context, xml []byte, situacao Situacao, canhoto bool) ([]byte, error) {
	tipo, ok := enums.SituacaoNotaPorSTC(int(situacao))
	if !ok {
		tipo = enums.VALIDA
	}
	return g.servico.Executar(ctx, xml, tipo, canhoto)
}
