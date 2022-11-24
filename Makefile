all: tidy build

tidy:
	go mod tidy; go mod vendor

build:
	go build -o bin/main cmd/main.go
	go build -o bin/client client/client.go
