package store

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type AnalyticsData struct {
	db *sql.DB
}

type Customer struct {
	Id int32 
	Name string
	TotalSpent float64
}

func New(db *sql.DB) *AnalyticsData {
	return &AnalyticsData{db: db}
}

func (ad AnalyticsData) GetTotalSales(ctx context.Context) (float64, error){
	var totalSales float64
	query := "SELECT SUM(price) FROM transactions"
	err := ad.db.QueryRow(query).Scan(&totalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no sales ")
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Print("Log, log")
	}
	
	return totalSales, nil
}

func (ad AnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*float64, error){
	var totalSales float64
	query := "SELECT SUM(price) FROM transactions where product_id = $1"
	err := ad.db.QueryRow(query, product_id).Scan(&totalSales)
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no sales")
		case err != nil:
			log.Printf("query error: %v\n", err)
			return nil, err
		default:
			log.Print("Log, log")
	}
	
	return &totalSales, nil
}


func (ad AnalyticsData) GetTop5Customers(ctx context.Context) ([]Customer, error){
	var customers[]Customer
	rows, err := ad.db.Query("SELECT customer_id, SUM(price) FROM transactions GROUP BY customer_id ORDER BY SUM(price) desc LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		customer := Customer{}
		err = rows.Scan(&customer.Id, &customer.TotalSpent)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	log.Printf("Log, log")

	return customers, nil
}