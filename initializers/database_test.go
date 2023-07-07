package initializers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectDatabase(t *testing.T) {
	Env()
	ConnectDatabase()
	require.NoError(t, Err)
	require.NotEmpty(t, DB)
}

// todo test error db
