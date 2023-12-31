package service

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"GoSalesStream/packages/transaction/store"
	transactionpbv1 "GoSalesStream/packages/proto/transaction/v1/genproto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type TransactionService struct{
	transactionpbv1.TransactionServiceServer
	td store.TransactionDataInterface 
}

func New(td store.TransactionDataInterface) *TransactionService { 
	return &TransactionService{td: td}
}

func (ts *TransactionService) GetTransactions(ctx context.Context, in *transactionpbv1.TransactionsRequest) (*transactionpbv1.TransactionsResponse, error){
	var d, err =  ts.td.GetTransactions(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var transactions []*transactionpbv1.Transaction
	//mapping
	for _, v := range *d {
		var tran = &transactionpbv1.Transaction{
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

	return &transactionpbv1.TransactionsResponse{Transaction: transactions}, nil
}

func (ts *TransactionService) GetTransaction(ctx context.Context, in *transactionpbv1.GetTransactionRequest) (*transactionpbv1.Transaction, error){
	var res, err = ts.td.GetTransaction(ctx, in.TransactionId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var tran = &transactionpbv1.Transaction{
		Id: res.Id, 
		CustomerId: res.CustomerId,
		ProductId: res.ProductId,
		Price: res.Price,
		Quantity: res.Quantity,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}

	return tran, nil
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, in *transactionpbv1.CreateTransactionRequest) (*transactionpbv1.CreateTransactionResponse, error){
	log.Print("service: CreateTransaction: ", in)
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

	var tran = &transactionpbv1.Transaction{
		Id: res.Id, 
		CustomerId: res.CustomerId,
		ProductId: res.ProductId,
		Price: res.Price,
		Quantity: res.Quantity,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}

	return &transactionpbv1.CreateTransactionResponse{Transaction: tran}, nil
}

func (as *TransactionService) GetTotalSales(ctx context.Context, in *transactionpbv1.TotalSalesRequest) (*transactionpbv1.TotalSales, error){
	var total_sales, err =  as.td.GetTotalSales(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &transactionpbv1.TotalSales{TotalSales: *total_sales}, nil
}

func (as *TransactionService) GetSalesByProduct(ctx context.Context, in *transactionpbv1.SalesByProductRequest) (*transactionpbv1.SalesByProductResponse, error){
	var total_sales, err = as.td.GetSalesByProduct(ctx, in.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &transactionpbv1.SalesByProductResponse{ProductId: in.ProductId, TotalSales: *total_sales}, err
}

func (as *TransactionService) GetTop5Customers(ctx context.Context, in *transactionpbv1.Top5CustomersRequest) (*transactionpbv1.Top5CustomersResponse, error){
	var res, err = as.td.GetTop5Customers(ctx) // []*CustomerTotalSpent
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var customers []*transactionpbv1.Customer
	//mapping
	for _, v := range res {
		var customer = &transactionpbv1.Customer{
			CustomerId: v.Id,
			CustomerName: v.Name,
			TotalSpent: v.TotalSpent,
		}
		customers = append(customers, customer)
	}

	return &transactionpbv1.Top5CustomersResponse{Customer: customers}, nil
}