package client

import (
	"net/http"
	"time"
)

// mazzarothOptions config options for client
type mazzarothClientOptions struct {
	httpClient *http.Client
	servers    []string
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
		o.httpClient = client
	})
}

// WithHttpClient used to set the http client that the mazzaroth client should use
func WithServers(servers ...string) Options {
	return newFuncPacketOption(func(o *mazzarothClientOptions) {
		o.servers = servers
	})
}

func defaultPacketOption() *mazzarothClientOptions {
	return &mazzarothClientOptions{
		httpClient: &http.Client{
			Timeout: 500 * time.Millisecond,
		},
	}
}
