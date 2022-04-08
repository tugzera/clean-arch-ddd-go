package entities_test

import (
	"testing"

	"aula-01/entities"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateNewOrderItemWithCorrectValues(t *testing.T) {
	item := entities.NewItem("1", "Coca-Cola", 5.32, "Alguma descrição")
	orderItem := entities.NewOrderItem(item, 3)
	require.Equal(t, orderItem.Total, item.Price*3)
}
