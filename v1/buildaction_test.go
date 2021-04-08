package mazzaroth

import (
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

func idFromString(s string) xdr.ID {
	var id xdr.ID

	copy(id[:], s)

	return id
}

// TestBuildActionForTransactionCall tests the BuildActionForTransactionCall function.
func TestBuildActionForTransactionCall(t *testing.T) {
	address := idFromString("the address")
	channel := idFromString("the channel")
	nonce := uint64(3)

	var call xdr.Call
	call.Function = "foo"
	call.Parameters = []xdr.Parameter{xdr.Parameter("p1"), xdr.Parameter("p2")}

	action := BuildActionForTransactionCall(address, channel, nonce, call)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, call, *action.Category.Call)
	require.Equal(t, xdr.ActionCategoryTypeCALL, *&action.Category.Type)
}

// TestBuildActionForUpdateContract tests the TestBuildActionForUpdateContract function.
func TestBuildActionForUpdateContract(t *testing.T) {
	address := idFromString("the address")
	channel := idFromString("the channel")
	nonce := uint64(3)

	var contract xdr.Contract
	contract.Version = "1"
	contract.Contract = []byte("the contract")

	action := BuildActionForContractUpdate(address, channel, nonce, contract)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, contract, *action.Category.Update.Contract)
	require.Equal(t, xdr.ActionCategoryTypeUPDATE, *&action.Category.Type)
}

// TestBuildActionForUpdateConfig tests the TestBuildActionForUpdateConfig function.
func TestBuildActionForUpdateConfig(t *testing.T) {
	address := idFromString("the address")
	channel := idFromString("the channel")
	nonce := uint64(3)

	var config xdr.ChannelConfig
	config.ChannelID = idFromString("the config address")
	config.ChannelName = "the config channel name"
	config.Admins = []xdr.ID{idFromString("admin1"), idFromString("admin2"), idFromString("admin3")}
	config.Owner = idFromString("the config owner")
	config.Version = "2.0"

	action := BuildActionForConfigUpdate(address, channel, nonce, config)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, config, *action.Category.Update.ChannelConfig)
	require.Equal(t, xdr.ActionCategoryTypeUPDATE, *&action.Category.Type)
}

// TestBuildActionForPermissionUpdate tests the TestBuildActionForPermissionUpdate function.
func TestBuildActionForPermissionUpdate(t *testing.T) {
	address := idFromString("the address")
	channel := idFromString("the channel")
	nonce := uint64(3)

	var permission xdr.Permission
	permission.Key = idFromString("the key")
	permission.Action = 1

	action := BuildActionForPermissionUpdate(address, channel, nonce, permission)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, permission, *action.Category.Update.Permission)
	require.Equal(t, xdr.ActionCategoryTypeUPDATE, *&action.Category.Type)
}
