package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	value := RandomInt(1, 343)

	require.NotZero(t, value)
	require.Positive(t, value)
}

func TestRandomString(t *testing.T) {
	value := RandomString(10)

	require.NotEmpty(t, value)
	require.Len(t, value, 10)
}

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInt(int64(i), int64(i))
	}
}

func TestRandomOwner(t *testing.T) {
	value := RandomOwner()
	require.NotEmpty(t, value)
	require.Len(t, value, 6)
}

func TestRandomMoney(t *testing.T) {
	value := RandomMoney()

	require.NotZero(t, value)
	require.Positive(t, value)
}

func TestRandomCurrency(t *testing.T) {
	value := RandomCurrency()

	require.NotEmpty(t, value)
	require.Len(t, value, 3)
}
