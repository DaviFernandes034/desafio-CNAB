package model


type status string

const (
	entrada = "entrada"
	saida = "saida"
)
type TransferType struct {

	Type int
	Description string
	Nature status
	Signal string


}

var Transfer = []TransferType{

	{1, "debito", entrada, "+"},
	{2, "boleto", saida, "-"},
	{3, "financiamento",saida,"-"},
	{4, "credito", entrada, "+"},
	{5,"recebimento emprestimo",entrada, "+"},
	{6, "vendas", entrada, "+"},
	{7, "recebimento TED", entrada, "+"},
	{8, "recebimento DOC", entrada, "+"},
	{9, "aluguel saida", saida, "-"},
}

func TypeGetById(Type int) (*TransferType){

	for _, t:= range Transfer {

		if t.Type == Type {
			return &t
		}
	}

	return nil

}



