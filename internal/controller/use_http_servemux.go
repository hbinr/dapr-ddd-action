package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (u UserController) SayHi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	_ = json.NewEncoder(w).Encode(vars["world"])
}
