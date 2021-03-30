package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// BuildActionForTransactionCall generates the action for a transaction call.
func BuildActionForTransactionCall(address, channel xdr.ID, nonce uint64, call xdr.Call) (*xdr.Action, error) {
	return nil, nil
}

// BuildActionForContractUpdate generates the action for a contract update transaction.
func BuildActionForContractUpdate(address, channel xdr.ID, nonce uint64, contract xdr.Contract) (*xdr.Action, error) {
	return nil, nil
}

// BuildActionForConfigUpdate generates the action for a config update transaction.
func BuildActionForConfigUpdate(address, channel xdr.ID, nonce uint64, config xdr.ChannelConfig) (*xdr.Action, error) {
	return nil, nil
}

// BuildActionForPermissionUpdate generates the action for a permission update transaction.
func BuildActionForPermissionUpdate(address, channel xdr.ID, nonce uint64, permission xdr.Permission) (*xdr.Action, error) {
	return nil, nil
}
