package main

import (
	"log"
	"net/http"
	"webAuthStudy/pkg/auth"
)

func main() {
	http.HandleFunc("/sign", auth.Signin)
	http.HandleFunc("/welcome", auth.Welcome)
	http.HandleFunc("/refresh", auth.Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
