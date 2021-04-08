test:
	go test ./... -count=1

integration:
	go test ./... -count=1 -tags=integration