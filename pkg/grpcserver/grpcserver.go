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
	var ret2 []int32
	ret2 = pkg.FibSliceRPC(int(req.GetN1()), int(req.GetN2()))
	ret3, _ := json.Marshal(ret2)
	ret4 := strings.Trim(string(ret3), "[]")
	return &api.FiboResponse{Result: ret4}, nil
}
