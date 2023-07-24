package store

import (
	"context"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions_success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	store := New(db) // Alocate memory but not initialize, (return a pointer)

	rows := sqlmock.NewRows([]string{"id", "customer_id", "product_id", "price", "quantity", "created_at"}).
		AddRow(1, 1, 1, 20.0, 2, time.Now()).
		AddRow(2, 2, 2, 30.0, 3, time.Now())

	query := "SELECT (.+) FROM transactions"
	mock.ExpectQuery(query).WillReturnRows(rows)

	transactions, err := store.GetTransactions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetTransaction_success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()	

	store := New(db)

	rows := sqlmock.NewRows([]string{"id", "customer_id", "product_id", "price", "quantity", "created_at"}).
		AddRow(1, 2, 3, 20.5, 15, time.Now())
		
	query := "SELECT (.+) FROM transactions where id = \\$1"
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	transaction, err := store.GetTransaction(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), transaction.Id)
	assert.Equal(t, int32(2), transaction.CustomerId)
	assert.Equal(t, int32(3), transaction.ProductId)
	assert.Equal(t, 20.5, transaction.Price)
	assert.Equal(t, int32(15), transaction.Quantity)

	assert.NoError(t, err)
}

func TestCreateTransaction_success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	newTransaction := Transaction{
		CustomerId: 2,
		ProductId: 19,
		Price: 1892.7,
		Quantity: 73,
		CreatedAt: time.Now(),
	}

	store := New(db)

	query := "INSERT INTO transactions(.+)"
	mock.ExpectQuery(query).
		WithArgs(newTransaction.CustomerId, newTransaction.ProductId, newTransaction.Price, newTransaction.Quantity).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	transaction, err := store.CreateTransaction(context.Background(), newTransaction)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), transaction.Id)
	assert.Equal(t, int32(2), transaction.CustomerId)
	assert.Equal(t, int32(19), transaction.ProductId)
	assert.Equal(t, 1892.7, transaction.Price)
	assert.Equal(t, int32(73), transaction.Quantity)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}