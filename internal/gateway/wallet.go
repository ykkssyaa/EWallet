package gateway

import (
	"EWallet/internal/model"
	lg "EWallet/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type WalletGateway interface {
	CreateWallet() (model.Wallet, error)
	GetWalletById(id string) (model.Wallet, error)
}

type WalletGatewayImpl struct {
	db     *sqlx.DB
	logger *lg.Logger
}

func (w WalletGatewayImpl) CreateWallet() (model.Wallet, error) {

	w.logger.Info.Println("Create wallet in postgres")

	var newWallet model.Wallet

	createQuery := "INSERT INTO Wallets (balance) VALUES ($1) RETURNING *"

	tx, err := w.db.DB.Begin()

	if err != nil {
		return model.Wallet{}, err
	}

	row := tx.QueryRow(createQuery, model.DefaultBalance)
	if err := row.Scan(&newWallet.Id, &newWallet.Balance); err != nil {
		tx.Rollback()
		return model.Wallet{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.Wallet{}, err
	}

	return newWallet, nil
}

func (w WalletGatewayImpl) GetWalletById(id string) (model.Wallet, error) {

	w.logger.Info.Println("Getting wallet by id: " + id + " in postgres")

	var wallet model.Wallet

	getQuery := "SELECT * FROM wallets WHERE id = $1"

	if err := w.db.Get(&wallet, getQuery, id); err != nil {
		return model.Wallet{}, err
	}

	return wallet, nil
}
