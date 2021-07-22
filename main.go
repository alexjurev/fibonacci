package main

import (
	"fibo/internal/app/apiserver"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/fibo", apiserver.FiboCounter)
	log.Fatal(http.ListenAndServe(":8000", router))
}
