package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

func contractAbi(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		abi := xdr.Abi{}
		if err := json.Unmarshal([]byte(args[0].String()), abi); err != nil {
			return err.Error()
		}
		contractBuilder.Abi(abi)
		return map[string]interface{}{
			"ContractBytes": contractVersion(contractBuilder),
		}
	})
}

func contractBytes(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		contractBuilder.ContractBytes([]byte(args[0].String()))
		return map[string]interface{}{
			"Version": contractVersion(contractBuilder),
		}
	})
}

func contractVersion(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		contractBuilder.Version(args[0].String())
		return map[string]interface{}{
			"Sign": sign(contractBuilder),
		}
	})
}
