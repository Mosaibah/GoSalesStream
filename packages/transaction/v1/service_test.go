package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	transactionpbv1 "GoSalesStream/packages/proto/transaction/v1/genproto"
	"GoSalesStream/packages/transaction/mock"
)

func TestGetTransactions(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	resp, err := mockData.GetTransactions(context.Background(), &transactionpbv1.TransactionsRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 2, len(resp.Transaction))
	assert.Equal(t, int32(1), resp.Transaction[0].Id)
	assert.Equal(t, int32(2), resp.Transaction[0].ProductId)
	assert.Equal(t, int32(100), resp.Transaction[0].Price)
	assert.Equal(t, int32(15), resp.Transaction[0].Quantity)
}

func TestGetTransaction(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	resp, err := mockData.GetTransaction(context.Background(), &transactionpbv1.GetTransactionRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, int32(1), resp.ProductId)
	assert.Equal(t, int32(87), resp.Price)
	assert.Equal(t, int32(9), resp.Quantity)
}

func TestCreateTransaction(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	req := &transactionpbv1.CreateTransactionRequest{
		Transaction: &transactionpbv1.Transaction{
			CustomerId: 1,
			ProductId:  1,
			Price: 100,
			Quantity: 1,
		},
	}

	res, err := mockData.CreateTransaction(context.Background(), req)

	assert.NoError(t, err)

	assert.Equal(t, int32(1), res.Transaction.Id)
	assert.Equal(t, int32(1), res.Transaction.CustomerId)
	assert.Equal(t, int32(1), res.Transaction.ProductId)
	assert.Equal(t, int32(100), res.Transaction.Price)
	assert.Equal(t, int32(1), res.Transaction.Quantity)
}


func TestGetTotalSales(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	res, err := mockData.GetTotalSales(context.Background(), &transactionpbv1.TotalSalesRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(999), res.TotalSales)
}

func TestGetSalesByProduct(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	res, err := mockData.GetSalesByProduct(context.Background(), &transactionpbv1.SalesByProductRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(998), res.TotalSales)
}

func TestGetTop5Customers(t *testing.T) {
	mockData := New(&mock.MockTransactionData{})

	res, err := mockData.GetTop5Customers(context.Background(), &transactionpbv1.Top5CustomersRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 5, len(res.Customer))
	assert.Equal(t, int32(1), res.Customer[0].CustomerId)
	assert.Equal(t, "Mohammed", res.Customer[1].CustomerName)
	assert.Equal(t, int32(75075), res.Customer[2].TotalSpent)
	assert.Equal(t, int32(10010), res.Customer[3].TotalSpent)
	assert.Equal(t, "Abdulaziz", res.Customer[4].CustomerName)
}
