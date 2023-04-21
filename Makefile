.PHONY: migrate-up
migrate-up:
	sql-migrate up

.PHONY: build
build:
	go build -o bin ./...

.PHONY: run
run: build
	./bin/cmd

.PHONY: start-services
start-db:
	docker-compose up 
