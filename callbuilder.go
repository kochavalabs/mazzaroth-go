package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type CallBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	functionName          string
	arguments             []xdr.Argument
}

// Call
func (cb *CallBuilder) Call(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *CallBuilder {
	cb.sender = sender
	cb.channel = channel
	cb.nonce = nonce
	cb.blockExpirationNumber = blockExpirationNumber
	return cb
}

// Function
func (cb *CallBuilder) Function(name string) *CallBuilder {
	cb.functionName = name
	return cb
}

// Arguments
func (cb *CallBuilder) Arguments(arguments ...xdr.Argument) *CallBuilder {
	cb.arguments = arguments
	return cb
}

// Sign
func (cb *CallBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	// check required values
	if len(cb.functionName) <= 0 {
		return nil, ErrEmptyFunctionName
	}

	data := &xdr.Data{
		ChannelID:             *cb.channel,
		Nonce:                 cb.nonce,
		BlockExpirationNumber: cb.blockExpirationNumber,
		Category: xdr.Category{
			Type: xdr.CategoryTypeCALL,
			Call: &xdr.Call{
				Function:  cb.functionName,
				Arguments: cb.arguments,
			},
		},
	}

	dataBytes, err := data.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in data.MarshalBinary")
	}

	signer, err := xdr.IDFromPublicKey(pk.Public())
	if err != nil {
		return nil, errors.Wrap(err, "in xdr.IDFromPublicKey")
	}

	signatureSlice := ed25519.Sign(pk, dataBytes)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}

	transaction := &xdr.Transaction{
		Sender:    *cb.sender,
		Signer:    signer,
		Signature: signature,
		Data:      data,
	}

	return transaction, nil
}
