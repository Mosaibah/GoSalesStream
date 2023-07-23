package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"Transaction/store"
	"Transaction/v1"
	"Transaction/proto"
)

const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable&parseTime=true"
// const connStr = "postgres://root:@host.minikube.internal:26260/GoSalesStream?sslmode=disable&parseTime=true"
func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	store := store.New(db)
	service := service.New(store)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterTransactionServiceServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}