package service

import (
	"EWallet/internal/gateway"
	"EWallet/internal/model"
	"EWallet/pkg/logger"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	return t.transactionGateway.GetAllTransactionsByWallet(id)
}

func (t TransactionServiceImpl) CreateTransaction(transaction model.Transaction) error {

	t.logger.Info.Println("Executing CreateTransaction on service")

	if transaction.To == transaction.From {
		return errors.New("transferring money to yourself")
	}

	ToWallet, err := t.walletGetter.GetWalletById(transaction.To)

	if err != nil {
		return err
	}

	if len(ToWallet.Id) == 0 {
		return errors.New(fmt.Sprintf("Wallet with id %s not found", transaction.To))
	}

	t.logger.Info.Println(fmt.Sprintf("Wallet with id %s was found", transaction.To))

	FromWallet, err := t.walletGetter.GetWalletById(transaction.From)

	if err != nil {
		return err
	}

	if len(FromWallet.Id) == 0 {
		return errors.New(spew.Sprintf("Wallet with id %s not found", transaction.From))
	}

	t.logger.Info.Println(fmt.Sprintf("Wallet with id %s was found", transaction.From))

	if FromWallet.Balance < transaction.Amount {
		return errors.New(spew.Sprintf("There are not enough funds in the wallet (%s) for the transaction", FromWallet.Id))
	}

	t.logger.Info.Println(fmt.Sprintf("There are enough funds in the wallet %s", transaction.From))

	return t.transactionGateway.CreateTransaction(transaction)
}
