# Fork da InvoiceX

Este repositório é um fork do [gerador oficial do DANFSe da Receita Federal](https://github.com/ReceitaFederal/calculadora-rtc-gerador-danfse).

O código da Receita não é alterado. O fork acrescenta apenas:

- `pkg/danfse`: pacote público que expõe o gerador da Receita (`NovoGerador` e `Gerar`) para ser importado por outros módulos Go, já que o código original fica em `internal/`;
- `.github/workflows/sync-upstream.yml`: workflow diário que traz as mudanças da Receita, roda os testes e abre um PR de sincronização.

## Uso

```go
import "danfse-service/pkg/danfse"

gerador, err := danfse.NovoGerador()
pdf, err := gerador.Gerar(ctx, xml, danfse.Valida, true)
```

No `go.mod` de quem usa:

```
require danfse-service v0.0.0-00010101000000-000000000000

replace danfse-service => github.com/invoicextech/calculadora-rtc-gerador-danfse <versão>
```
