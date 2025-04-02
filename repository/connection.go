package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)


func Connection()(*sql.DB, error){

	db_server:=os.Getenv("DB_SERVER")
	port:=os.Getenv("DB_PORT")
	database:=os.Getenv("DATABASE")


	//string de conexão

	stringConn:= fmt.Sprintf("server=%s;port=%s;database=%s;trusted_connection=yes", db_server,port,
	database)

	//abrindo a conexão

	db, err:= sql.Open("sqlserver",stringConn)
	if err != nil{
		return nil, fmt.Errorf("erro ao abrir o banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil ,fmt.Errorf("erro ao conectar com o banco de dados", err)

	}

	fmt.Println("conexao estabelecida")

	return db, nil

}