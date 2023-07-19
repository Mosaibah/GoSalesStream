package logic

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	TransactionServiceServer
}


const connStr = "postgres://root:@localhost:26260/GoSalesStream?sslmode=disable&parseTime=true"
var db, err = sql.Open("postgres", connStr)
func (s *Server) GetTransactions(ctx context.Context, in *TransactionsRequest) (*TransactionsResponse, error){
	if err != nil {
		log.Fatal(err)
	}
	var trans []*TransactionResponse
	var created time.Time
	rows, err := db.Query("SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		transaction := TransactionResponse{}
		err = rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductId,
		&transaction.Price, &transaction.Quantity, &created )
		if err != nil {
			return nil, err
		}
		trans = append(trans, &TransactionResponse{
			Id: transaction.Id,
			CustomerId: transaction.CustomerId, 
			ProductId: transaction.ProductId, 
			Price: transaction.Price, 
			Quantity: transaction.Quantity, 
			CreatedAt: timestamppb.New(created)})
	}


	log.Printf("this is a log log log ")

	return &TransactionsResponse{Transactions: trans}, nil
}

