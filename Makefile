build:
	go build -o bin/pancakes cmd/main.go

run:
	./bin/pancakes

test:
	go test -v ./...