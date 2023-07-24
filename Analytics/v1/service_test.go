package service

import (
	"context"
	"testing"

	"analytics/proto"
	"analytics/store"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalSales(t *testing.T) {
	mockService := New(&store.MockAnalyticsData{})

	res, err := mockService.GetTotalSales(context.Background(), &proto.TotalSalesRequest{})

	assert.Nil(t, err)

	assert.Equal(t, float64(999), res.TotalSales)
}

func TestGetSalesByProduct(t *testing.T) {
	mockService := New(&store.MockAnalyticsData{})

	res, err := mockService.GetSalesByProduct(context.Background(), &proto.SalesByProductRequest{})

	assert.Nil(t, err)

	assert.Equal(t, float64(998), res.TotalSales)
}

func TestGetTop5Customers(t *testing.T) {
	mockService := New(&store.MockAnalyticsData{})

	res, err := mockService.GetTop5Customers(context.Background(), &proto.Top5CustomersRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 5, len(res.Customer))
	assert.Equal(t, int32(1), res.Customer[0].CustomerId)
	assert.Equal(t, "Mohammed", res.Customer[1].CustomerName)
	assert.Equal(t, 750.75, res.Customer[2].TotalSpent)
	assert.Equal(t, 1000.10, res.Customer[3].TotalSpent)
	assert.Equal(t, "Abdulaziz", res.Customer[4].CustomerName)
}