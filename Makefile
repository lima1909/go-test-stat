.PHONY: run
run: 
	@go test -count=1 -json ./... | go run main.go

.PHONY: fail
fail: 
	@TESTFAIL="true" go test -count=1  -json ./... | go run main.go

.PHONY: jtest
jtest:
	go test -count=1  -json ./...
