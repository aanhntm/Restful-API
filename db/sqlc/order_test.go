package db

import (
	"context"
	"testing"

	"github.com/aanhntm/restful-api/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		UserName:    util.RandomName(),
		ProductName: util.RandomProductName(),
		Amount:      util.RandomAmount(),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.NotZero(t, order.ID)
	require.NotZero(t, order.Amount)

	return order
}

func TestCreateOrder(t *testing.T) {
	CreateRandomOrder(t)
}

func GetOrder(t *testing.T) Order {
	order1 := CreateRandomOrder(t)
	order2, err := testQueries.GetOneOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserName, order2.UserName)

	return order2
}

func TestGetManyOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomOrder(t)
	}

	arg := GetManyOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.GetManyOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}
