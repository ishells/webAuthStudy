package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome")
	jwt.ParseWithClaims()
}
