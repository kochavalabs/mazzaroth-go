package mazzaroth

import (
	"errors"
	"sync/atomic"
)

// ServerSelector is the behaviour to select a new server.
type ServerSelector interface {
	Peek() string
}

// RoundRobinServerSelector implements a round robin server selection.
type RoundRobinServerSelector struct {
	current    uint64
	servers    []string
	numServers uint64
}

// NewRoundRobinServerSelector creates a RoundRobinServerSelector object.
// If the server list is empty the function fails.
// It follows a minimalistic approach and it doesn't validate the structure of the server elements passed in.
func NewRoundRobinServerSelector(servers ...string) (*RoundRobinServerSelector, error) {
	numServers := uint64(len(servers))

	if numServers == 0 {
		return nil, errors.New("could not create the server selector with an empty server list")
	}

	return &RoundRobinServerSelector{
		current:    0,
		servers:    servers,
		numServers: numServers,
	}, nil
}

// Peek returs the next server.
func (rr *RoundRobinServerSelector) Peek() string {
	atomic.AddUint64(&rr.current, 1)

	n := (rr.current - 1) % rr.numServers

	return rr.servers[n]
}
