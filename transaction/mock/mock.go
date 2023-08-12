package mock

import (
	"context"
	"time"
	"GoSalesStream/transaction/store"
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