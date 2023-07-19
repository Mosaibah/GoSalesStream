package logic

import (
	"context"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

type Server struct {
	TransactionServiceServer
}

var transactions = []*TransactionResponse{
	{
		Id:    1,
		CustomerId:  1,
		ProductId:  55,
		Price: 93,
		Quantity: 12,
		// CreatedAt: time.Now(),
	},
	{
		Id:    3,
		CustomerId:  2,
		ProductId:  55,
		Price: 93,
		Quantity: 12,
		// CreatedAt: now.Unix(),
	},
}

const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable"
var db, err = sql.Open("postgres", connStr)
func (s *Server) GetTransactions(ctx context.Context, in *TransactionsRequest) (*TransactionsResponse, error){
	if err != nil {
		log.Fatal(err)
	}
	var trans []*TransactionResponse
	rows, err := db.Query("SELECT id, customer_id, product_id, price, quantity FROM transactions ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		transaction := TransactionResponse{}
		err = rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductId,
		&transaction.Price, &transaction.Quantity)
		if err != nil {
			return nil, err
		}
		trans = append(trans, &transaction)

	}


	log.Printf("this is a log log log ")

	return &TransactionsResponse{Transactions: trans}, nil
}

