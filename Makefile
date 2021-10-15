.PHONY: all $(MAKECMDGOALS)

test:
	go test ./... -count=1

integration:
	go test ./... -count=1 -tags=integration

testall: test integration
