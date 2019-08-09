package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//go mod init github.com/Ulbora/fflLoad
func main() {
	router := mux.NewRouter()
	port := "3000"

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)
}
