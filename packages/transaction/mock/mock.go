package mock

import (
	"context"
	"time"
	"GoSalesStream/packages/transaction/store"
)

type MockTransactionData struct{}

func (m *MockTransactionData) GetTransactions(ctx context.Context) (*[]store.Transaction, error) {
	return &[]store.Transaction{
		{
			Id: 1,
			CustomerId: 2,
			ProductId: 2,
			Price: 100,
			Quantity: 15,
			CreatedAt: time.Now(),
		},{
			Id: 2,
			CustomerId: 2,
			ProductId: 2,
			Price: 100,
			Quantity: 15,
			CreatedAt: time.Now(),
		},
	}, nil
}

func (m *MockTransactionData) GetTransaction(ctx context.Context, id int32) (*store.Transaction, error) {
	return &store.Transaction{
		Id: 1,
		CustomerId: 1,
		ProductId: 1,
		Price: 87,
		Quantity: 9,
		CreatedAt: time.Now(),
	}, nil
}

func (m *MockTransactionData) CreateTransaction(ctx context.Context, t store.Transaction) (*store.Transaction, error) {
	t.Id = 1
	t.CreatedAt = time.Now()
	return &t, nil
}

func (m *MockTransactionData) GetProductById(ctx context.Context, id int32) (*store.Product, error) {
	return &store.Product{
		Id:   1,
		Name: "Product 1",
	}, nil
}

func (m *MockTransactionData) GetCustomerById(ctx context.Context, id int32) (*store.Customer, error) {
	return &store.Customer{
		Id:   1,
		Name: "Customer 1",
	}, nil
}

func (m *MockTransactionData) GetTotalSales(ctx context.Context) (*int32, error) {
	var a int32 = 999
	return &a, nil
}

func (m *MockTransactionData) GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error) {
	var a int32 = 998
	return &a, nil
}

func (m *MockTransactionData) GetTop5Customers(ctx context.Context) ([]store.CustomerTotalSpent, error) {
	customers := []store.CustomerTotalSpent{
		{Id: 1, Name: "Abdulrahman", TotalSpent: 5050},
		{Id: 2, Name: "Mohammed", TotalSpent: 3030},
		{Id: 3, Name: "Ibramim", TotalSpent: 75075},
		{Id: 4, Name: "Omar", TotalSpent: 10010},
		{Id: 5, Name: "Abdulaziz", TotalSpent: 2525},
	}
	return customers, nil
}