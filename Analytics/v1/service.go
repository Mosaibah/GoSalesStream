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
	ad *store.AnalyticsData
}

func New(ad *store.AnalyticsData) *AnalyticsService {
	return &AnalyticsService{ad: ad}
}

func (as *AnalyticsService) GetTotalSales(ctx context.Context, in *proto.TotalSalesRequest) (*proto.TotalSales, error){
	var res, err =  as.ad.GetTotalSales(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (as *AnalyticsService) GetSalesByProduct(ctx context.Context, in *proto.SalesByProductRequest) (*proto.SalesByProductResponse, error){
	var res, err = as.ad.GetSalesByProduct(ctx, in.ProductId)
	if err != nil {
		log.Fatal(err)
	}
	return res, err
}

func (as *AnalyticsService) GetTop5Customers(ctx context.Context, in *proto.Top5CustomersRequest) (*proto.Top5CustomersResponse, error){
	var res, err = as.ad.GetTop5Customers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &proto.Top5CustomersResponse{Customer: res}, nil
}