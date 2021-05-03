package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type CallBuilder struct {
	transaction *xdr.Transaction
}

func (cb *CallBuilder) Call(address, channel [32]byte, nonce uint64) *CallBuilder {
	if cb.transaction == nil {
		cb.transaction = new(xdr.Transaction)
	}
	call := new(xdr.Call)
	cb.transaction.Action = xdr.Action{
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeCALL,
			Call: call,
		},
	}
	return cb
}

func (cb *CallBuilder) Function(name string) *CallBuilder {
	if cb.transaction.Action.Category.Call == nil {
		cb.transaction.Action.Category.Call = new(xdr.Call)
	}
	cb.transaction.Action.Category.Call.Function = name
	return cb
}

func (cb *CallBuilder) Parameters(f ...Field) *CallBuilder {
	if cb.transaction.Action.Category.Call == nil {
		cb.transaction.Action.Category.Call = new(xdr.Call)
	}
	for _, field := range f {
		cb.transaction.Action.Category.Call.Parameters =
			append(cb.transaction.Action.Category.Call.Parameters, xdr.Parameter(field))
	}
	return cb
}

func (cb *CallBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	if len(cb.transaction.Action.Category.Call.Function) <= 0 {
		return nil, errors.New("Missing required function name")
	}

	actionStream, err := cb.transaction.Action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}

	signatureSlice := ed25519.Sign(pk, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}
	cb.transaction.Signature = signature
	if &cb.transaction.Signer == nil {
		cb.transaction.Signer, err = xdr.NewAuthority(xdr.AuthorityTypeNONE, nil)
		if err != nil {
			return nil, errors.Wrap(err, "in xdr.NewAuthority(xdr.AuthorityTypeNONE)")
		}
	}
	return cb.transaction, nil
}
