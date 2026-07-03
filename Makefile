build:
	go build -o ./bin/my-go-blockchain

run: build
	./bin/my-go-blockchain

test:
	go test -v ./...