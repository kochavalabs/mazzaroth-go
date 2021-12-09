package main

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
	"syscall/js"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

type transactionBuilderJsWrapper struct{}

func (tb *transactionBuilderJsWrapper) account() js.Func {
	accountBuilder := &mazzaroth.AccountBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		accountBuilder.Account(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"Alias": accountAlias(accountBuilder),
		}
	})
}

func (tb *transactionBuilderJsWrapper) authorization() js.Func {
	authorizationBuilder := &mazzaroth.AuthorizationBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(authorizationBuilder)
		return nil
	})
}

func (tb *transactionBuilderJsWrapper) call() js.Func {
	callBuilder := &mazzaroth.CallBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(callBuilder)
		return nil
	})
}

func (tb *transactionBuilderJsWrapper) config() js.Func {
	configBuilder := &mazzaroth.ConfigBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(configBuilder)
		return nil
	})
}

func (tb *transactionBuilderJsWrapper) contract() js.Func {
	contractBuilder := &mazzaroth.ContractBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(contractBuilder)
		return nil
	})
}

func accountAlias(accountBuilder *mazzaroth.AccountBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		accountBuilder.Alias(args[0].String())
		return map[string]interface{}{
			"Sign": accountSign(accountBuilder),
		}
	})
}

func accountSign(accountBuilder *mazzaroth.AccountBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		privateKeyBytes, err := crypto.FromHex(args[0].String())
		if err != nil {
			return err.Error()
		}

		key := ed25519.NewKeyFromSeed(privateKeyBytes)

		tx, err := accountBuilder.Sign(key)
		if err != nil {
			return err.Error()
		}

		b := new(bytes.Buffer)
		if _, err := xdr.Marshal(b, tx); err != nil {
			return err.Error()
		}

		return js.ValueOf(string(b.Bytes()))
	})
}

func authorizationAccount(authorizationBuilder *mazzaroth.AuthorizationBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func authorizationAuthorize(authorizationBuilder *mazzaroth.AuthorizationBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func authorizationSign(authorizationBuilder *mazzaroth.AuthorizationBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func callFunction(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func callArguments(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func callSign(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractBytes(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractVersion(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractAbi(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractSign(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func configOwner(configBuilder *mazzaroth.ConfigBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func configAdmins(configBuilder *mazzaroth.ConfigBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func configSign(configBuilder *mazzaroth.ConfigBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}
