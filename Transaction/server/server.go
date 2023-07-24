package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"transaction/store"
	"transaction/v1"
	"transaction/proto"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(grpcServer)

	proto.RegisterTransactionServiceServer(grpcServer, service)

	defer db.Close()
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}