
all: clean test build

clean:
	@rm -rf worktracker

test:
	@go test ./wtime

build:
	@go build .
