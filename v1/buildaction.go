package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// BuildXDRActionForTransactionCall generates the action in xdr format for a transaction call.
func BuildXDRActionForTransactionCall(address, channel xdr.ID, nonce uint64, function string, parameters ...string) []byte {
	return nil
}

// BuildXDRActionForTransactionContractUpdate generates the action in xdr format for a contract update transaction.
func BuildXDRActionForTransactionContractUpdate(address, channel xdr.ID, contract, version string) []byte {
	return nil
}

// BuildXDRActionForTransactionConfigUpdate generates the action in xdr format for a config update transaction.
func BuildXDRActionForTransactionConfigUpdate(address, channel xdr.ID, channelID, contractHash, version, owner, channelName string, admins []string) []byte {
	return nil
}

// BuildXDRActionForReadOnlyCall generates the action in xdr format for a read only transaction call.
func BuildXDRActionForReadOnlyCall(address, channel xdr.ID, function string, parameters ...string) []byte {
	return nil
}
