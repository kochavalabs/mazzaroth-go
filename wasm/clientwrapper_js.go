package main

import (
	"context"
	"encoding/json"
	"syscall/js"
	"time"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

const timeout = 10 * time.Second

type mazzarothJsWrapperClient struct {
	client mazzaroth.Client
}

func (m *mazzarothJsWrapperClient) blockHeaderLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		blockHeader, err := m.client.BlockHeaderLookup(ctx, args[0].String(), args[1].String())
		if err != nil {
			return err.Error()
		}
		blockHeaderJson, err := json.Marshal(blockHeader)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(blockHeaderJson))
	})
}

func (m *mazzarothJsWrapperClient) blockHeaderList() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		blockHeaderList, err := m.client.BlockHeaderList(ctx, args[0].String(), args[1].Int(), args[2].Int())
		if err != nil {
			return err.Error()
		}
		blockHeaderListJson, err := json.Marshal(blockHeaderList)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(blockHeaderListJson))
	})
}

func (m *mazzarothJsWrapperClient) blockHeight() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		blockHeight, err := m.client.BlockHeight(ctx, args[0].String())
		if err != nil {
			return err.Error()
		}
		blockHeightJson, err := json.Marshal(blockHeight)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(blockHeightJson))
	})
}

func (m *mazzarothJsWrapperClient) blockLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		block, err := m.client.BlockLookup(ctx, args[0].String(), args[1].String())
		if err != nil {
			return err.Error()
		}
		blockJson, err := json.Marshal(block)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(blockJson))
	})
}

func (m *mazzarothJsWrapperClient) blockList() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		blockList, err := m.client.BlockList(ctx, args[0].String(), args[1].Int(), args[2].Int())
		if err != nil {
			return err.Error()
		}
		blockListJson, err := json.Marshal(blockList)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(blockListJson))
	})
}

func (m *mazzarothJsWrapperClient) channelAbi() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		channelAbi, err := m.client.ChannelAbi(ctx, args[0].String())
		if err != nil {
			return err.Error()
		}
		channelAbiJson, err := json.Marshal(channelAbi)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(channelAbiJson))
	})
}

func (m *mazzarothJsWrapperClient) receiptLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		receipt, err := m.client.ReceiptLookup(ctx, args[0].String(), args[1].String())
		if err != nil {
			return err.Error()
		}
		receiptJson, err := json.Marshal(receipt)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(receiptJson))
	})
}

func (m *mazzarothJsWrapperClient) transactionLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		tx, err := m.client.TransactionLookup(ctx, args[0].String(), args[1].String())
		if err != nil {
			return err.Error()
		}
		txJson, err := json.Marshal(tx)
		if err != nil {
			return err.Error()
		}
		return js.ValueOf(string(txJson))
	})
}

func (m *mazzarothJsWrapperClient) transactionSubmit() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		tx := &xdr.Transaction{}
		if err := json.Unmarshal([]byte(args[0].String()), tx); err != nil {
			return err.Error()
		}

		id, receipt, err := m.client.TransactionSubmit(ctx, tx)
		if err != nil {
			return err.Error()
		}
		if receipt != nil {
			receiptJson, err := json.Marshal(receipt)
			if err != nil {
				return err.Error()
			}
			return map[string]interface{}{
				"id":      &id,
				"receipt": receiptJson,
			}
		}
		return &id
	})
}
