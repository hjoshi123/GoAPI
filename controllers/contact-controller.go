package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hjoshi123/GORest/models"
	u "github.com/hjoshi123/GORest/utils"
)

// CreateContact creates contact for the user
var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	//Grab the id of the user that send the request
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Response(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserID = user
	response := contact.Create()
	u.Response(w, response)
}

// GetContactsForUser gets a contact for a certain user
var GetContactsForUser = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	response := u.Message(true, "success")
	response["data"] = data
	u.Response(w, response)
}
