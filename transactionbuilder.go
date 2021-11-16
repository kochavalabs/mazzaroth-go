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
func (txb *TransactionBuilder) Authority(address *xdr.ID) *TransactionBuilder {
	if address != nil {
		txb.authority = &xdr.Authority{
			Type:   xdr.AuthorityTypeAUTHORIZED,
			Origin: address,
		}
	}
	return txb
}

func (txb *TransactionBuilder) Call(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *CallBuilder {
	callbuilder := new(CallBuilder)
	if txb.authority != nil {
		callbuilder.signer = txb.authority
	}
	return callbuilder.Call(address, channel, nonce, blockExpirationNumber)
}

func (txb *TransactionBuilder) UpdateConfig(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateConfigBuilder {
	updateConfigBuilder := new(UpdateConfigBuilder)
	return updateConfigBuilder.UpdateConfig(address, channel, nonce, blockExpirationNumber)
}

func (txb *TransactionBuilder) UpdateContract(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateContractBuilder {
	updateContractBuilder := new(UpdateContractBuilder)
	return updateContractBuilder.UpdateContract(address, channel, nonce, blockExpirationNumber)
}

func (txb *TransactionBuilder) UpdatePermission(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateAuthorizationBuilder {
	updatePermissionBuilder := new(UpdateAuthorizationBuilder)
	return updatePermissionBuilder.UpdatePermission(address, channel, nonce, blockExpirationNumber)
}
