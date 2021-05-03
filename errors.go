package mazzaroth

import "errors"

var (
	// ErrEmptyServerList triggered if empty server list is passed
	ErrEmptyServerList = errors.New("unable to create the server selector with an empty server list")
	//ErrTransactionActionEmpty triggered if empty action is passed on transaction
	ErrTransactionActionEmpty = errors.New("transaction action can not be empty")
	//ErrActionAddressNil triggered if action address is nil
	ErrActionAddressNil = errors.New("action address can not be nil")
	//ErrChannelIDNil triggered if action channel id is nil
	ErrChannelIDNil = errors.New("action channel id can not be nil")
	//ErrEmptyFunction name triggered if an empty function name is used to sign a call transaction
	ErrEmptyFunctionName = errors.New("function name can not be empty")
)
