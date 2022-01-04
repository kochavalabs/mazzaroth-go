package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

// TransactionBuilder builds a xdr transaction object. This is a helper struct
// that will build a transaction object.
type TransactionBuilder struct {
	sender  xdr.ID
	channel xdr.ID
}

// Transaction returns a transactionBuilder with a empty xdr.transaction
func Transaction(sender, channel xdr.ID) *TransactionBuilder {
	return &TransactionBuilder{
		sender:  sender,
		channel: channel,
	}
}

// Call
func (txb *TransactionBuilder) Call(nonce, blockExpirationNumber uint64) *CallBuilder {
	return new(CallBuilder).Call(&txb.sender, &txb.channel, nonce, blockExpirationNumber)
}

// Contract
func (txb *TransactionBuilder) Contract(nonce, blockExpirationNumber uint64) *ContractBuilder {
	return new(ContractBuilder).Contract(&txb.sender, &txb.channel, nonce, blockExpirationNumber)
}
