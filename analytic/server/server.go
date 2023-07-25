package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"analytics/store"
	"analytics/v1"
	"analytics/proto"
	"database/sql"
	"google.golang.org/grpc/reflection"
	"os"
	"fmt"
)

func main() {

	username := os.Getenv("username")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")
  
	connStr := fmt.Sprintf("postgres://%s:@%s:%s/%s?sslmode=%s&parseTime=true", username, host, port, dbname, sslmode)
  

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	store := store.New(db)
	service := service.New(store)

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	proto.RegisterAnalyticsServiceServer(grpcServer, service)

	defer db.Close()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}