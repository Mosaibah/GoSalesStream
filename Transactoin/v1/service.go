package service

import (
	"context"
	"log"
	// "time"

	_ "github.com/lib/pq"
	// "google.golang.org/protobuf/types/known/timestamppb"
	"Transaction/store"
	"Transaction/proto"
)

type TransactionService struct{
	proto.TransactionServiceServer
	td *store.TransactionData
}

func New(td *store.TransactionData) *TransactionService {
	return &TransactionService{td: td}
}

func (ts *TransactionService) GetTransactions(ctx context.Context, in *proto.TransactionsRequest) (*proto.TransactionsResponse, error){
	var d, err =  ts.td.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}
	return &proto.TransactionsResponse{Transactions: d}, nil
}

func (ts *TransactionService) GetTransaction(ctx context.Context, in *proto.GetTransactionRequest) (*proto.Transaction, error){
	var d, err = ts.td.GetTransaction(in.TransactionId)
	if err != nil {
		log.Fatal(err)
	}
	return d, nil
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, in *proto.CreateTransactionRequest) (*proto.CreateTransactionResponse, error){
	var d, err = ts.td.CreateTransaction(in.Transaction )
	if err != nil {
		log.Fatal(err)
	}
	return &proto.CreateTransactionResponse{Transaction: d}, nil
}