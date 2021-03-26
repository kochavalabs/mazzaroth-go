package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// BuildActionForContractCall generates the action in xdr format for a transaction call.
func BuildActionForContractCall(address, channel xdr.ID, nonce uint64, function string, parameters ...xdr.Parameter) []byte {
	return nil
}

// BuildActionForContractUpdate generates the action in xdr format for a contract update transaction.
func BuildActionForContractUpdate(address, channel xdr.ID, contract, version string) []byte {
	return nil
}

// BuildActionForConfigUpdate generates the action in xdr format for a config update transaction.
func BuildActionForConfigUpdate(address, channel xdr.ID, channelID, contractHash, version, owner, channelName string, admins []string) []byte {
	return nil
}
