PROJECT_NAME:=$(shell basename `pwd`)

include .env
export

BUILD_IMAGE:=$(PROJECT_NAME)-go

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

VARIABLES := PROJECT_NAME=$(PROJECT_NAME)
COMPOSE:=$(VARIABLES) && docker-compose -f docker-compose.yml

DOCKER_OPS=--rm -u $(CURRENT_UID):$(CURRENT_GID) --env-file .env

ROOT=`pwd`

.PHONY: build
build:
	docker build -t ${BUILD_IMAGE} .

up-db: 
	${COMPOSE} up --build -d postgres

up: 
	${COMPOSE} up --build -d postgres app

down: 
	${COMPOSE} down

migrations-create:
	goose -dir=migrations postgres ${POSTGRES_URL} create $(name) sql

migrations-up:
	goose -dir=migrations -allow-missing postgres ${POSTGRES_URL} up

migrations-down:
	goose -dir=migrations postgres ${POSTGRES_URL} down

migrations-reset:
	goose -dir=migrations postgres ${POSTGRES_URL} reset

.PHONY: code-analyze
code-analyze:
	docker run --rm -v ${ROOT}:/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v cmd/... internal/... pkg/...

.PHONY: code-analyze-fix
code-analyze-fix:
	docker run --rm -v ${ROOT}:/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run --fix -v cmd/... internal/... pkg/...

.PHONY: code-style
code-style:
	gofumpt -l -extra internal/ cmd/ pkg/

.PHONY: code-style-fix
code-style-fix:
	gofumpt -w -extra internal/ cmd/ pkg/

.PHONY: deps-security-audit
deps-security-audit:
	go list -json -deps ./... | docker run --rm -i sonatypecommunity/nancy:latest sleuth

generate:
	go generate ./...

test:
	go test -cover ./internal/... ./pkg/...
