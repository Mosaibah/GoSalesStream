package service

import (
	"context"
	"testing"

	analyticspbv1 "GoSalesStream/packages/proto/analytics/v1/genproto"
	 "GoSalesStream/packages/transaction/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalSales(t *testing.T) {
	mockService := New(&mock.MockTransactionData{})

	res, err := mockService.GetTotalSales(context.Background(), &analyticspbv1.TotalSalesRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(999), res.TotalSales)
}

func TestGetSalesByProduct(t *testing.T) {
	mockService := New(&mock.MockTransactionData{})

	res, err := mockService.GetSalesByProduct(context.Background(), &analyticspbv1.SalesByProductRequest{})

	assert.Nil(t, err)

	assert.Equal(t, int32(998), res.TotalSales)
}

func TestGetTop5Customers(t *testing.T) {
	mockService := New(&mock.MockTransactionData{})

	res, err := mockService.GetTop5Customers(context.Background(), &analyticspbv1.Top5CustomersRequest{})

	assert.Nil(t, err)

	assert.Equal(t, 5, len(res.Customer))
	assert.Equal(t, int32(1), res.Customer[0].CustomerId)
	assert.Equal(t, "Mohammed", res.Customer[1].CustomerName)
	assert.Equal(t, int32(75075), res.Customer[2].TotalSpent)
	assert.Equal(t, int32(10010), res.Customer[3].TotalSpent)
	assert.Equal(t, "Abdulaziz", res.Customer[4].CustomerName)
}
