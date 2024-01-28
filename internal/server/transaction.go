package server

import (
	"EWallet/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *HttpServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	walletId := mux.Vars(r)["walletId"]
	s.logger.Info.Println("Invoked CreateTransaction on server. walletId = " + walletId)

	if r.Header.Get("Content-Type") != "application/json" {
		s.logger.Err.Println("Error: Content Type is not application/json")
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var transaction model.Transaction
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&transaction)

	if err != nil {
		s.logger.Err.Println("Error while decoding request body " + err.Error())
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	s.logger.Info.Println("Request body: ", transaction)

	if len(transaction.To) == 0 || transaction.Amount == 0 {

		errorMes := "Error: Empty fields in transaction request"
		s.logger.Err.Println(errorMes)
		errorResponse(w, errorMes, http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		errorMes := "Error: Amount must be positive number"
		s.logger.Err.Println(errorMes)
		errorResponse(w, errorMes, http.StatusBadRequest)
		return
	}

	transaction.From = walletId

	if err := s.services.TransactionService.CreateTransaction(transaction); err != nil {
		s.logger.Err.Println(err.Error())
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) GetHistoryOfTransactions(w http.ResponseWriter, r *http.Request) {

	walletId := mux.Vars(r)["walletId"]
	s.logger.Info.Println("Invoked GetHistoryOfTransactions on server. walletId = " + walletId)

	transactions, err := s.services.TransactionService.GetAllTransactionsByWallet(walletId)

	if err != nil {
		s.logger.Err.Println(err.Error())
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(transactions)
}
