package service

import (
	"EWallet/internal/gateway"
	"EWallet/pkg/logger"
)

type Services struct {
	TransactionService TransactionService
	WalletService      WalletService
}

func NewService(gateways *gateway.Gateways, logger *logger.Logger) *Services {
	return &Services{
		TransactionService: TransactionServiceImpl{
			transactionGateway: gateways.TransactionGateway,
			walletGetter:       gateways.WalletGateway,
			logger:             logger,
		},
		WalletService: WalletServiceImpl{
			walletGateway: gateways.WalletGateway,
			logger:        logger,
		},
	}
}
