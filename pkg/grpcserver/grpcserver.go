package grpcserver

import (
	"context"
	"fibo/proto/api"
	"fibo/pkg"	
	"strings"
	"encoding/json"	
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedFibonacciServer
}

// Fibo ...
func (s *GRPCServer) FiboCount(ctx context.Context, req *api.FiboRequest) (*api.FiboResponse, error) {
	var resp []int32
	resp = pkg.FibSliceRPC(int(req.GetN1()), int(req.GetN2()))
	jsonRes, _ := json.Marshal(resp)
	ret := strings.Trim(string(jsonRes), "[]")
	return &api.FiboResponse{Result: ret}, nil
}
