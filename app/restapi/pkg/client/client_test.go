package client

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Client(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		test := true
		require.True(t, test)
	})
}
