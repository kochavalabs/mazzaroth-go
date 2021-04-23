package client

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type TransactionBuilder struct {
	tx *xdr.Transaction
}

func NewTransactionBuilder(tx *xdr.Transaction) *TransactionBuilder {
	return &TransactionBuilder{
		tx: tx,
	}
}

func (txb *TransactionBuilder) WithAction(action xdr.Action) *TransactionBuilder {
	txb.tx.Action = action
	return txb
}

func (txb *TransactionBuilder) WithAuthority(authority xdr.Authority) *TransactionBuilder {
	txb.tx.Signer = authority
	return txb
}

func (txb *TransactionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	actionStream, err := txb.tx.Action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}
	signatureSlice := ed25519.Sign(pk, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}
	txb.tx.Signature = signature
	if &txb.tx.Signer == nil {
		txb.tx.Signer, err = xdr.NewAuthority(xdr.AuthorityTypeNONE, nil)
		if err != nil {
			return nil, errors.Wrap(err, "in xdr.NewAuthority(xdr.AuthorityTypeNONE)")
		}
	}
	return txb.tx, nil
}
