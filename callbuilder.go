package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type CallBuilder struct {
	address, channel xdr.ID
	nonce            uint64
	functionName     string
	parameters       []xdr.Parameter
	signer           *xdr.Authority
}

//Call
func (cb *CallBuilder) Call(address, channel xdr.ID, nonce uint64) *CallBuilder {
	cb.address = address
	cb.channel = channel
	cb.nonce = nonce
	return cb
}

//Function
func (cb *CallBuilder) Function(name string) *CallBuilder {
	cb.functionName = name
	return cb
}

//Parameters
func (cb *CallBuilder) Parameters(parameters ...xdr.Parameter) *CallBuilder {
	cb.parameters = parameters
	return cb
}

//Sign
func (cb *CallBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	// check required values
	if len(cb.functionName) <= 0 {
		return nil, ErrEmptyFunctionName
	}

	action := xdr.Action{
		Address:   cb.address,
		ChannelID: cb.channel,
		Nonce:     cb.nonce,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeCALL,
			Call: &xdr.Call{
				Function:   cb.functionName,
				Parameters: cb.parameters,
			},
		},
	}

	actionStream, err := action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}

	signatureSlice := ed25519.Sign(pk, actionStream)

	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}

	transaction := &xdr.Transaction{
		Signature: signature,
		Action:    action,
	}

	if cb.signer != nil {
		transaction.Signer = *cb.signer
	}
	return transaction, nil
}
