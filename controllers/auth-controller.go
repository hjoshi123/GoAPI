package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hjoshi123/GORest/models"
	u "github.com/hjoshi123/GORest/utils"
)

// CreateAccount is Controller for creating account
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Response(w, u.Message(false, "Invalid Request"))
		return
	}

	response := account.Create()
	u.Response(w, response)
}

// Auth handles Authentication of the suer
var Auth = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Response(w, u.Message(false, "Invalid Request"))
		return
	}

	response := models.Login(account.Email, account.Password)
	u.Response(w, response)

}
