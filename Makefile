all: vet test

vet:
	go vet ./...

test:
	go test ./...