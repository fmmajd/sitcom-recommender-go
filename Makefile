all: say_hello

say_hello:
	@echo "Hello"

bash:
	@docker-compose run --rm go bash

run:
	@docker-compose run --rm go go run main.go

build:
	@docker-compose run --rm go build main.go
