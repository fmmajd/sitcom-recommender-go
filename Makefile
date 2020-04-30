all: say_hello

say_hello:
	@echo "Hello"

run:
	@docker-compose run --rm go go run main.go

build:
	@docker-compose run --rm go build main.go
