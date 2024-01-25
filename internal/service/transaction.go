package service

import (
	"EWallet/internal/gateway"
	"EWallet/internal/model"
	"EWallet/pkg/logger"
)

type TransactionService interface {
	GetAllTransactionsByWallet(id string) ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) error
}

type WalletGetter interface {
	GetWalletById(id string) (model.Wallet, error)
}

type TransactionServiceImpl struct {
	transactionGateway gateway.TransactionGateway
	walletGetter       WalletGetter
	logger             *logger.Logger
}

func (t TransactionServiceImpl) GetAllTransactionsByWallet(id string) ([]model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t TransactionServiceImpl) CreateTransaction(transaction model.Transaction) error {
	//TODO implement me
	panic("implement me")
}
