package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

// tx := mazzaroth.Transaction(....).Call().function().Parameters(field...).Sign()
// tx := Transaction().UpdateConfig(Channel 2 ).Owner("").ChannelName("").Admins().Sign()
// tx := Transaction().UpdateContract("address","Channel",nonce).Contract("Bytes").WithVersion("").Sign()
// tx := Trnasaction().UpdatePermission("address","Channel",nonce).Key("").Action("enum")
// tx := Transaction(Address, Nonce).Call(channel).function(name).Parmaeters(field).Sign()

// TransactionBuilder builds a xdr transaction object. This is a helper struct
// that will build a transaction object.
type TransactionBuilder struct {
	transaction *xdr.Transaction
}

// Transaction returns a transactionBuilder with a empty xdr.transaction
func Transaction() *TransactionBuilder {
	return &TransactionBuilder{
		transaction: &xdr.Transaction{},
	}
}
func (txb *TransactionBuilder) Call() *CallBuilder {
	call := new(xdr.Call)
	cat, _ := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)
	txb.transaction.Action = xdr.Action{
		Category: cat,
	}
	return &CallBuilder{
		transaction: txb.transaction,
	}
}

func (txb *TransactionBuilder) UpdateConfig() *UpdateConfigBuilder {
	return &UpdateConfigBuilder{}
}

func (txb *TransactionBuilder) UpdateContract() *UpdateContractBuilder {
	return &UpdateContractBuilder{}
}

func (txb *TransactionBuilder) UpdatePermission() *UpdatePermissionBuilder {
	return &UpdatePermissionBuilder{}
}

// WithAction sets the action on a transations object through the transactionbuilder. Multiple
// calls will overwrite the action set on a transactions.
func (txb *TransactionBuilder) WithAction(action xdr.Action) *TransactionBuilder {
	txb.transaction.Action = action
	return txb
}

// WithAuthority set the authority on a transaction object through the transactionbuilder. Multiple
// calls will overwrite the authority on a transaction.
func (txb *TransactionBuilder) WithAuthority(authority xdr.Authority) *TransactionBuilder {
	txb.transaction.Signer = authority
	return txb
}

// Sign signs a transaction through the transaction builder. Multiple calls to sign will overwrite
// the transactions signature objects
func (txb *TransactionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if &txb.transaction.Action == nil {
		return nil, ErrTransactionActionEmpty
	}
	actionStream, err := txb.transaction.Action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}

	signatureSlice := ed25519.Sign(pk, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}
	txb.transaction.Signature = signature
	if &txb.transaction.Signer == nil {
		txb.transaction.Signer, err = xdr.NewAuthority(xdr.AuthorityTypeNONE, nil)
		if err != nil {
			return nil, errors.Wrap(err, "in xdr.NewAuthority(xdr.AuthorityTypeNONE)")
		}
	}
	return txb.transaction, nil
}
