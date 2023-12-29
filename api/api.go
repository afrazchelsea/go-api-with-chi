package api 

import (
	"net/http"
	"encoding/json"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string 
}

// Error printing function
func writeError(w http.ResponseWriter, message string, code int){
	res := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}

// Different types of errors which ultimately call the above function
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalServiceErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}

)