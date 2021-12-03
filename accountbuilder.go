package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

type AccountBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
}

func (acctb *AccountBuilder) Account(sending, channel *xdr.ID, nonce, blockExpirationNumber uint64) *AccountBuilder {
	return nil
}

func (acctb *AccountBuilder) Alias() *AccountBuilder {
	return nil
}

func (acctb *AccountBuilder) Sign() (*xdr.Transaction, error) {
	return nil, nil
}
