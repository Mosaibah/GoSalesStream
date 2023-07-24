package store

import (
	"context"
)

type MockAnalyticsData struct{}


func (m *MockAnalyticsData) GetTotalSales(ctx context.Context) (*float64, error){
	var a float64 = 999
	return &a, nil
}

func (m *MockAnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*float64, error){
	var a float64 = 998
	return &a, nil
}


func (m *MockAnalyticsData) GetTop5Customers(ctx context.Context) ([]Customer, error){
	customers := []Customer{
		{Id: 1, Name: "Abdulrahman", TotalSpent: 500.50},
		{Id: 2, Name: "Mohammed", TotalSpent: 300.30},
		{Id: 3, Name: "Ibramim", TotalSpent: 750.75},
		{Id: 4, Name: "Omar", TotalSpent: 1000.10},
		{Id: 5, Name: "Abdulaziz", TotalSpent: 250.25},
	}
	return customers, nil
}