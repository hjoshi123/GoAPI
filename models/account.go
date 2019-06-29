package models

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	u "github.com/hjoshi123/GORest/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token is a struct that will be encoded to JWT. jwt.StandardClaims is added to include
// details like expiry etc
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Account represents user account
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Validate users details
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 && !strings.ContainsAny(account.Password, "!@#$%^&*") {
		return u.Message(false, "Password should be more than 6 characters and should have special characters"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Validation Passed"), true
}

// Create an account for user
func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error")
	}

	tokenInstance := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenInstance)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

// Login allows the users to login
func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		//Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

// GetUser returns the user from id
func GetUser(uid uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", uid).First(acc)
	if acc.Email == "" {
		//User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
