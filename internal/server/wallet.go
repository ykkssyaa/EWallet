package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *HttpServer) CreateWallet(w http.ResponseWriter, r *http.Request) {

	s.logger.Info.Println("Invoked CreateWallet on server")

	wallet, err := s.services.WalletService.CreateWallet()

	if err != nil {

		s.logger.Err.Println("Error with creating wallet: " + err.Error())

		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.logger.Info.Println(fmt.Sprintf("Created wallet with id: %s, balance: %f", wallet.Id, wallet.Balance))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(wallet)
}

func (s *HttpServer) GetWallet(w http.ResponseWriter, r *http.Request) {

	walletId := mux.Vars(r)["walletId"]

	s.logger.Info.Println("Invoked GetWallet on server. walletId = " + walletId)

	wallet, err := s.services.WalletService.GetWalletById(walletId)

	if err != nil {
		s.logger.Err.Println("Error with getting wallet: " + err.Error())
		errorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	if len(wallet.Id) == 0 {
		errorMes := fmt.Sprintf("Wallet with id %s not found", walletId)
		s.logger.Err.Println(errorMes)
		errorResponse(w, errorMes, http.StatusNotFound)
		return
	}

	s.logger.Info.Println(fmt.Sprintf("Wallet with id %s has balance %f", wallet.Id, wallet.Balance))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(wallet)
}
