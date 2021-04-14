package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

// BuildActionForTransactionCall generates the action for a transaction call.
func BuildActionForTransactionCall(address, channel xdr.ID, nonce uint64, call xdr.Call) (*xdr.Action, error) {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for transaction call")
	}

	retAction.Category = cat

	return &retAction, nil
}

// BuildActionForContractUpdate generates the action for a contract update transaction.
func BuildActionForContractUpdate(address, channel xdr.ID, nonce uint64, contract xdr.Contract) (*xdr.Action, error) {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Contract: &contract})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for contract update")
	}

	retAction.Category = act

	return &retAction, nil
}

// BuildActionForConfigUpdate generates the action for a config update transaction.
func BuildActionForConfigUpdate(address, channel xdr.ID, nonce uint64, config xdr.ChannelConfig) (*xdr.Action, error) {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{ChannelConfig: &config})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for config update")
	}

	retAction.Category = act

	return &retAction, nil
}

// BuildActionForPermissionUpdate generates the action for a permission update transaction.
func BuildActionForPermissionUpdate(address, channel xdr.ID, nonce uint64, permission xdr.Permission) (*xdr.Action, error) {
	retAction := xdr.Action{}

	retAction.Address = address
	retAction.ChannelID = channel
	retAction.Nonce = nonce

	// The type is safe. No need to check the error.
	act, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Permission: &permission})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for permission update")
	}

	retAction.Category = act

	return &retAction, nil
}
