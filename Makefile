.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.PHONY: create-certs
create-certs: ## Create containers by docker-compose.
	cd gondola/certificates && go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

.PHONY: docker-compose-build
docker-compose-build: ## Build containers by docker-compose.
	docker-compose -f docker-compose-local.yml build

.PHONY: docker-compose-up
docker-compose-up: ## Run containers by docker-compose.
	docker-compose -f docker-compose-local.yml up

.PHONY: docker-compose-up-d
docker-compose-up-d: ## Run containers in the background by docker-compose.
	docker-compose -f docker-compose-local.yml up -d

.PHONY: docker-compose-pull
docker-compose-pull: ## Pull images by docker-compose.
	docker-compose -f docker-compose-local.yml pull

.PHONY: setup-buildx
setup-buildx: ## Set up buildx builder.
	docker buildx create --name buildx-builder
	docker buildx use buildx-builder

.PHONY: build-and-push
build-and-push: ## Build and push image to dockerhub.
	docker buildx build --no-cache --push --platform linux/amd64,linux/arm64 --file app/Dockerfile --tag bmfsan/bmf-tech-client app/

.PHONY: mod
mod: ## Run go mod download.
	cd app && go mod download

.PHONY: install-staticcheck
install-staticcheck: ## Install staticcheck.
ifeq ($(shell command -v staticcheck 2> /dev/null),)
	cd app && go install honnef.co/go/tools/cmd/staticcheck@latest
endif

.PHONY: staticcheck
staticcheck: ## Run staticcheck.
	cd app && staticcheck ./...

.PHONY: gofmt
gofmt: ## Run gofmt.
	cd app && test -z "$(gofmt -s -l . | tee /dev/stderr)"

.PHONY: vet
vet: ## Run vet.
	cd app && go vet -v ./...

.PHONY: test
test: ## Run unit tests.
	cd app && go test -v -race ./...

.PHONY: test-cover
test-cover: ## Run unit tests with cover options. ex. make test-cover OUT="c.out"
	cd app && go test -v -race -cover -coverprofile=$(OUT) -covermode=atomic ./...

.PHONY: build
build: ## Run go build
	cd app && go build