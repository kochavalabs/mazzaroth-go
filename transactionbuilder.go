package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

// TransactionBuilder builds a xdr transaction object. This is a helper struct
// that will build a transaction object.
type TransactionBuilder struct {
	sender  *xdr.ID
	channel *xdr.ID
}

// Transaction returns a transactionBuilder with a empty xdr.transaction
func Transaction(sender, channel *xdr.ID) *TransactionBuilder {
	return &TransactionBuilder{
		sender:  sender,
		channel: channel,
	}
}

// Acount
func (txb *TransactionBuilder) Account(nonce, blockExpirationNumber uint64) *AccountBuilder {
	return new(AccountBuilder).Account(txb.sender, txb.channel, nonce, blockExpirationNumber)
}

// Authorization
func (txb *TransactionBuilder) Authorization(nonce, blockExpirationNumber uint64) *AuthorizationBuilder {
	return new(AuthorizationBuilder).Authorization(txb.sender, txb.channel, nonce, blockExpirationNumber)
}

// Call
func (txb *TransactionBuilder) Call(nonce, blockExpirationNumber uint64) *CallBuilder {
	return new(CallBuilder).Call(txb.sender, txb.channel, nonce, blockExpirationNumber)
}

// Config
func (txb *TransactionBuilder) Config(nonce, blockExpirationNumber uint64) *ConfigBuilder {
	return new(ConfigBuilder).Config(txb.sender, txb.channel, nonce, blockExpirationNumber)
}

// Contract
func (txb *TransactionBuilder) Contract(nonce, blockExpirationNumber uint64) *ContractBuilder {
	return new(ContractBuilder).Contract(txb.sender, txb.channel, nonce, blockExpirationNumber)
}
