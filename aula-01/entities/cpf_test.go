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