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
	Id         int32
	Name       string
	TotalSpent int32
}

type AnalyticsDataInterface interface {
	GetTotalSales(ctx context.Context) (*int32, error)
	GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error)
	GetTop5Customers(ctx context.Context) ([]Customer, error)
}

func New(db *sql.DB) *AnalyticsData {
	return &AnalyticsData{db: db}
}

func (ad *AnalyticsData) GetTotalSales(ctx context.Context) (*int32, error) {
	var totalSales int32
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

	return &totalSales, nil
}

func (ad *AnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error) {
	var totalSales int32
	query := "SELECT SUM(price) FROM transactions where product_id = $1"
	err := ad.db.QueryRow(query, product_id).Scan(&totalSales)
	switch {
	case err == sql.ErrNoRows:
		log.Fatal("no sales")
	case err != nil:
		log.Fatal("query error: %v\n", err)
	default:
		log.Print("Log, log")
	}

	return &totalSales, nil
}

func (ad *AnalyticsData) GetTop5Customers(ctx context.Context) ([]Customer, error) {
	var customers []Customer
	rows, err := ad.db.Query("SELECT c.Name, t.customer_id, SUM(t.price) FROM transactions t inner join customers c on c.Id = t.customer_id GROUP BY c.Name, t.customer_id ORDER BY SUM(t.price) desc LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.Name ,&customer.Id, &customer.TotalSpent)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	log.Printf("Log, log")

	return customers, nil
}
