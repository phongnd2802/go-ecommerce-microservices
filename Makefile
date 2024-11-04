GOOSE_DBSTRING ?= "user=root password=secret host=127.0.0.1 port=5432 dbname=ecommerce sslmode=disable"
GOOSE_MIGRATION_DIR ?= migrations
GOOSE_DRIVER ?= postgres

network:
	docker network create ecommerce-network

postgres:
	docker run --name ecommerce-db --network ecommerce-network -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:14-alpine

createdb:
	docker exec -it ecommerce-db createdb --username=root --owner=root ecommerce

dropdb:
	docker exec -it ecommerce-db dropdb ecommerce

redis:
	docker run --name redis --network ecommerce-network -p 6379:6379 -d redis:7.4.1-alpine

migration:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) -s create $(name) sql

db-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

db-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down-to 0

db-cli:
	docker exec -it ecommerce-db psql -U root -d ecommerce

proto:
	scripts/proto_gen.sh

sqlc:
	sqlc generate

wire:
	cd internal/user/app && wire && cd -

evans:
	evans --host localhost --port $(port) -r repl

redis-cli:
	docker exec -it redis redis-cli

user:
	go run ./cmd/user

proxy:
	go run ./cmd/proxy

docker-compose-dev-up:
	docker compose -f docker-compose-dev.yml up

docker-compose-dev-down:
	docker compose -f docker-compose-dev.yml down

templ:
	templ generate

web:
	go run ./web/
	
.PHONY: network postgres createdb dropdb db-up db-down migration db-cli proto sqlc evans redis redis-cli server \
	user proxy docker-compose-dev-up docker-compose-dev-down wire templ web
