package model

import "errors"

type Status string

const (
	entrada = "entrada"
	saida   = "saída"
) //const para padronizar a escrita

type TransferType struct {
	Type        int
	Description string
	Nature      Status
	Signal      string
} //struct da transferencia

var Transfer = []TransferType{ //slice de todas as tranferencias

	{1, "debito", entrada, "+"},
	{2, "boleto", saida, "-"},
	{3, "financiamento", saida, "-"},
	{4, "credito", entrada, "+"},
	{5, "recebimento emprestimo", entrada, "+"},
	{6, "vendas", entrada, "+"},
	{7, "recebimento TED", entrada, "+"},
	{8, "recebimento DOC", entrada, "+"},
	{9, "aluguel saida", saida, "-"},
}

// funcao para pega cada tipo de tranferencia pelo id
func (tp *TransferType) TypeGetById(Type int) ( error) {

	for _, t := range Transfer {

		if t.Type == Type {
			
			*tp = t
			return nil
		} 
					
	}
	return  errors.New("invalid type!")

}
