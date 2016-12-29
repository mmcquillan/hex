package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

//TokenClaims Struct for representing a users claims for JWT
type TokenClaims struct {
	UserID uuid.Uuid `json:"userid"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.StandardClaims
}
