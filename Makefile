.PHONY: run
run: 
	@go test -json ./... | go run main.go

.PHONY: fail
fail: 
	@TESTFAIL="true" go test -json ./... | go run main.go

.PHONY: jtest
jtest:
	go test -json ./...
