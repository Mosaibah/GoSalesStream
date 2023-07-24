package store

import (
	"context"
	"time"
)

type MockTransactionData struct{}

func (m *MockTransactionData) GetTransactions(ctx context.Context) ([]Transaction, error) {
	return []Transaction{
		{
			Id: 1,
			CustomerId: 2,
			ProductId: 2,
			Price: 100.0,
			Quantity: 15,
			CreatedAt: time.Now(),
		},{
			Id: 2,
			CustomerId: 2,
			ProductId: 2,
			Price: 100.0,
			Quantity: 15,
			CreatedAt: time.Now(),
		},
	}, nil
}

func (m *MockTransactionData) GetTransaction(ctx context.Context, id int32) (Transaction, error) {
	return Transaction{
		Id: 1,
		CustomerId: 1,
		ProductId: 1,
		Price: 87.3,
		Quantity: 9,
		CreatedAt: time.Now(),
	}, nil
}

func (m *MockTransactionData) CreateTransaction(ctx context.Context, t Transaction) (Transaction, error) {
	t.Id = 1
	t.CreatedAt = time.Now()
	return t, nil
}

func (m *MockTransactionData) GetProductById(ctx context.Context, id int32) (*Product, error) {
	return &Product{
		Id:   1,
		Name: "Product 1",
	}, nil
}

func (m *MockTransactionData) GetCustomerById(ctx context.Context, id int32) (*Customer, error) {
	return &Customer{
		Id:   1,
		Name: "Customer 1",
	}, nil
}