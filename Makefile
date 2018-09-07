# Go parameters
BINARY_NAME=SwagIP-go

all: test build
build:
	packr
	go build -o $(BINARY_NAME)_ -v
test:
	go test -v ./...
clean:
	go clean
	packr clean
	rm -f $(BINARY_NAME)_*
run:
	go build -o $(BINARY_NAME)_ -v ./...
	./$(BINARY_NAME)_
deps:
	go mod download

# Cross compilation
build-all:
	packr
	gox -output "bin/{{.Dir}}_{{.OS}}_{{.Arch}}"