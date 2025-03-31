package repository

import (
	"database/sql"
	"desafia-CNAB/model"
)

type TypeRepository interface {
	GetById(id int) (*model.TransferType, error)
} //interface para typeRepository implementar

type typeRepository struct { //conexão com o banco de dados

	db *sql.DB
}//essa estrutura que irá fazer as querys


func NewtypeRepository(db *sql.DB) TypeRepository {

	return &typeRepository{db: db}
}


// GetById implementado do typeRepository.
//A estrutura typeRepository implementa a interface TypeRepository porque ela tem o método GetById
func (t *typeRepository) GetById(id int) (*model.TransferType, error) {
	panic("unimplemented")
}
