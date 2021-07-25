package main

import (
	"fibo/internal/app/apiserver"
	"github.com/gorilla/mux"
	"net/http"
	
	"log"
	"fibo/internal/app/proto/api"
	
	"fibo/pkg/fibocounter"
	"google.golang.org/grpc"
	"net"

	
)

func main() {
	// REST
	router := mux.NewRouter()
	router.HandleFunc("/fibo", apiserver.FiboCounter)

	go func() {
        http.ListenAndServe("localhost:8000", router)
    }()
	
	//GRPC
	s := grpc.NewServer()
	srv := &fibocounter.GRPCServer{}
	api.RegisterFibonacciServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
