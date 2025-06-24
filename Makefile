SOURCES=$(shell find . \( -name '*.go' -not -name '*.gen.go' \) -or -name 'go.*' -or -name 'Makefile')
SCHEMA_SRC=$(shell find ./oas -name '*.yaml')
DISTDIR=dist
SCHEMA_DST=blueprint-schema.json

.PHONY: help
help: ## print this help
	@echo "make [TARGETS...]"
	@echo
	@echo 'Targets:'
	@awk 'match($$0, /^([a-zA-Z_\/-]+):.*?## (.*)$$/, m) {printf "  \033[36m%-30s\033[0m %s\n", m[1], m[2]}' $(MAKEFILE_LIST) | sort

$(DISTDIR):
	mkdir -p $(DISTDIR)

.PHONY: write-fixtures
write-fixtures: ## Write new test fixtures
	@rm -f ./testdata/*.out.yaml ./testdata/*.validator.out ./testdata/*.validator.out
	@WRITE_FIXTURES=1 go test -count=1 . ./pkg/...

.PHONY: pkg-go-dev-update
pkg-go-dev-update: ## Schedule https://pkg.go.dev/github.com/osbuild/blueprint-schema for update
	GOPROXY=https://proxy.golang.org go get github.com/osbuild/blueprint-schema

.PHONY: test
test: ## Run all tests
	@go test -count=1 . ./pkg/...

.PHONY: test-diff
test-diff: ## Diff all fields test data
	@go run github.com/homeport/dyff/cmd/dyff@latest between testdata/all-fields.in.yaml testdata/all-fields.out.yaml

.PHONY: image-builder-blueprint
image-builder-blueprint: $(SOURCES) $(SCHEMA_SRC) ## Build the image-builder-blueprint binary
	go build -o image-builder-blueprint ./cmd/image-builder-blueprint

SCHEMA_BUILD_CLI=go run ./cmd/image-builder-blueprint
# If you find yourself in a loop being unable to build the CLI, switch to the "main" branch
# and build the CLI command via "make image-builder-blueprint" and use it.
blueprint-oas3.yaml: $(SCHEMA_SRC) $(SOURCES)
	$(SCHEMA_BUILD_CLI) -print-yaml-schema > blueprint-oas3.yaml

blueprint-oas3.json: $(SCHEMA_SRC) $(SOURCES)
	$(SCHEMA_BUILD_CLI) -print-json-schema > blueprint-oas3.json

blueprint-oas3-ext.json: $(SCHEMA_SRC) $(SOURCES)
	$(SCHEMA_BUILD_CLI) -print-json-extended-schema > blueprint-oas3-ext.json

.PHONY: schema-bundle
schema-bundle: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json ## Bundle OpenAPI schema

pkg/ubp/types.gen.go: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json oapi-codegen.cfg.yml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -config oapi-codegen.cfg.yml -generate types -o pkg/ubp/types.gen.go blueprint-oas3.json

pkg/ubp/http.gen.go: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json oapi-codegen.cfg.yml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -config oapi-codegen.cfg.yml -generate std-http -o pkg/ubp/http.gen.go blueprint-oas3.json

schema: pkg/ubp/types.gen.go pkg/ubp/http.gen.go ## Generate bundled schema and Go code

.PHONY: yamllint
yamllint: ## Lint YAML files
	yamllint -c .yamllint.yaml oas/

cmd/wasm/wasm_exec.js:
# will not work until Go 1.24+ download instead
# @cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" cmd/wasm/wasm_exec.js
	@curl -sSL https://raw.githubusercontent.com/golang/go/master/lib/wasm/wasm_exec.js -o cmd/wasm/wasm_exec.js

cmd/wasm/blueprint.wasm: $(SCHEMA_SRC) $(SOURCES) cmd/wasm/wasm_exec.js cmd/wasm/index.html
	GOOS=js GOARCH=wasm go build -o cmd/wasm/blueprint.wasm ./cmd/wasm

.PHONY: wasm
wasm: cmd/wasm/blueprint.wasm cmd/wasm/wasm_exec.js ## Build the WASM binary

.PHONY: run-wasm
run-wasm: cmd/wasm/blueprint.wasm cmd/wasm/wasm_exec.js ## Run the WASM binary
	@echo "Open http://localhost:8080 and Ctrl+C to stop the server"
	@python3 -m http.server 8080 --directory cmd/wasm

.PHONY: golint
golint: ## Run golint on the codebase
	@golangci-lint run --config=.golangci.yaml ./pkg/...

.PHONY: golint-podman
golint-podman: ## Run golint via podman on the codebase
	@podman run -t --rm -v $(shell pwd):/app:ro -v ~/.cache:/root/.cache -w /app golangci/golangci-lint:latest golangci-lint run -v ./pkg/...

.PHONY: clean
clean: ## Clean up all build artifacts
	rm -rf $(DISTDIR) blueprint-oas3*.{yaml,json}
	rm -f ./testdata/*.out.{yaml,toml,json} ./testdata/*.validator.out
