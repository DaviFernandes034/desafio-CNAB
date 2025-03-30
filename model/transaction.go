package model

import (
	"time"

	
)

type Transaction struct {
	Type   string
	Nature string
	Signal string
	Date   time.Time
	Value float64
	Cpf string
	Card string
	Time time.Time
	Store_owner string
	Store_name string

}