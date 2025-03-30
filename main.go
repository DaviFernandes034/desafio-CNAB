package main

import (
	"desafia-CNAB/model"
	_"desafia-CNAB/repository"
	"fmt"
)


func main(){


	//repository.Init()

	transfer:= model.TypeGetById(1)
	fmt.Println(transfer.Type, transfer.Description, transfer.Nature, transfer.Signal)

}