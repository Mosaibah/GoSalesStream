package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"GoSalesStream/packages/transaction/store"
	"GoSalesStream/packages/transaction/v1"
	transactionpbv1 "GoSalesStream/packages/proto/transaction/v1/genproto"
	// "google.golang.org/grpc/reflection"
	// "fmt"
	// "os"
)

func main() {

	const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable&parseTime=true"


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
	
	transactionpbv1.RegisterTransactionServiceServer(grpcServer, service)
	// reflection.Register(grpcServer)

	defer db.Close()
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}