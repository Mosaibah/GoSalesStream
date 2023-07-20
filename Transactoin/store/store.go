// data/transactionData.go
package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"Transaction/proto"

)

const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable&parseTime=true"
var db, err = sql.Open("postgres", connStr)

type TransactionData struct{
	db *sql.DB
}

func New(db *sql.DB) *TransactionData {
	return &TransactionData{db: db}
}

func (td *TransactionData) GetTransactions() ([]*proto.Transaction, error){
	var trans []*proto.Transaction
	var created time.Time
	rows, err := db.Query("SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		transaction := proto.Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductId,
		&transaction.Price, &transaction.Quantity, &created )
		if err != nil {
			return nil, err
		}
		trans = append(trans, &proto.Transaction{
			Id: transaction.Id,
			CustomerId: transaction.CustomerId, 
			ProductId: transaction.ProductId, 
			Price: transaction.Price, 
			Quantity: transaction.Quantity, 
			CreatedAt: timestamppb.New(created)})
	}
	return trans, nil
}

func (td *TransactionData) GetTransaction(id int64) (*proto.Transaction, error){
	var trans proto.Transaction
	var created time.Time
	query := "SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions where id = $1"
	err := db.QueryRow(query, id).Scan(&trans.Id, &trans.CustomerId, &trans.ProductId,
		&trans.Price, &trans.Quantity, &created )
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no transaction with id %d\n", id)
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Print("Log, log")
	}
	
	return &proto.Transaction{
			Id: trans.Id, 
			CustomerId: trans.CustomerId, 
			ProductId: trans.ProductId, 
			Price: trans.Price, 
			Quantity: trans.Quantity, 
			CreatedAt: 
			timestamppb.New(created),
			}, nil
}

func (td *TransactionData) CreateTransaction(t *proto.Transaction) (*proto.Transaction, error){
	insert_query := "INSERT INTO transactions(customer_id, product_id, price, quantity) VALUES( $1, $2, $3, $4 )"
	
	if _, err := db.Exec(insert_query,t.CustomerId, t.ProductId, t.Price, t.Quantity); 
	err != nil {
		log.Fatal(err)
	}
	
	log.Println("data => ", t)
	return &proto.Transaction{
		
			Id: t.Id,
			CustomerId: t.CustomerId, 
			ProductId: t.ProductId, 
			Price: t.Price, 
			Quantity: t.Quantity, 
			CreatedAt: timestamppb.Now(),
		}, nil
}