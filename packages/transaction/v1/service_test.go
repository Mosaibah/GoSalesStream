package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	transactionpbv1 "GoSalesStream/packages/proto/transaction/v1/genproto"
	"GoSalesStream/packages/transaction/mock"
)

func TestGetTransactions(t *testing.T) {
	s := New(&mock.MockTransactionData{})

	resp, err := s.GetTransactions(context.Background(), &transactionpbv1.TransactionsRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 2, len(resp.Transaction))
	assert.Equal(t, int32(1), resp.Transaction[0].Id)
	assert.Equal(t, int32(2), resp.Transaction[0].ProductId)
	assert.Equal(t, int32(100), resp.Transaction[0].Price)
	assert.Equal(t, int32(15), resp.Transaction[0].Quantity)
}

func TestGetTransaction(t *testing.T) {
	s := New(&mock.MockTransactionData{})

	resp, err := s.GetTransaction(context.Background(), &transactionpbv1.GetTransactionRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, int32(1), resp.ProductId)
	assert.Equal(t, int32(87), resp.Price)
	assert.Equal(t, int32(9), resp.Quantity)
}

func TestCreateTransaction(t *testing.T) {
	s := New(&mock.MockTransactionData{})

	req := &transactionpbv1.CreateTransactionRequest{
		Transaction: &transactionpbv1.Transaction{
			CustomerId: 1,
			ProductId:  1,
			Price: 100,
			Quantity: 1,
		},
	}

	res, err := s.CreateTransaction(context.Background(), req)

	assert.NoError(t, err)

	assert.Equal(t, int32(1), res.Transaction.Id)
	assert.Equal(t, int32(1), res.Transaction.CustomerId)
	assert.Equal(t, int32(1), res.Transaction.ProductId)
	assert.Equal(t, int32(100), res.Transaction.Price)
	assert.Equal(t, int32(1), res.Transaction.Quantity)
}
