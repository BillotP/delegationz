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
BUILD           := $(shell git rev-parse --short HEAD)
APP_NAME        := $(shell head -n 1 README.md | cut -d ' ' -f2 |  tr '[:upper:]' '[:lower:]')

DB_URL        ?= "postgres://postgres:supersecret@localhost:5432/dev"
REGISTRY_DEV	:= "ghcr.io/BillotP/delegationz"

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

# build: ## Build a CMD container image
# 	@export chartFname="deployments/charts/$(CMD)/Chart.yaml" && \
# 	export valueFname="deployments/charts/$(CMD)/values.yaml" && \
# 	export version=$$(sed -n 's/^appVersion:[[:space:]]*//p' "$$chartFname" | tr -d '"') && \
# 	export build=$$(git rev-parse --short HEAD) && \
# 	export registry=$(shell if [ $(STAGE) == "preprod" ]; then echo $(REGISTRY_PREPROD); else echo $(REGISTRY_DEV);fi) && \
# 	export imageName="$$registry/$(CMD)" && \
# 	export dockerTemplate=$$(cat $$valueFname | grep dockerTemplate: | grep -oP 'dockerTemplate: "\K[^"]+') && \
# 	export dockerfile=deployments/dockerfiles/$$dockerTemplate.dockerfile && \
# 	echo "$$chartFname version: $$version" && \
# 	$(DOCKER) build -t "$$imageName:$$version" \
# 	--build-arg cmd=$(CMD) \
# 	--build-arg version=$$version-$(BUILD) \
# 	-f $$dockerfile .

# push: ## Push a previously built CMD container image
# 	@export chartFname="deployments/charts/$(CMD)/Chart.yaml" && \
# 	export version=$$(sed -n 's/^appVersion:[[:space:]]*//p' "$$chartFname" | tr -d '"') && \
# 	export registry=$(shell if [ $(STAGE) == "preprod" ]; then echo $(REGISTRY_PREPROD); else echo $(REGISTRY_DEV);fi) && \
# 	export imageName="$$registry/$(CMD)" && \
# 	echo "Will push $$imageName:$$version" && \
# 	$(DOCKER) push "$$imageName:$$version"

# all: ## Build and push all CMDs container images
# 	@for d in cmd/**/* ; \
# 	do ! [[ $$d == *"__utils"* ]] && make build CMD=$${d#cmd/} || true ; \
# 	done
# 	@for d in cmd/**/* ; \
# 	do ! [[ $$d == *"__utils"* ]] && make push CMD=$${d#cmd/} || true ; \
# 	done

# bump_all: ## Increment all charts version to eg trigger a build / deploy in CI...
# 	@for f in deployments/charts/**/**/Chart.yaml; do  [[ $$f == *"utils"* ]] && continue ; echo "Processing $$f file..." && \
# 	 VERSION=$$(grep appVersion $$f | sed 's/appVersion[: "]*//' | sed 's/"//') && \
# 	 NVERSION=$$(./tools/increment_version.sh -p $$VERSION) \
# 	 app=$$(echo $$f | sed 's|deployments/charts/||g' | sed 's|/Chart.yaml||g') \
# 	 && \
# 	 echo "🩹 New patch from $$VERSION to $$NVERSION for $$app" && \
# 	 sed -i "s/$$VERSION/$$NVERSION/"  $$f; done

# bump: ## Increment CMD charts version to eg trigger a build / deploy in CI...
# 	@for f in deployments/charts/$(CMD)/**/Chart.yaml; do  echo "Processing $$f file..." && \
# 	 VERSION=$$(grep appVersion $$f | sed 's/appVersion[: "]*//' | sed 's/"//') && \
# 	 NVERSION=$$(./tools/increment_version.sh -p $$VERSION) \
# 	 app=$$(echo $$f | sed 's|deployments/charts/||g' | sed 's|/Chart.yaml||g') \
# 	 && \
# 	 echo "🩹 New patch from $$VERSION to $$NVERSION for $$app" && \
# 	 sed -i "s/$$VERSION/$$NVERSION/"  $$f; done


gen: ## Generate a fully typed golang orm from postgres db introspection cf pkg/_generate.go
	@sqlboiler psql -c utils/config/sqlboiler.toml -o "pkg/repository"

# updatesecrets: ## Update dotenv secrets files
# 	@for f in secrets/**/*.yaml ; do \
# 	                [[ $$f == *dev* ]] && echo "[INFO] no need to encrypt $$f secret" || \
# 	                echo "[INFO] Adding $$f" && git secret add $$f; \
# 	done
# 	git secret hide

# reveal:
# 	@git secret reveal

clean: ## Stop background cmd(s) and cleaning binary and temporary files 
	@rm -rf bin || true
	@rm -rf out || true
	@rm -rf logs || true
	@rm *.log **/*.log *.pid **/*.pid || true


.PHONY: help clean setup test all
