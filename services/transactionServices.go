package services

import (
	"bufio"
	"desafia-CNAB/model"
	"desafia-CNAB/repository"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

type TransactionServices struct {

	TransactionRepository repository.TypeRepository
}

//metado para pegar um arquivo e ler uma linha por vez, tranformando a linha em string
func (t *TransactionServices) ArchiveProcess(file multipart.File){

	Reader:= bufio.NewScanner(file)
	for Reader.Scan() {//enquanto estiver linhas, o loop continua

		line := Reader.Text()//string
		
		transfer:=parseLinha(line)
		t.TransactionRepository.Save(transfer)//salvando o resultado de "tranfer" no banco de dados

	}

	
}

//funcao para dividir a linha vinda da função anterior em varias substrings
func parseLinha(line string) model.Transaction {

	if len(line) < 2 {

		panic("linha não tem os elemetos esperados")
	}

	tipocodigo, err:= strconv.Atoi(string(line[0]))
	if err != nil {
		panic("Erro ao converter o primeiro valor")//transformando a substrings em inteiro
	}

	datasrv, err:= time.Parse("20060102", line[1:9])
	if err != nil {
		fmt.Errorf("erro ao converter a data: %v", err)//transformando a substrings em data
	}

	value, err:= strconv.ParseFloat(line[9:19], 64)
	if err != nil {
		fmt.Errorf("erro ao converter o valor: %v", err)//transformando a substrings em float
	}

	horasrv, err:= time.Parse("150405", line[42:48])//transformando a substrings em hora
	if err != nil {
		fmt.Errorf("erro ao converter a hora: %v", err)
	}


	var tp model.TransferType

	tp.TypeGetById(tipocodigo)

	Transaction:= model.Transaction {
		Type: tp.Description,
		Nature: string(tp.Nature),
		Signal: model.Status(tp.Signal),
		Date: datasrv,
		Value: value/100,
		Cpf: line[19:30],
		Card: line[30:42],
		Time: horasrv,
		Store_owner: line[48:62],
		Store_name: line[62:81],

	}
	
	return Transaction
}

func (t *TransactionServices)ListTransactions(){

	t.TransactionRepository.FindAll()
}