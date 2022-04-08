package entities_test

import (
	"testing"

	"aula-01/entities"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateANewItem(t *testing.T) {
	item := entities.NewItem("1", "Coca-Cola", 3.60)
	require.Equal(t, item.ID, "1")
	require.Equal(t, item.Name, "Coca-Cola")
	require.Equal(t, item.Price, 3.60)
}
