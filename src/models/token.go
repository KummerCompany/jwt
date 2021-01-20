package models

import "github.com/dgrijalva/jwt-go"

// Token of API
type Token struct {
	Data interface{}
	Exp  string `json:"exp"`
	jwt.StandardClaims
}

// TokenDefault is a generic structure
type TokenDefault struct {
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	Admin     bool   `json:"admin"`
	Exp       string `json:"exp"`
	jwt.StandardClaims
}
