package main

import (
	
	"desafia-CNAB/repository"
	"fmt"

	"desafia-CNAB/controller"
	"desafia-CNAB/services"

	"github.com/gin-gonic/gin"
)


func main(){


	db, err:= repository.Init()
	if err != nil{
		fmt.Errorf("erro ao conectar com o banco de dados: %v", err.Error())
	}
	defer db.Close()
	router:= gin.Default()

	tr:= repository.NewtypeRepository(db)
	ts:=services.TransactionServices{
		TransactionRepository: tr,
	}
	cb:= &controller.CnabController{
		TransactionServices: ts,
	}


	router.POST("/upload", cb.UploadFile())
	router.GET("/getAll", cb.GetAll())

	router.Run(":8080")


}