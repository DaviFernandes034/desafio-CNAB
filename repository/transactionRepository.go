package repository

import (
	"database/sql"
	"desafia-CNAB/model"
	"fmt"
)

type TypeRepository interface {
	GetById(id int) (*model.TransferType, error)
	Save(transactin model.Transaction) error
	FindAll() ([]model.Transaction, error)
} //interface para typeRepository implementar

type typeRepository struct { //conexão com o banco de dados

	db *sql.DB
}



// essa estrutura que irá fazer as querys
func NewtypeRepository(db *sql.DB) TypeRepository {

	return &typeRepository{db: db}
}

// GetById implementado do typeRepository.
// A estrutura typeRepository implementa a interface TypeRepository porque ela tem o método GetById
func (t *typeRepository) GetById(id int) (*model.TransferType, error) {
	panic("unimplemented")
}

// Save implements TypeRepository.
func (t *typeRepository) Save(transactin model.Transaction) error {

	querys := `insert into CNAB (Tipo, Natureza, Sinal, Data, Valor,
		 Cpf, Cartao, Hora, Dono_loja, Nome_Loja)
		 values(@tipo, @natureza, @sinal, @data, @valor, @cpf, @cartao, @hora, @dono_loja, @nome_loja)`

	stmt, err := t.db.Prepare(querys)
	if err != nil {
		return fmt.Errorf("erro ao preparar a query: %v", err)
	}

	_, err = stmt.Exec(

		sql.Named("tipo", transactin.Type),
		sql.Named("natureza", transactin.Nature),
		sql.Named("sinal", transactin.Signal),
		sql.Named("data", transactin.Date),
		sql.Named("valor", transactin.Value),
		sql.Named("cpf", transactin.Cpf),
		sql.Named("cartao", transactin.Card),
		sql.Named("hora", transactin.Time),
		sql.Named("dono_loja", transactin.Store_owner),
		sql.Named("nome_loja", transactin.Store_name),
	)

	if err != nil {

		return fmt.Errorf("erro no stmt: %v", err)
	}

	return nil
}


// FindAll implements TypeRepository.(fazer outro dia)
func (t *typeRepository) FindAll() ([]model.Transaction, error) {
	
	query:= `select * from CNAB`

	rows, err:= t.db.Query(query)
	if err != nil {

		return nil, fmt.Errorf("erro ao buscar as transações: %v", err)
	}
	defer rows.Close()

	var Transactions []model.Transaction

	for rows.Next(){

		var Transaction model.Transaction

		if err:= rows.Scan(&Transaction.Type, &Transaction.Nature, &Transaction.Signal,
			&Transaction.Date, &Transaction.Value, &Transaction.Cpf,&Transaction.Card,&Transaction.Time,
			&Transaction.Store_owner, &Transaction.Store_name); err != nil{

			return nil, fmt.Errorf("erro ao ler a linha: %v", err)
		}

		Transactions = append(Transactions, Transaction)
	}

	if err:= rows.Err(); err != nil {

		return nil, fmt.Errorf("erro ao iterar sobre os resultados: %v", err)
	}

	return Transactions, nil
}