package service

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"analytics/store"
	"analytics/proto"
)

type AnalyticsService struct{
	proto.AnalyticsServiceServer
	ad store.AnalyticsDataInterface
}

func New(ad store.AnalyticsDataInterface) *AnalyticsService {
	return &AnalyticsService{ad: ad}
}

func (as *AnalyticsService) GetTotalSales(ctx context.Context, in *proto.TotalSalesRequest) (*proto.TotalSales, error){
	var total_sales, err =  as.ad.GetTotalSales(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	return &proto.TotalSales{TotalSales: *total_sales}, nil
}

func (as *AnalyticsService) GetSalesByProduct(ctx context.Context, in *proto.SalesByProductRequest) (*proto.SalesByProductResponse, error){
	var total_sales, err = as.ad.GetSalesByProduct(ctx, in.ProductId)
	if err != nil {
		log.Fatal(err)
	}
	return &proto.SalesByProductResponse{ProductId: in.ProductId, TotalSales: *total_sales}, err
}

func (as *AnalyticsService) GetTop5Customers(ctx context.Context, in *proto.Top5CustomersRequest) (*proto.Top5CustomersResponse, error){
	var res, err = as.ad.GetTop5Customers(ctx) // []*Customer
	if err != nil {
		log.Fatal(err)
	}

	var customers []*proto.Customer
	//mapping
	for _, v := range res {
		var customer = &proto.Customer{
			CustomerId: v.Id,
			CustomerName: v.Name,
			TotalSpent: v.TotalSpent,
		}
		customers = append(customers, customer)
	}

	return &proto.Top5CustomersResponse{Customer: customers}, nil
}