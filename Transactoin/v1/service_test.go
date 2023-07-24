package service

import (
	"context"
	"testing"

	"Transaction/proto"
	"Transaction/store"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	mockService := New(&store.MockTransactionData{})

	resp, err := mockService.GetTransactions(context.Background(), &proto.TransactionsRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 2, len(resp.Transactions))
	assert.Equal(t, int32(1), resp.Transactions[0].Id)
	assert.Equal(t, int32(2), resp.Transactions[0].ProductId)
	assert.Equal(t, float64(100.0), resp.Transactions[0].Price)
	assert.Equal(t, int32(15), resp.Transactions[0].Quantity)
}

func TestGetTransaction(t *testing.T) {
	mockService := New(&store.MockTransactionData{})

	resp, err := mockService.GetTransaction(context.Background(), &proto.GetTransactionRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, int32(1), resp.ProductId)
	assert.Equal(t, float64(87.3), resp.Price)
	assert.Equal(t, int32(9), resp.Quantity)
}

func TestCreateTransaction(t *testing.T) {
	mockService := New(&store.MockTransactionData{})

	req := &proto.CreateTransactionRequest{
		Transaction: &proto.Transaction{
			CustomerId: 1,
			ProductId: 1,
			Price: 100.0,
			Quantity: 1,
		},
	}

	res, err := mockService.CreateTransaction(context.Background(), req)

	assert.NoError(t, err)

	assert.Equal(t, int32(1), res.Transaction.Id)
	assert.Equal(t, int32(1), res.Transaction.CustomerId)
	assert.Equal(t, int32(1), res.Transaction.ProductId)
	assert.Equal(t, 100.0, res.Transaction.Price)
	assert.Equal(t, int32(1), res.Transaction.Quantity)
}