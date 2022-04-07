package entities_test

import (
	"aula-01/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldNotCreateAnOrderWithInvalidCpf(t *testing.T) {
	_, err := entities.NewOrder("013.213.412-75")
	require.Error(t, err)
}

func TestShouldCreateAnOrderWithValidCPf(t *testing.T) {
	order, err := entities.NewOrder("039.187.211-75")
	require.Nil(t, err) 
	require.Equal(t, order.Cpf.Value, "03918721175")
}
