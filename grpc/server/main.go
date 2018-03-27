package main

import (
	"net"
	"google.golang.org/grpc"
	pb "github.com/evilabandon/gonetstudy/grpc/protobuf"
	"github.com/evilabandon/gonetstudy/grpc/response"
	"log"
)

const (
	PORT = ":10023"
)

func main()  {
	lis, err:= net.Listen("tcp",PORT)
	if err!=nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataServer(s, &response.Server{})
	s.Serve(lis)
}

