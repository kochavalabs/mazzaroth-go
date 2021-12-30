package main

import (
	"crypto/ed25519"
	"encoding/json"
	"syscall/js"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

type signer interface {
	Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error)
}

func sign(signer signer) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		privateKeyBytes, err := crypto.FromHex(args[0].String())
		if err != nil {
			return err.Error()
		}
		key := ed25519.NewKeyFromSeed(privateKeyBytes)

		tx, err := signer.Sign(key)
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
