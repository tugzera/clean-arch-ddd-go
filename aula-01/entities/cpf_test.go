package entities_test

import (
	"aula-01/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldValidateCpfLength(t *testing.T) {
	_, err := entities.NewCpf("039.187.211-7512")
	require.Error(t, err)
}

func TestShouldValidateIfCpfHasAllDigitsEqual(t *testing.T) {
	_, err := entities.NewCpf("111.111.111-11")
	require.Error(t, err)
}

func TestShouldValidateIfCpfIsInvalid(t *testing.T) {
	_, err := entities.NewCpf("111.444.777.53")
	require.Error(t, err)
}

func TestShouldValidateIfCpfIsValid(t *testing.T) {
	cpf, err := entities.NewCpf("111.444.777.35")
	require.Nil(t, err)
	require.Equal(t, cpf.Value, "11144477735")
}