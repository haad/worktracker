
all: clean test build

clean:
	@rm -rf worktracker

test:
	@go test .

build:
	@go build .