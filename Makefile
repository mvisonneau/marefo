NAME          := marefo
VERSION       := $(shell git describe --tags --abbrev=1)
FILES         := $(shell git ls-files '*.go')
LDFLAGS       := -X 'main.version=$(VERSION)'
.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Install required libraries/tools
	go get -u -v github.com/golang/dep/cmd/dep
	go get -u -v golang.org/x/tools/cmd/goimports
	go get -u -v golang.org/x/tools/cmd/cover
	go get -u -v github.com/golang/lint/golint

certs: ## Generate self-signed certificates for TLS
	mkdir -p tls
	openssl req -new -newkey rsa:4096 -nodes -x509 \
    -subj "/C=UK/ST=Foo/L=Bar/O=Devel/CN=localhost" \
    -keyout tls/server.key -out tls/server.crt

.PHONY: fmt
fmt: ## Format source code
	goimports -w $(FILES)

.PHONY: lint
lint: ## Run golint and go vet against the codebase
	golint -set_exit_status . app config rand
	go vet ./...

.PHONY: test
test: ## Run the tests against the codebase
	go test -v ./...

.PHONY: install
install: ## Build and install locally the binary (dev purpose)
	go install .

.PHONY: build
build: ## Build the binary
	CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o $(NAME) main.go
	strip $(NAME)

.PHONY: deps
deps: ## Fetch all dependencies
	dep ensure -v

.PHONY: imports
imports: ## Fixes the syntax (linting) of the codebase
	goimports -d $(FILES)

.PHONY: clean
clean: ## Remove binary if it exists
	rm -f $(NAME)

.PHONY: coverage
coverage: ## Generates coverage report
	rm -rf *.out
	go test -coverprofile=coverage.out
	@for i in app config; do \
	 	go test -coverprofile=$$i.coverage.out github.com/mvisonneau/$(NAME)/$$i; \
		tail -n +2 $$i.coverage.out >> coverage.out; \
	done

.PHONY: all
all: lint imports test coverage build ## Test, builds and ship package for all supported platforms

.PHONY: help
help: ## Displays this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
