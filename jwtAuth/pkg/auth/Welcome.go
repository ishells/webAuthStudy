package auth

import (
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome")

}
