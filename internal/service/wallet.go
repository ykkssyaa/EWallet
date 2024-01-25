package service

import (
	"EWallet/internal/gateway"
	"EWallet/internal/model"
	"EWallet/pkg/logger"
)

type WalletService interface {
	CreateWallet(wallet model.Wallet) (string, error)
	GetWalletById(id string) (model.Wallet, error)
}

type WalletServiceImpl struct {
	walletGateway gateway.WalletGateway
	logger        *logger.Logger
}

func (w WalletServiceImpl) CreateWallet(wallet model.Wallet) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletServiceImpl) GetWalletById(id string) (model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}
