package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"Transaction/proto"
)

func TestGetTransactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	store := New(db)

	rows := sqlmock.NewRows([]string{"id", "customer_id", "product_id", "price", "quantity", "created_at"}).
		AddRow(1, 1, 1, 20.0, 2, time.Now()).
		AddRow(2, 2, 2, 30.0, 3, time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM transactions$").WillReturnRows(rows)

	transactions, err := store.GetTransactions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	transactionData := New(db)

	rows := sqlmock.NewRows([]string{"id", "customer_id", "product_id", "price", "quantity", "created_at"}).
		AddRow(1, 1, 1, 20.0, 2, time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM transactions where id = \\$1$").WithArgs(1).WillReturnRows(rows)

	transaction, err := transactionData.GetTransaction(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), transaction.Id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	newTransaction := &proto.Transaction{
		CustomerId: 1,
		ProductId:  1,
		Price:      20.0,
		Quantity:   2,
		CreatedAt:  timestamppb.Now(),
	}

	transactionData := New(db)

	mock.ExpectQuery("^INSERT INTO transactions(.+)$").
		WithArgs(newTransaction.CustomerId, newTransaction.ProductId, newTransaction.Price, newTransaction.Quantity).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	transaction, err := transactionData.CreateTransaction(context.Background(), newTransaction)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), transaction.Id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}