package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

// TransactionBuilder builds a xdr transaction object. This is a helper struct
// that will build a transaction object.
type TransactionBuilder struct {
	tx *xdr.Transaction
}

// NewTransactionBuilder returns a transactionBuilder with a empty transaction
func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		tx: &xdr.Transaction{},
	}
}

// WithAction sets the action on a transations object through the transactionbuilder. Multiple
// calls will overwrite the action set on a transactions.
func (txb *TransactionBuilder) WithAction(action xdr.Action) *TransactionBuilder {
	txb.tx.Action = action
	return txb
}

// WithAuthority set the authority on a transaction object through the trnasactionbuilder. Multiple
// calls will overwrite the authority on a transaction.
func (txb *TransactionBuilder) WithAuthority(authority xdr.Authority) *TransactionBuilder {
	txb.tx.Signer = authority
	return txb
}

// Sign signs a transaction through the transaction builder. Multiple calls to sign will overwrite
// the transactions signature objects
func (txb *TransactionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if &txb.tx.Action == nil {
		return nil, ErrTransactionActionEmpty
	}
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
