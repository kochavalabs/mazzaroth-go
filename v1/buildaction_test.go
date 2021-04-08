package mazzaroth

import (
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

func IDFromString(s string) xdr.ID {
	var buffer [32]byte

	var maxLenth int = 32
	if len(s) < maxLenth {
		maxLenth = len(s)
	}

	for i := 0; i < len(s); i++ {
		buffer[i] = byte(s[i])
	}

	return xdr.ID(buffer)
}

func TestBuildActionForTransactionCall(t *testing.T) {
	address := IDFromString("the address")
	channel := IDFromString("the channel")
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

func TestBuildActionForUpdateContract(t *testing.T) {
	address := IDFromString("the address")
	channel := IDFromString("the channel")
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

func TestBuildActionForUpdateConfig(t *testing.T) {
	address := IDFromString("the address")
	channel := IDFromString("the channel")
	nonce := uint64(3)

	var config xdr.ChannelConfig
	config.ChannelID = IDFromString("the config address")
	config.ChannelName = "the config channel name"
	config.Admins = []xdr.ID{IDFromString("admin1"), IDFromString("admin2"), IDFromString("admin3")}
	config.Owner = IDFromString("the config owner")
	config.Version = "2.0"

	action := BuildActionForConfigUpdate(address, channel, nonce, config)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, config, *action.Category.Update.ChannelConfig)
	require.Equal(t, xdr.ActionCategoryTypeUPDATE, *&action.Category.Type)
}

func TestBuildActionForPermissionUpdate(t *testing.T) {
	address := IDFromString("the address")
	channel := IDFromString("the channel")
	nonce := uint64(3)

	var permission xdr.Permission
	permission.Key = IDFromString("the key")
	permission.Action = 1

	action := BuildActionForPermissionUpdate(address, channel, nonce, permission)
	require.Equal(t, address, action.Address)
	require.Equal(t, channel, action.ChannelID)
	require.Equal(t, nonce, action.Nonce)
	require.Equal(t, permission, *action.Category.Update.Permission)
	require.Equal(t, xdr.ActionCategoryTypeUPDATE, *&action.Category.Type)
}
