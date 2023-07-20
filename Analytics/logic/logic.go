package logic

import (
	"context"
	"database/sql"
	"log"
	// "time"
	_ "github.com/lib/pq"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	AnalyticsServiceServer
}

const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable&parseTime=true"
var db, err = sql.Open("postgres", connStr)

func (s *Server) GetTotalSales(ctx context.Context, in *TotalSalesRequest) (*TotalSales, error){
	var totalSales TotalSales
	query := "SELECT SUM(price) FROM transactions"
	err := db.QueryRow(query).Scan(&totalSales.TotalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no transaction  %d\n")
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Print("Log, log")
	}
	
	return &totalSales, nil
}

func (s *Server) GetSalesByProduct(ctx context.Context, in *SalesByProductRequest) (*SalesByProductResponse, error){
	var res SalesByProductResponse
	res.ProductId = in.ProductId
	query := "SELECT SUM(price) FROM transactions where product_id = $1"
	err := db.QueryRow(query, in.ProductId).Scan(&res.TotalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no transaction  %d\n")
		case err != nil:
			log.Printf("query error: %v\n", err)
			return nil, err
		default:
			log.Print("Log, log")
	}
	
	return &res, nil
}