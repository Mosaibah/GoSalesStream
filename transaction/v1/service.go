package service

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"transaction/store"
	"transaction/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type TransactionService struct{
	proto.TransactionServiceServer
	td store.TransactionDataInterface 
}

func New(td store.TransactionDataInterface) *TransactionService { 
	return &TransactionService{td: td}
}

func (ts *TransactionService) GetTransactions(ctx context.Context, in *proto.TransactionsRequest) (*proto.TransactionsResponse, error){
	var d, err =  ts.td.GetTransactions(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var transactions []*proto.Transaction
	//mapping
	for _, v := range *d {
		var tran = &proto.Transaction{
			Id: v.Id, 
			CustomerId: v.CustomerId,
			ProductId: v.ProductId,
			Price: v.Price,
			Quantity: v.Quantity,
			CreatedAt: timestamppb.New(v.CreatedAt),
		}

		transactions = append(transactions, tran)

		if err != nil {
			log.Print("error while mapping store Transaction to proto message")
			return nil, err
		}
	}

	return &proto.TransactionsResponse{Transactions: transactions}, nil
}

func (ts *TransactionService) GetTransaction(ctx context.Context, in *proto.GetTransactionRequest) (*proto.Transaction, error){
	var res, err = ts.td.GetTransaction(ctx, in.TransactionId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var tran = &proto.Transaction{
		Id: res.Id, 
		CustomerId: res.CustomerId,
		ProductId: res.ProductId,
		Price: res.Price,
		Quantity: res.Quantity,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}

	return tran, nil
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, in *proto.CreateTransactionRequest) (*proto.CreateTransactionResponse, error){
	
	if in.Transaction == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid")
	}
	if in.Transaction.CustomerId <= 0 {
		return nil, fmt.Errorf("customer id must be greater than 0")
	}
	if in.Transaction.ProductId <= 0 {
		return nil, fmt.Errorf("product id must be greater than 0")
	}
	if in.Transaction.Price <= 0 {
		return nil, fmt.Errorf("price must be greater than 0")
	}
	if in.Transaction.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}

	_, err := ts.td.GetCustomerById(ctx, in.Transaction.CustomerId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	_, err = ts.td.GetProductById(ctx, in.Transaction.ProductId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	
	res, err := ts.td.CreateTransaction(
		ctx, 
		store.Transaction{CustomerId: in.Transaction.CustomerId, ProductId: in.Transaction.ProductId, Price: in.Transaction.Price, Quantity: in.Transaction.Quantity})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var tran = &proto.Transaction{
		Id: res.Id, 
		CustomerId: res.CustomerId,
		ProductId: res.ProductId,
		Price: res.Price,
		Quantity: res.Quantity,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}

	return &proto.CreateTransactionResponse{Transaction: tran}, nil
}