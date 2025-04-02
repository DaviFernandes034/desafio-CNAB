package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Init()(*sql.DB,error){

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	
	}

		// Conectar ao banco de dados
		db, err := Connection()
		if err != nil {
			return nil, fmt.Errorf("erro ao conectar com o banco de dados: %v", err.Error())
		}
		
	
		log.Println("Conexão com o banco de dados estabelecida com sucesso")
		

		return db, nil
	
}