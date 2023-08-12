package mock

import (
	"context"
	"GoSalesStream/packages/analytics/store"
)

type MockAnalyticsData struct{}

func (m *MockAnalyticsData) GetTotalSales(ctx context.Context) (*int32, error) {
	var a int32 = 999
	return &a, nil
}

func (m *MockAnalyticsData) GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error) {
	var a int32 = 998
	return &a, nil
}

func (m *MockAnalyticsData) GetTop5Customers(ctx context.Context) ([]store.Customer, error) {
	customers := []store.Customer{
		{Id: 1, Name: "Abdulrahman", TotalSpent: 5050},
		{Id: 2, Name: "Mohammed", TotalSpent: 3030},
		{Id: 3, Name: "Ibramim", TotalSpent: 75075},
		{Id: 4, Name: "Omar", TotalSpent: 10010},
		{Id: 5, Name: "Abdulaziz", TotalSpent: 2525},
	}
	return customers, nil
}
