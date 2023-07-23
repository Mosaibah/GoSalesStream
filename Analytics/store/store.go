package store

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"analytics/proto"
)

type AnalyticsData struct {
	db *sql.DB
}

func New(db *sql.DB) *AnalyticsData {
	return &AnalyticsData{db: db}
}

func (ad AnalyticsData) GetTotalSales(ctx context.Context) (*proto.TotalSales, error){
	var totalSales proto.TotalSales
	query := "SELECT SUM(price) FROM transactions"
	err := ad.db.QueryRow(query).Scan(&totalSales.TotalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no sales ")
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Print("Log, log")
	}
	
	return &totalSales, nil
}

func (ad AnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*proto.SalesByProductResponse, error){
	var res proto.SalesByProductResponse
	res.ProductId = product_id
	query := "SELECT SUM(price) FROM transactions where product_id = $1"
	err := ad.db.QueryRow(query, product_id).Scan(&res.TotalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no sales")
		case err != nil:
			log.Printf("query error: %v\n", err)
			return nil, err
		default:
			log.Print("Log, log")
	}
	
	return &res, nil
}


func (ad AnalyticsData) GetTop5Customers(ctx context.Context) ([]*proto.Customer, error){
	var customers []*proto.Customer
	rows, err := ad.db.Query("SELECT customer_id, SUM(price) FROM transactions GROUP BY customer_id ORDER BY SUM(price) desc LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		customer := proto.Customer{}
		err = rows.Scan(&customer.CustomerId, &customer.TotalSpent)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	log.Printf("Log, log")

	return customers, nil
}