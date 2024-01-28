package service

import (
	"EWallet/internal/gateway"
	"EWallet/internal/model"
	"EWallet/pkg/logger"
)

type WalletService interface {
	CreateWallet() (model.Wallet, error)
	GetWalletById(id string) (model.Wallet, error)
}

type WalletServiceImpl struct {
	walletGateway gateway.WalletGateway
	logger        *logger.Logger
}

func (w WalletServiceImpl) CreateWallet() (model.Wallet, error) {
	return w.walletGateway.CreateWallet()
}

func (w WalletServiceImpl) GetWalletById(id string) (model.Wallet, error) {
	return w.walletGateway.GetWalletById(id)
}
