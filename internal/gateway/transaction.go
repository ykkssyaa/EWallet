package gateway

import (
	"EWallet/internal/model"
	lg "EWallet/pkg/logger"
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
	//TODO implement me
	panic("implement me")
}

func (t TransactionGatewayImpl) CreateTransaction(transaction model.Transaction) error {
	//TODO implement me
	panic("implement me")
}
