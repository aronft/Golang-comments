package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim es una solicitud
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
