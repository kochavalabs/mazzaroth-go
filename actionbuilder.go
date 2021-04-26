package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

// ActionBuilder builds a xdr action object. This is a helper struct
// that will build an action object with type of Call,ContractUpdate
// PermissionUpdate or ConfigUpdate.
type ActionBuilder struct {
	action *xdr.Action
}

// NewActionBuilder returns empty actionbuilder
func NewActionBuilder() *ActionBuilder {
	return &ActionBuilder{
		action: &xdr.Action{},
	}
}

// WithAddress sets the Action address through the action builder. Multiple
// calls will overwrite the address on an action
func (ab *ActionBuilder) WithAddress(address xdr.ID) *ActionBuilder {
	ab.action.Address = address
	return ab
}

// WithChannel sets the Action channel id through the action builder. Multiple
// calls will overwrite the Channel id on an action
func (ab *ActionBuilder) WithChannel(channelID xdr.ID) *ActionBuilder {
	ab.action.ChannelID = channelID
	return ab
}

// WithNonce sets the Action nonce value through the action builder. Multiple
// calls will overwrite the Nonce value on an action
func (ab *ActionBuilder) WithNonce(nonce uint64) *ActionBuilder {
	ab.action.Nonce = nonce
	return ab
}

// Call sets the Action Call Category through the action builder. Overwrites any pervious
// set action category by the actionbuilder
func (ab *ActionBuilder) Call(call xdr.Call) (*xdr.Action, error) {
	if err := ab.checkRequiredValues(); err != nil {
		return nil, err
	}

	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for transaction call")
	}
	ab.action.Category = cat
	return ab.action, nil
}

// ContractUpdate sets the Action Update Category through the action builder. Overwrites any pervious
// set action category by the actionbuilder
func (ab *ActionBuilder) ContractUpdate(contract xdr.Contract) (*xdr.Action, error) {
	if err := ab.checkRequiredValues(); err != nil {
		return nil, err
	}

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

// PermissionUpdate sets the Action Permission Category through the action builder. Overwrites any pervious
// set action category by the actionbuilder
func (ab *ActionBuilder) PermissionUpdate(permission xdr.Permission) (*xdr.Action, error) {
	if err := ab.checkRequiredValues(); err != nil {
		return nil, err
	}

	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{Permission: &permission})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for permission update")
	}
	ab.action.Category = cat
	return ab.action, nil
}

// ConfigUpdate sets the Action Config Category through the action builder. Overwrites any pervious
// set action category by the actionbuilder
func (ab *ActionBuilder) ConfigUpdate(config xdr.ChannelConfig) (*xdr.Action, error) {
	if err := ab.checkRequiredValues(); err != nil {
		return nil, err
	}

	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeUPDATE, xdr.Update{ChannelConfig: &config})
	if err != nil {
		return nil, errors.Wrap(err, "could not crate a new action category for config update")
	}
	ab.action.Category = cat
	return ab.action, nil
}

func (ab *ActionBuilder) checkRequiredValues() error {
	if &ab.action.Address == nil {
		return ErrActionAddressNil
	}
	if &ab.action.ChannelID == nil {
		return ErrChannelIDNil
	}
	return nil
}
