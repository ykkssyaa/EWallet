package model

import "time"

type Transaction struct {
	Id     string    `json:"id" db:"id"`
	Time   time.Time `json:"time" db:"time"`
	From   string    `json:"from" db:"from_wallet"`
	To     string    `json:"to" db:"to_wallet"`
	Amount float32   `json:"amount" db:"amount"`
}
