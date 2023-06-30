###
### A Makefile to make things
###
### Note : use `make xxx DOCKER=podman SED=gsed` overrides
### 			style to match your available bins
###
SHELL           := /bin/bash
DOCKER			?= docker
STAGE        	?= dev
SED					?= sed
LOGS			?= logs
CMD				?= delegationz
BUILD           := $(shell git rev-parse --short HEAD)
APP_NAME        := $(shell head -n 1 README.md | cut -d ' ' -f2 |  tr '[:upper:]' '[:lower:]')

DB_URL        ?= "postgres://postgres:supersecret@localhost:5432/dev"
REGISTRY_DEV	:= "ghcr.io/billotp/delegationz"

help: ## Print this help message and exit
	@echo -e "\n\t\t$(APP_NAME)-$(BUILD) \033[1mmake\033[0m options:\n"
	@perl -nle 'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "- \033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo -e "\n"

default: help

start: ## Start a cmd in background
	@mkdir -p logs
	@echo "[INFO] Will start $(CMD) listening on $(PORT)"
	@FOO=$(CMD) ; \
	FOO=$${FOO##*/} ; \
	{ PORT=$(PORT) HEALTH_PORT=$(HEALTH_PORT) CMD=$(CMD) go run cmd/$(CMD)/main.go &>"$(LOGS)/$$FOO.log" & }
	@echo "[INFO] Started $(CMD)"

devmigrate: ## Run migrations
	@cd migrations && npx --yes prisma@latest migrate deploy

stopdb: ## Stop and delete test-postgres
	@$(DOCKER) stop test-postgres || true
	@$(DOCKER) rm test-postgres || true

testdb: ## Run a fresh migrated local database
	@make stopdb
	@$(DOCKER) run --name test-postgres -d -p 5432:5432 \
	--env POSTGRES_PASSWORD=supersecret \
	--env POSTGRES_DB=dev \
	docker.io/postgres:alpine
	@sleep 5 && make devmigrate
	@echo "[INFO] Postgres Test available at postgres://postgres:supersecret@localhost:5432/dev"

test: ## Run all unit tests
	@rm cover.out coverage.html || true
	@go test -coverprofile cover.out -v ./... || true
	@go tool cover -func cover.out
	@go tool cover -html=cover.out -o coverage.html

redis: ## Launch a local redis store
	@$(DOCKER) kill local-redis || true
	@$(DOCKER) rm local-redis || true
	@echo "[INFO] Will start local redis store"
	@$(DOCKER) run -d --name local-redis -p 6379:6379 docker.io/redis:alpine
	@echo "[INFO] Local redis instance listening on 'redis://localhost:6379'"

build: ## Build a CMD container image
	@export version=$$(git rev-parse --short HEAD) && \
	export registry=$(REGISTRY_DEV) && \
	export imageName="$$registry/$(CMD)" && \
	export dockerfile=$(shell if [ $(CMD) == "delegationz" ]; then echo "utils/dockerfiles/delegationz.dockerfile"; else echo "utils/dockerfiles/light.dockerfile";fi) && \
	echo "$$dockerfile" && \
	$(DOCKER) build -t "$$imageName:$$version" \
	--build-arg version=$$version \
	--build-arg cmd=$(CMD) \
	-f $$dockerfile .

push: ## Push a previously built CMD container image
	@export version=$(git rev-parse --short HEAD) && \
	export registry=$(shell if [ $(STAGE) == "preprod" ]; then echo $(REGISTRY_PREPROD); else echo $(REGISTRY_DEV);fi) && \
	export imageName="$$registry/$(CMD)" && \
	echo "Will push $$imageName:$$version" && \
	$(DOCKER) push "$$imageName:$$version"

gen: ## Generate a fully typed golang orm from postgres db introspection cf pkg/_generate.go
	@sqlboiler psql -c utils/config/sqlboiler.toml -o "pkg/repository"

clean: ## Stop background cmd(s) and cleaning binary and temporary files 
	@rm -rf bin || true
	@rm -rf out || true
	@rm -rf logs || true
	@rm *.log **/*.log *.pid **/*.pid || true


.PHONY: help clean setup test all build push gen
