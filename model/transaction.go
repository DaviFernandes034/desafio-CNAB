package model

import (
	"time"

	
)

type Transaction struct {
	Type   string
	Nature string
	Signal Status
	Date   time.Time
	Value float64
	Cpf string
	Card string
	Time string
	Store_owner string
	Store_name string

}