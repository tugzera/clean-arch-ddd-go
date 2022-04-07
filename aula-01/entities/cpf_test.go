package entities_test

import (
	"aula-01/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldValidateCpf(t *testing.T) {
	cpf :=entities.NewCpf("123")
	require.Equal(t, cpf.Value, "123")
}
