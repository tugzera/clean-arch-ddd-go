package entities_test

import (
	"testing"

	"aula-01/entities"

	"github.com/stretchr/testify/require"
)

func TestShouldNotCreateAnOrderWithInvalidCpf(t *testing.T) {
	item := entities.NewItem("1", "Coca-Cola", 5.50, "Alguma description")
	item2 := entities.NewItem("2", "Açaí", 15.90, "Alguma description")
	item3 := entities.NewItem("3", "Heineken", 6.80, "Alguma description")
	orderItem := entities.NewOrderItem(item, 2)
	orderItem2 := entities.NewOrderItem(item2, 2)
	orderItem3 := entities.NewOrderItem(item3, 2)
	_, err := entities.NewOrder("013.213.412-75", []*entities.OrderItem{orderItem, orderItem2, orderItem3})
	require.Error(t, err)
}

func TestShouldCreateAnOrderWithValidCPf(t *testing.T) {
	item := entities.NewItem("1", "Coca-Cola", 5.50, "Alguma description")
	item2 := entities.NewItem("2", "Açaí", 15.90, "Alguma description")
	item3 := entities.NewItem("3", "Heineken", 6.80, "Alguma description")
	orderItem := entities.NewOrderItem(item, 2)
	orderItem2 := entities.NewOrderItem(item2, 1)
	orderItem3 := entities.NewOrderItem(item3, 4)
	order, err := entities.NewOrder("039.187.211-75", []*entities.OrderItem{orderItem, orderItem2, orderItem3})
	require.Nil(t, err)
	require.Equal(t, order.Cpf.Value, "03918721175")
	require.Equal(t, order.Total, (item.Price*2)+(item2.Price*1)+(item3.Price*4))
}
