GOOSE_DBSTRING ?= "user=root password=secret host=127.0.0.1 port=5432 dbname=ecommerce sslmode=disable"
GOOSE_MIGRATION_DIR ?= database/migrations
GOOSE_DRIVER ?= postgres

network:
	docker network create ecommerce-network

postgres:
	docker run --name ecommerce-db --network ecommerce-network -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:14-alpine

createdb:
	docker exec -it ecommerce-db createdb --username=root --owner=root ecommerce

dropdb:
	docker exec -it ecommerce-db dropdb ecommerce

migration:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) -s create $(name) sql

db-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

db-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down-to 0

db-cli:
	docker exec -it ecommerce-db psql -U root -d ecommerce

.PHONY: network postgres createdb dropdb db-up db-down migration db-cli
