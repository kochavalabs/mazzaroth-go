package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

type UpdatePermissionBuilder struct {
	transaction *xdr.Transaction
	permission  *xdr.Permission
}

func (upb *UpdatePermissionBuilder) UpdatePermission(address, channel [32]byte, nonce uint64) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Action(permissionAction int32) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Address(address [32]byte) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Sign() (*xdr.Transaction, error) {
	return nil, nil
}
