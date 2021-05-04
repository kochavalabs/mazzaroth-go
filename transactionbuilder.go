package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

// TransactionBuilder builds a xdr transaction object. This is a helper struct
// that will build a transaction object.
type TransactionBuilder struct {
	transaction *xdr.Transaction
	authority   *xdr.Authority
}

// Transaction returns a transactionBuilder with a empty xdr.transaction
func Transaction() *TransactionBuilder {
	return &TransactionBuilder{
		transaction: &xdr.Transaction{},
	}
}

// Authority - call out public key
func (txb *TransactionBuilder) Authority(address [32]byte) *TransactionBuilder {
	origin := xdr.ID(address)
	txb.authority = &xdr.Authority{
		Type:   xdr.AuthorityTypePERMISSIONED,
		Origin: &origin,
	}
	return nil
}

func (txb *TransactionBuilder) Call(address, channel [32]byte, nonce uint64) *CallBuilder {
	callbuilder := new(CallBuilder)
	if txb.authority != nil {
		callbuilder.signer = txb.authority
	}
	return callbuilder.Call(address, channel, nonce)
}

func (txb *TransactionBuilder) UpdateConfig(address, channel [32]byte, nonce uint64) *UpdateConfigBuilder {
	updateConfigBuilder := new(UpdateConfigBuilder)
	return updateConfigBuilder.UpdateConfig(address, channel, nonce)
}

func (txb *TransactionBuilder) UpdateContract(address, channel [32]byte, nonce uint64) *UpdateContractBuilder {
	updateContractBuilder := new(UpdateContractBuilder)
	return updateContractBuilder.UpdateContract(address, channel, nonce)
}

func (txb *TransactionBuilder) UpdatePermission(address, channel [32]byte, nonce uint64) *UpdatePermissionBuilder {
	updatePermissionBuilder := new(UpdatePermissionBuilder)
	return updatePermissionBuilder.UpdatePermission(address, channel, nonce)
}
