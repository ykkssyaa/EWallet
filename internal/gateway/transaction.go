package gateway

import (
	"EWallet/internal/model"
	lg "EWallet/pkg/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionGateway interface {
	GetAllTransactionsByWallet(id string) ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) error
}

type TransactionGatewayImpl struct {
	db     *sqlx.DB
	logger *lg.Logger
}

func (t TransactionGatewayImpl) GetAllTransactionsByWallet(id string) ([]model.Transaction, error) {

	t.logger.Info.Println(fmt.Sprintf("Getting all transaction of wallet with id %s  in postgres", id))

	var transactions []model.Transaction

	getQuery := "SELECT * FROM transactions WHERE from_wallet = $1 OR to_wallet = $1 ORDER BY time"

	if err := t.db.Select(&transactions, getQuery, id); err != nil {
		return nil, err
	}

	return transactions, nil

}

func (t TransactionGatewayImpl) CreateTransaction(transaction model.Transaction) error {

	t.logger.Info.Println(fmt.Sprintf("Creating Transaction. to: %s, from: %s, amount:%f",
		transaction.To, transaction.From, transaction.Amount))

	createQuery := "INSERT INTO transactions (time, from_wallet, to_wallet, amount) VALUES (now(), $1, $2, $3)"
	updateToWalletQuery := "UPDATE wallets SET balance = balance + $1 WHERE id = $2"
	updateFromWalletQuery := "UPDATE wallets SET balance = balance - $1 WHERE id = $2"

	tx, err := t.db.Begin()

	if err != nil {
		return err
	}

	t.logger.Info.Println(fmt.Sprintf("Transfer %f from %s", transaction.Amount, transaction.From))

	_, err = tx.Exec(updateFromWalletQuery, transaction.Amount, transaction.From)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	t.logger.Info.Println(fmt.Sprintf("Transfer %f to %s", transaction.Amount, transaction.To))

	_, err = tx.Exec(updateToWalletQuery, transaction.Amount, transaction.To)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	t.logger.Info.Println(fmt.Sprintf("Saving transaction (%s  -(%f)->  %s)",
		transaction.From, transaction.Amount, transaction.To))

	_, err = tx.Exec(createQuery, transaction.From, transaction.To, transaction.Amount)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
