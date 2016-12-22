package common

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//VerifyKey Key for verifying the JWTs
var VerifyKey *rsa.PublicKey
var signKey *rsa.PrivateKey

const (
	privKeyFileName = "app.rsa"
	pubKeyFileName  = "app.rsa.pub"
)

func initKeys(keysDir string) {
	var err error

	privKeyPath := fmt.Sprintf("%s/%s", keysDir, privKeyFileName)
	pubKeyPath := fmt.Sprintf("%s/%s", keysDir, pubKeyFileName)

	privBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		panic(err)
	}

	pubBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		panic(err)
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		panic(err)
	}
}

//GenerateJWT Generates a JWT based on the user
func GenerateJWT() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "ProjectJane",
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

type authHandler struct {
	next http.HandlerFunc
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorize(w, r, h.next)
}

func authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// var claims models.TokenClaims

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				displayAppError(w, err, "Access Token is expired", 401)
				return
			default:
				displayAppError(w, err, "Error while parsing Access Token", 500)
				return
			}
		default:
			displayAppError(w, err, "Error while parsing Access Token", 500)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		displayAppError(w, err, "Invalid Access Token", 401)
	}
}

//MustAuth Middleware Handler for enforcing authentication
func MustAuth(handler http.HandlerFunc) http.Handler {
	return &authHandler{next: handler}
}

func displayAppError(w http.ResponseWriter, err error, message string, statusCode int) {
	http.Error(w, message+" : "+err.Error(), statusCode)
}
