# Go parameters
BINARY_NAME=SwagIP

all: test build
build:
	packr
	go build -o $(BINARY_NAME) -v
test:
	go test -v ./...
clean:
	go clean
	packr clean
	rm -f $(BINARY_NAME)*
run:
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	go mod download

# Cross compilation
build-all:
	packr
	gox -output "bin/{{.Dir}}_{{.OS}}_{{.Arch}}"