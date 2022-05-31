package main

import (
	"fibo/pkg/apiserver"
	"fibo/pkg/grpcserver"
	"fibo/proto/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
	srv := &grpcserver.GRPCServer{}
	api.RegisterFibonacciServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
