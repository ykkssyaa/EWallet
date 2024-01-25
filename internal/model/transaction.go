package model

import "time"

type Transaction struct {
	Time   time.Time `json:"time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

type TransactionRequest struct {
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
