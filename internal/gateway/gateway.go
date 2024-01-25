package gateway

import (
	"EWallet/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Gateways struct {
	TransactionGateway TransactionGateway
	WalletGateway      WalletGateway
}

func NewGateway(db *sqlx.DB, logger *logger.Logger) *Gateways {
	return &Gateways{
		TransactionGateway: TransactionGatewayImpl{db: db, logger: logger},
		WalletGateway:      WalletGatewayImpl{db: db, logger: logger},
	}
}
