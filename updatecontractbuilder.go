package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

type UpdateContractBuilder struct {
	transaction *xdr.Transaction
	contract    *xdr.Contract
}

func (ucb *UpdateContractBuilder) Contract(b []byte) *UpdateContractBuilder {
	return nil
}

func (ucb *UpdateContractBuilder) Version(version string) *UpdateContractBuilder {
	return nil
}

func (ucb *UpdateContractBuilder) Sign() (*xdr.Transaction, error) {
	return nil, nil
}
