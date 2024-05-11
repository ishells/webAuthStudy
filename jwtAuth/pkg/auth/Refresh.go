package auth

import (
	"log"
	"net/http"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")

}
