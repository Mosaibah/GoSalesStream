package store

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type TransactionData struct {
	db *sql.DB
}

type TransactionDataInterface interface {
	GetTransactions(ctx context.Context) (*[]Transaction, error)
	GetTransaction(ctx context.Context, id int32) (*Transaction, error)
	CreateTransaction(ctx context.Context, t Transaction) (*Transaction, error)
	GetProductById(ctx context.Context, id int32) (*Product, error)
	GetCustomerById(ctx context.Context, id int32) (*Customer, error)
	GetTotalSales(ctx context.Context) (*int32, error)
	GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error)
	GetTop5Customers(ctx context.Context) ([]CustomerTotalSpent, error)
}

func New(db *sql.DB) *TransactionData {
	return &TransactionData{db: db}
}

type Transaction struct {
	Id   int32
	CustomerId int32
	ProductId  int32
	Price int32
	Quantity int32
	CreatedAt time.Time
}

type Product struct {
	Id   int32
	Name string
}

type Customer struct {
	Id   int32
	Name string
}

type CustomerTotalSpent struct {
	Id         int32
	Name       string
	TotalSpent int32
}

func (td *TransactionData) GetTransactions(ctx context.Context) (*[]Transaction, error) {
	var trans []Transaction
	var created time.Time
	rows, err := td.db.QueryContext(ctx, "SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		transaction := Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductId,
			&transaction.Price, &transaction.Quantity, &created)
		if err != nil {
			return nil, err
		}
		trans = append(trans, Transaction{
			Id:         transaction.Id,
			CustomerId: transaction.CustomerId,
			ProductId:  transaction.ProductId,
			Price:      transaction.Price,
			Quantity:   transaction.Quantity,
			CreatedAt:  created})
	}
	return &trans, nil
}

func (td *TransactionData) GetTransaction(ctx context.Context, id int32) (*Transaction, error) {
	var trans Transaction
	var created time.Time
	log.Print("Store:", ctx, "id: ", id)
	query := "SELECT * FROM transactions where id = $1"
	err := td.db.QueryRowContext(ctx, query, id).Scan(&trans.Id, &trans.CustomerId, &trans.ProductId,
		&trans.Price, &trans.Quantity, &created)
	switch {
	case err != nil:
		return nil, err
	default:
		log.Print("Log, log")
	}

	return &Transaction{
		Id:  trans.Id,
		CustomerId: trans.CustomerId,
		ProductId:  trans.ProductId,
		Price:  trans.Price,
		Quantity:  trans.Quantity,
		CreatedAt:  created,
	}, nil
}

func (td *TransactionData) CreateTransaction(ctx context.Context, t Transaction) (*Transaction, error) {
	insert_query := "INSERT INTO transactions(customer_id, product_id, price, quantity) VALUES( $1, $2, $3, $4 ) RETURNING id"

	err := td.db.QueryRowContext(ctx, insert_query, t.CustomerId, t.ProductId, t.Price, t.Quantity).Scan(&t.Id)
	if err != nil {
		return nil, err
	}
	return &Transaction{

		Id: t.Id,
		CustomerId: t.CustomerId,
		ProductId:  t.ProductId,
		Price: t.Price,
		Quantity: t.Quantity,
		CreatedAt: time.Now(),
	}, nil
}

func (td *TransactionData) GetProductById(ctx context.Context, id int32) (*Product, error) {
	var product Product
	query := "SELECT * FROM products where id = $1"
	err := td.db.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name)
	switch {
	case err != nil:
		return nil, err
	default:
		log.Print("Log, log")
	}

	return &product, nil
}

func (td *TransactionData) GetCustomerById(ctx context.Context, id int32) (*Customer, error) {
	var customer Customer
	query := "SELECT * FROM customers where id = $1"
	err := td.db.QueryRowContext(ctx, query, id).Scan(&customer.Id, &customer.Name)
	switch {
	case err != nil:
		return nil, err
	default:
		log.Print("Log, log")
	}

	return &customer, nil
}


func (ad *TransactionData) GetTotalSales(ctx context.Context) (*int32, error) {
	var totalSales int32
	query := "SELECT SUM(price) FROM transactions"
	err := ad.db.QueryRow(query).Scan(&totalSales)
	switch {
	case err != nil:
		return nil, err
	default:
		log.Print("Log, log")
	}

	return &totalSales, nil
}

func (ad *TransactionData) GetSalesByProduct(ctx context.Context, product_id int32) (*int32, error) {
	var totalSales int32
	query := "SELECT SUM(price) FROM transactions where product_id = $1"
	err := ad.db.QueryRow(query, product_id).Scan(&totalSales)
	switch {
	case err != nil:
		return nil, err
	default:
		log.Print("Log, log")
	}

	return &totalSales, nil
}

func (ad *TransactionData) GetTop5Customers(ctx context.Context) ([]CustomerTotalSpent, error) {
	var customers []CustomerTotalSpent
	rows, err := ad.db.Query("SELECT c.Name, t.customer_id, SUM(t.price) FROM transactions t inner join customers c on c.Id = t.customer_id GROUP BY c.Name, t.customer_id ORDER BY SUM(t.price) desc LIMIT 5")
	if err != nil {
		return nil, err
	}
	log.Println(rows)
	for rows.Next() {
		customer := CustomerTotalSpent{}
		err = rows.Scan(&customer.Name ,&customer.Id, &customer.TotalSpent)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	log.Printf("Log, log")

	return customers, nil
}