.PHONY: build
build:
	go build -o bin ./...

.PHONY: run
run: build
	./bin/cmd

.PHONY: start-db
start-db:
	docker-compose up -d
