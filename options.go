package mazzaroth

import (
	"net/http"
	"time"
)

// mazzarothOptions config options for client
type mazzarothClientOptions struct {
	httpClient *http.Client
	address    string
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

// WithAddress used to set the http client that the mazzaroth client should use
func WithAddress(address string) Options {
	return newFuncPacketOption(func(o *mazzarothClientOptions) {
		o.address = address
	})
}

// defaultOption defines a set of default options for the mazzaroth client
func defaultOption() *mazzarothClientOptions {
	return &mazzarothClientOptions{
		httpClient: &http.Client{
			Timeout: 500 * time.Millisecond,
		},
		address: "http://localhost:6299",
	}
}
