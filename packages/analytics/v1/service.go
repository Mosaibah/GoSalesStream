package service

import (
	"context"
	_ "github.com/lib/pq"
	"GoSalesStream/packages/analytics/store"
	analyticspbv1 "GoSalesStream/packages/proto/analytics/v1/genproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AnalyticsService struct{
	analyticspbv1.AnalyticsServiceServer
	ad store.AnalyticsDataInterface
}

func New(ad store.AnalyticsDataInterface) *AnalyticsService {
	return &AnalyticsService{ad: ad}
}

func (as *AnalyticsService) GetTotalSales(ctx context.Context, in *analyticspbv1.TotalSalesRequest) (*analyticspbv1.TotalSales, error){
	var total_sales, err =  as.ad.GetTotalSales(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &analyticspbv1.TotalSales{TotalSales: *total_sales}, nil
}

func (as *AnalyticsService) GetSalesByProduct(ctx context.Context, in *analyticspbv1.SalesByProductRequest) (*analyticspbv1.SalesByProductResponse, error){
	var total_sales, err = as.ad.GetSalesByProduct(ctx, in.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &analyticspbv1.SalesByProductResponse{ProductId: in.ProductId, TotalSales: *total_sales}, err
}

func (as *AnalyticsService) GetTop5Customers(ctx context.Context, in *analyticspbv1.Top5CustomersRequest) (*analyticspbv1.Top5CustomersResponse, error){
	var res, err = as.ad.GetTop5Customers(ctx) // []*Customer
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var customers []*analyticspbv1.Customer
	//mapping
	for _, v := range res {
		var customer = &analyticspbv1.Customer{
			CustomerId: v.Id,
			CustomerName: v.Name,
			TotalSpent: v.TotalSpent,
		}
		customers = append(customers, customer)
	}

	return &analyticspbv1.Top5CustomersResponse{Customer: customers}, nil
}