run:
	go test -json ./... | go run main.go

jtest:
	go test -json ./...
