package mazzaroth

import (
	"net/http"
	"time"
)

// mazzarothOptions config options for client
type mazzarothClientOptions struct {
	HttpClient *http.Client
	Servers    []string
}

// Options interface for applying service options
type Options interface {
	apply(*mazzarothClientOptions)
}

// funcMazzarothClientOption wraps a function that modifies Options into an
// implementation of the Option interface.
type funcMazzarothClientOption struct {
	f func(*mazzarothClientOptions)
}

func (fmso *funcMazzarothClientOption) apply(opt *mazzarothClientOptions) {
	fmso.f(opt)
}

func newFuncPacketOption(f func(*mazzarothClientOptions)) *funcMazzarothClientOption {
	return &funcMazzarothClientOption{
		f: f,
	}
}

// WithHttpClient used to set the http client that the mazzaroth client should use
func WithHttpClient(client *http.Client) Options {
	return newFuncPacketOption(func(o *mazzarothClientOptions) {
		o.HttpClient = client
	})
}

// WithHttpClient used to set the http client that the mazzaroth client should use
func WithServers(servers ...string) Options {
	return newFuncPacketOption(func(o *mazzarothClientOptions) {
		o.Servers = servers
	})
}

func defaultOption() *mazzarothClientOptions {
	return &mazzarothClientOptions{
		HttpClient: &http.Client{
			Timeout: 500 * time.Millisecond,
		},
		Servers: []string{"localhost:8080"},
	}
}
