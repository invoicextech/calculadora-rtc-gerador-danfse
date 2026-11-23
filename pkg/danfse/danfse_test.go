package danfse

import (
	"context"
	"testing"
)

func TestNovoGeradorInicializaOServicoDaReceita(t *testing.T) {
	g, err := NovoGerador()
	if err != nil {
		t.Fatalf("NovoGerador: %v", err)
	}
	if g == nil {
		t.Fatal("gerador nulo")
	}
}

func TestGerarRejeitaXMLInvalido(t *testing.T) {
	g, err := NovoGerador()
	if err != nil {
		t.Fatalf("NovoGerador: %v", err)
	}
	if _, err := g.Gerar(context.Background(), []byte("<nao-e-nfse/>"), Valida, true); err == nil {
		t.Fatal("esperava erro para XML que não é NFS-e")
	}
}

func TestGerarRejeitaEntradaVazia(t *testing.T) {
	g, err := NovoGerador()
	if err != nil {
		t.Fatalf("NovoGerador: %v", err)
	}
	if _, err := g.Gerar(context.Background(), nil, Cancelada, false); err == nil {
		t.Fatal("esperava erro para XML vazio")
	}
}
