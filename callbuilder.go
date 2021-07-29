package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type CallBuilder struct {
	address, channel      *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	functionName          string
	arguments             []xdr.Argument
	signer                *xdr.Authority
}

//Call
func (cb *CallBuilder) Call(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *CallBuilder {
	cb.address = address
	cb.channel = channel
	cb.nonce = nonce
	cb.blockExpirationNumber = blockExpirationNumber
	return cb
}

//Function
func (cb *CallBuilder) Function(name string) *CallBuilder {
	cb.functionName = name
	return cb
}

//Arguments
func (cb *CallBuilder) Arguments(arguments ...xdr.Argument) *CallBuilder {
	cb.arguments = arguments
	return cb
}

//Sign
func (cb *CallBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	// check required values
	if len(cb.functionName) <= 0 {
		return nil, ErrEmptyFunctionName
	}

	action := xdr.Action{
		Address:               *cb.address,
		ChannelID:             *cb.channel,
		Nonce:                 cb.nonce,
		BlockExpirationNumber: cb.blockExpirationNumber,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeCALL,
			Call: &xdr.Call{
				Function:  cb.functionName,
				Arguments: cb.arguments,
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
