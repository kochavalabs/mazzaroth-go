package mazzaroth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRoundRobinServerSelectorHappyPath checks the correct behaviour of the component.
func TestRoundRobinServerSelectorHappyPath(t *testing.T) {
	rr, err := NewRoundRobinServerSelector("a", "b", "c", "d")
	require.NoError(t, err)

	require.Equal(t, "a", rr.Pick())
	require.Equal(t, "b", rr.Pick())
	require.Equal(t, "c", rr.Pick())
	require.Equal(t, "d", rr.Pick())
	require.Equal(t, "a", rr.Pick())
}

// TestRoundRobinServerDetectsEmptyServerList checks the constructor function returns an error when an empty list is passed in.
func TestRoundRobinServerDetectsEmptyServerList(t *testing.T) {
	rr, err := NewRoundRobinServerSelector()
	require.Error(t, err)
	require.Equal(t, "could not create the server selector with an empty server list", err.Error())
	require.Nil(t, rr)
}
