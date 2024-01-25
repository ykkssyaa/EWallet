package server

import (
	"EWallet/internal/service"
	"EWallet/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HttpServer struct {
	services *service.Services
	logger   *logger.Logger
}

func NewHttpServer(services *service.Services, logger *logger.Logger, addr string) *http.Server {
	server := &HttpServer{services: services, logger: logger}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet", server.CreateWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet/{walletId}", server.GetWallet).Methods("GET")
	r.HandleFunc("/api/v1/wallet/{walletId}/send", server.CreateTransaction).Methods("POST")
	r.HandleFunc("/api/v1/wallet/{walletId}/history", server.GetHistoryOfTransactions).Methods("GET")

	return &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           addr,
	}
}
