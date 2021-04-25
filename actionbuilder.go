package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

// TODO : Add go docs and safety checks

type ActionBuilder struct {
	action *xdr.Action
}

func NewActionBuilder() *ActionBuilder {
	return &ActionBuilder{
		action: &xdr.Action{},
	}
}

func (ab *ActionBuilder) WithAddress(address xdr.ID) *ActionBuilder {
	ab.action.Address = address
	return ab
}

func (ab *ActionBuilder) WithChannel(channelID xdr.ID) *ActionBuilder {
	ab.action.ChannelID = channelID
	return ab
}

func (ab *ActionBuilder) WithNonce(nonce uint64) *ActionBuilder {
	ab.action.Nonce = nonce
	return ab
}

func (ab *ActionBuilder) Call(call xdr.Call) (*xdr.Action, error) {
	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for transaction call")
	}
	ab.action.Category = cat
	return ab.action, nil
}

func (ab *ActionBuilder) ContractUpdate(contract xdr.Contract) (*xdr.Action, error) {

	update, err := xdr.NewUpdate(xdr.UpdateTypeCONTRACT, contract)
	if err != nil {
		return nil, errors.Wrap(err, "could not create a new update for type UpdateTypeCONTRACT")
	}

	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, update)
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for contract update")
	}
	ab.action.Category = cat
	return ab.action, nil
}

func (ab *ActionBuilder) PermissionUpdate(permission xdr.Permission) (*xdr.Action, error) {
	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Permission: &permission})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for permission update")
	}
	ab.action.Category = cat
	return ab.action, nil
}

func (ab *ActionBuilder) ConfigUpdate(config xdr.ChannelConfig) (*xdr.Action, error) {
	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{ChannelConfig: &config})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for config update")
	}
	ab.action.Category = cat
	return ab.action, nil
}
