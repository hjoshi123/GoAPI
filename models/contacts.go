package models

import (
	"fmt"

	u "github.com/hjoshi123/GORest/utils"
	"github.com/jinzhu/gorm"
)

// Contact structure provides the outline of the model
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

// Validate returns message if validation of parameters is true
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name should be provided"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone Number should be provided"), false
	}

	if contact.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

// Create an user
func (contact *Contact) Create() map[string]interface{} {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "Success")
	resp["Contact"] = contact
	return resp
}

// GetContact gets the contact of the user
func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

// GetContacts returns the array of contacts
func GetContacts(user uint) []*Contact {

	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
