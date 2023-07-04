package initializers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectDatabase(t *testing.T) {
	Env()
	ConnectDatabase()

	require.NotEmpty(t, DB)
}
