PJT_NAME = $(notdir $(PWD))
NET = app
SVC = api
DB_SVC = db
DB_NAME = app_dev

## Container up
.PHONY: up
up: down
	docker compose up

## Container down
.PHONY: down
down:
	docker compose down

# Cotainer attach
.PHONY: attach
attach:
	docker exec -it $(SVC) sh

## Lint
.PHONY: lint
lint:
	docker compose run --rm $(SVC) sh bin/golangci-lint.sh

## Test
.PHONY: test
test:
	docker compose run --rm $(SVC) go test ./...

## Generate Mocks
.PHONY: gen-mock
gen-mock:
	docker compose run --rm ${SVC} sh bin/mockgen.sh

## Init ent schema # make ent-init flags="Example"
.PHONY: init-ent
init-ent:
	docker compose run --rm ${SVC} go run entgo.io/ent/cmd/ent init $(flags)

## Generate ent
.PHONY: gen-ent
gen-ent:
	docker compose run --rm ${SVC} go generate ./ent

## Migration
.PHONY: run-migrate
run-migrate:
	docker compose run --rm ${SVC} go run ./cmd/migrate

## Show ent Schema
.PHONY: show-schema
show-schema:
	docker compose run --rm ${SVC} go run entgo.io/ent/cmd/ent describe ./ent/schema

## Generate gqlgen
.PHONY: gen-gql
gen-gql:
	 docker compose run --rm ${SVC} go run github.com/99designs/gqlgen generate

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
