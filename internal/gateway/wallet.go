package gateway

import (
	"EWallet/internal/model"
	lg "EWallet/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type WalletGateway interface {
	CreateWallet(wallet model.Wallet) (string, error)
	GetWalletById(id string) (model.Wallet, error)
}

type WalletGatewayImpl struct {
	db     *sqlx.DB
	logger *lg.Logger
}

func (w WalletGatewayImpl) CreateWallet(wallet model.Wallet) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletGatewayImpl) GetWalletById(id string) (model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}
