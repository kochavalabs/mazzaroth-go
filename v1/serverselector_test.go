package mazzaroth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRoundRobinServerSelectorHapyPath checks the correct behaviour of the component.
func TestRoundRobinServerSelectorHapyPath(t *testing.T) {
	rr, err := NewRoundRobinServerSelector("a", "b", "c", "d")
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		require.Equal(t, "a", rr.Peek())
		require.Equal(t, "b", rr.Peek())
		require.Equal(t, "c", rr.Peek())
		require.Equal(t, "d", rr.Peek())
	}
}

// TestRoundRobinServerDetectsEmptyServerList checks the constructor function returns an error when an empty list is passed in.
func TestRoundRobinServerDetectsEmptyServerList(t *testing.T) {
	rr, err := NewRoundRobinServerSelector()
	require.Error(t, err)
	require.Equal(t, "could not create the server selector with an empty server list", err.Error())
	require.Nil(t, rr)
}
