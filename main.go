package main

import (
	"fibo/internal/app/apiserver"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"fibo/internal/app/proto/api"
	"fibo/pkg"
	"fibo/pkg/fibocounter"
	"google.golang.org/grpc"
	"net"
	"strings"
	"encoding/json"
)

func main() {
	// REST
	router := mux.NewRouter()
	router.HandleFunc("/fibo", apiserver.FiboCounter)

	go func() {
        http.ListenAndServe("localhost:8000", router)
    }()
	
	//GRPC
	
	ret2 := pkg.FibSliceRPC(1, 4)
	ret3, _ := json.Marshal(ret2)
	ret4 := strings.Trim(string(ret3), "[]")
	fmt.Println(ret4)
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
