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
	var trans []*Transaction
	var created time.Time
	rows, err := db.Query("SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
	for rows.Next(){
		transaction := Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductId,
		&transaction.Price, &transaction.Quantity, &created )
		if err != nil {
			return nil, err
		}
		trans = append(trans, &Transaction{
			Id: transaction.Id,
			CustomerId: transaction.CustomerId, 
			ProductId: transaction.ProductId, 
			Price: transaction.Price, 
			Quantity: transaction.Quantity, 
			CreatedAt: timestamppb.New(created)})
	}

	log.Printf("Log, log")

	return &TransactionsResponse{Transactions: trans}, nil
}

func (s *Server) GetTransaction(ctx context.Context, in *GetTransactionRequest) (*Transaction, error){
	var trans Transaction
	var created time.Time
	query := "SELECT id, customer_id, product_id, price, quantity, created_at FROM transactions where id = $1"
	err := db.QueryRow(query, in.TransactionId).Scan(&trans.Id, &trans.CustomerId, &trans.ProductId,
		&trans.Price, &trans.Quantity, &created )
	switch {
		case err == sql.ErrNoRows:
			log.Printf("no transaction with id %d\n", in.TransactionId)
		case err != nil:
			log.Fatalf("query error: %v\n", err)
		default:
			log.Print("Log, log")
	}
	
	return &Transaction{
			Id: trans.Id, 
			CustomerId: trans.CustomerId, 
			ProductId: trans.ProductId, 
			Price: trans.Price, 
			Quantity: trans.Quantity, 
			CreatedAt: 
			timestamppb.New(created),
			}, nil
}

func (s *Server) CreateTransaction(ctx context.Context, in *CreateTransactionRequest) (*CreateTransactionResponse, error){
	
	insert_query := "INSERT INTO transactions(customer_id, product_id, price, quantity) VALUES( $1, $2, $3, $4 )"
	
	if _, err := db.Exec(insert_query,in.Transaction.CustomerId, in.Transaction.ProductId, in.Transaction.Price, in.Transaction.Quantity); 
	err != nil {
		log.Fatal(err)
	}
	
	log.Println("data => ", in)
	return &CreateTransactionResponse{Transaction: 	&Transaction{
		
			Id: in.Transaction.Id,
			CustomerId: in.Transaction.CustomerId, 
			ProductId: in.Transaction.ProductId, 
			Price: in.Transaction.Price, 
			Quantity: in.Transaction.Quantity, 
			CreatedAt: timestamppb.Now(),
		}}, nil
}
