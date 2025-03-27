package repository

import (

	"log"

	"github.com/joho/godotenv"
)

func Init(){

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	
	}

		// Conectar ao banco de dados
		db, err := Connection()
		if err != nil {
			log.Fatalf("Erro: %v", err)
		}
		
	
		log.Println("Conexão com o banco de dados estabelecida com sucesso")
		defer db.Close()
	
}