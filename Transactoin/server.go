package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	 "Transaction/logic"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := logic.Server{}

	grpcServer := grpc.NewServer()

	logic.RegisterTransactionServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}