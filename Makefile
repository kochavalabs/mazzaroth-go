.PHONY: $(MAKECMDGOALS)

test:
	go test -v ./... 

build-wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/mazzarothclient.wasm ./wasm/*.go
