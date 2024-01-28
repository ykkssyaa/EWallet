package model

const DefaultBalance float32 = 100

type Wallet struct {
	Id      string  `json:"id" db:"id"`
	Balance float32 `json:"balance" db:"balance"`
}
