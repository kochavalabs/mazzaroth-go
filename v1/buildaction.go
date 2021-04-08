package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// BuildActionForTransactionCall generates the action for a transaction call.
func BuildActionForTransactionCall(address, channel xdr.ID, nonce uint64, call xdr.Call) *xdr.Action {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	cat, _ := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)

	retAction.Category = cat

	return &retAction
}

// BuildActionForContractUpdate generates the action for a contract update transaction.
func BuildActionForContractUpdate(address, channel xdr.ID, nonce uint64, contract xdr.Contract) *xdr.Action {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, _ := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Contract: &contract})

	retAction.Category = act

	return &retAction
}

// BuildActionForConfigUpdate generates the action for a config update transaction.
func BuildActionForConfigUpdate(address, channel xdr.ID, nonce uint64, config xdr.ChannelConfig) *xdr.Action {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, _ := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{ChannelConfig: &config})

	retAction.Category = act

	return &retAction
}

// BuildActionForPermissionUpdate generates the action for a permission update transaction.
func BuildActionForPermissionUpdate(address, channel xdr.ID, nonce uint64, permission xdr.Permission) *xdr.Action {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, _ := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Permission: &permission})

	retAction.Category = act

	return &retAction
}
