package main

import (
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

func delete(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return map[string]interface{}{
			"Sign": sign(contractBuilder),
		}
	})
}

func deploy(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 3 {
			return errors.New("missing arguments")
		}
		version := args[0].String()
		abi := &xdr.Abi{}
		if err := json.Unmarshal([]byte(args[1].String()), abi); err != nil {
			return err.Error()
		}
		contractBytes := []byte(args[2].String())
		contractBuilder.Deploy(version, *abi, contractBytes)
		return map[string]interface{}{
			"Sign": sign(contractBuilder),
		}
	})
}

func pause(contractBuilder *mazzaroth.ContractBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		contractBuilder.Pause(args[0].Bool())
		return map[string]interface{}{
			"Sign": sign(contractBuilder),
		}
	})
}
