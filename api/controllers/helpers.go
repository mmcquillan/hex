package controllers

import (
	"encoding/json"
	"linksscoreapi/models"
	"log"
	"net/http"

	"linksscoreapi/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// GetClaimsFromToken Gets the claims from the token
func GetClaimsFromToken(r *http.Request) (*models.TokenClaims, error) {
	// Don't need to verify because our middleware does that
	token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return common.VerifyKey, nil
	})

	if err != nil {
		log.Println("Error getting claims:", err)
		return &models.TokenClaims{}, err
	}

	if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return &models.TokenClaims{}, err
}

func sendError(w http.ResponseWriter, err error, status int, message string) {
	log.Println(err)
	resp, _ := json.Marshal(message)

	w.WriteHeader(status)
	w.Write(resp)
}

func jsonResponse(w http.ResponseWriter, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
