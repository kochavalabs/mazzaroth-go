package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

type UpdateConfigBuilder struct {
	transaction *xdr.Transaction
	config      *xdr.ChannelConfig
}

func (ucb *UpdateConfigBuilder) Owner(ID [32]byte) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) ChannelName(name string) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) Admins(IDs [][32]byte) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) Sign() (*xdr.Transaction, error) {
	return nil, nil
}
