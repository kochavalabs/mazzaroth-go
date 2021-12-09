package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

func mazzarothClient(this js.Value, args []js.Value) interface{} {
	var client mazzaroth.Client
	var err error

	if len(args) == 0 {
		// allow client to be created with default values
		client, err = mazzaroth.NewMazzarothClient()
		if err != nil {
			return err.Error()
		}
	}

	wrapperClient := &mazzarothJsWrapperClient{
		client: client,
	}

	// list of client functions
	return js.ValueOf(map[string]interface{}{
		"AccountLookup":       wrapperClient.accountLookup(),
		"AuthorizationLookup": wrapperClient.authorizationLookup(),
		"BlockHeaderLookup":   wrapperClient.blockHeaderLookup(),
		"BlockHeaderList":     wrapperClient.blockHeaderList(),
		"BlockHeight":         wrapperClient.blockHeight(),
		"BlockLookup":         wrapperClient.blockLookup(),
		"BlockList":           wrapperClient.blockList(),
		"ChannelAbi":          wrapperClient.channelAbi(),
		"ChannelConfig":       wrapperClient.channelConfig(),
		"ReceiptLookup":       wrapperClient.receiptLookup(),
		"TransactionLookup":   wrapperClient.transactionLookup(),
		"TransactionSubmit":   wrapperClient.transactionSubmit(),
	})
}

func transactionBuilder(this js.Value, args []js.Value) interface{} {
	txBuilder := &transactionBuilderJsWrapper{}
	return js.ValueOf(map[string]interface{}{
		"Account":       txBuilder.account(),
		"Authorization": txBuilder.authorization(),
		"Call":          txBuilder.call(),
		"Config":        txBuilder.config(),
		"Contract":      txBuilder.contract(),
	})
}

func main() {
	c := make(chan struct{})
	js.Global().Set("NewMazzarothClient", js.FuncOf(mazzarothClient))
	js.Global().Set("TransactionBuilder", js.FuncOf(transactionBuilder))

	<-c
	// Must keep go program alive when instantiated to allow access to functions
}
