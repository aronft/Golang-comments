package controllers

import (
	"context"
	"net/http"

	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/commons"
	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// ValidateToken validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)

	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token ha expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del token coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "su token no es válido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}

	if token.Valid {
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		next(w, r.WithContext(ctx))
	} else {
		m.Code = http.StatusUnauthorized
		m.Message = "su token no es válido"
		commons.DisplayMessage(w, m)
	}

}
