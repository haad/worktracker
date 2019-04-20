
all: clean test build

clean:
	@rm -rf worktracker

test:
	@go test -v ./wtime

build:
	@go build .
