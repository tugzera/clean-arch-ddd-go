package entities_test

import (
	"aula-01/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateNewOrderItemWithCorrectValues(t *testing.T) {
	item := entities.NewItem("1", "Coca-Cola", 5.32)
	orderItem := entities.NewOrderItem(item, 3)
	require.Equal(t, orderItem.Total, item.Price * 3)
}