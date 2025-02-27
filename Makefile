# TinyGo requires a compatible Go version, typically older than the current one.
WASM_GOROOT?=$(HOME)/sdk/go1.21.13

SOURCES=$(shell find . -name '*.go' -name 'blueprint-schema.json' -name 'go.mod' -name 'go.sum' -name 'Makefile')
DISTDIR=dist

.PHONY: help
help: ## print this help
	@echo "make [TARGETS...]"
	@echo
	@echo 'Targets:'
	@awk 'match($$0, /^([a-zA-Z_\/-]+):.*?## (.*)$$/, m) {printf "  \033[36m%-30s\033[0m %s\n", m[1], m[2]}' $(MAKEFILE_LIST) | sort

$(DISTDIR):
	mkdir -p $(DISTDIR)

.PHONY: generate-schema
generate-schema: ## Generate schema
	@go run ./cmd/generate-schema > blueprint-schema.json

.PHONY: write-fixtures
write-fixtures: ## Write new test fixtures
	@rm -f ./fixtures/*.out.yaml ./fixtures/*.valid.out
	@WRITE_FIXTURES=1 go test -count=1 .

.PHONY: pkg-go-dev-update
pkg-go-dev-update: ## Schedule https://pkg.go.dev/github.com/osbuild/blueprint-schema for update
	GOPROXY=https://proxy.golang.org go get github.com/osbuild/blueprint-schema

.PHONY: test
test: ## Run all tests
	@go test -count=1 .

PLATFORMS:=$(DISTDIR)/blueconv_linux_amd64 \
	$(DISTDIR)/blueconv_linux_arm64 \
	$(DISTDIR)/blueconv_windows_amd64 \
	$(DISTDIR)/blueconv_windows_arm64 \
	$(DISTDIR)/blueconv_darwin_amd64 \
	$(DISTDIR)/blueconv_darwin_arm64

temp = $(subst _, ,$@)
os = $(word 2, $(temp))
arch = $(word 3, $(temp))

.PHONY: build
build: build-cli build-wasm ## Builds all binaries

.PHONY: build-cli
build-cli: $(PLATFORMS) ## Builds cli binaries

.PHONY: build-wasm
build-wasm: $(DISTDIR)/blueprint_go.wasm $(DISTDIR)/blueprint_tgo.wasm ## Builds wasm binaries

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o $(DISTDIR)/blueconv_$(os)_$(arch) ./cmd/blueconv/

.PHONY: $(DISTDIR)/blueprint_go.wasm
$(DISTDIR)/blueprint_go.wasm: $(DISTDIR) ## Builds wasm via go
	GOOS=js GOARCH=wasm go build -o $(DISTDIR)/blueprint_go.wasm ./cmd/blueconv/

$(DISTDIR)/blueprint_tgo.wasm: $(SOURCES) $(DISTDIR) ## Builds wasm via tinygo - GOROOT and GOPATH must be set to compatible Go
	GOROOT=$(WASM_GOROOT) PATH="$(WASM_GOROOT)/bin:$(PATH)" GOOS=js GOARCH=wasm tinygo build -scheduler=none -o $(DISTDIR)/blueprint_tgo.wasm ./cmd/blueconv/

.PHONY: run-web-editor-json
run-web-editor-json: ## show a demo-web editor for the json format
	xdg-open ./autocomplete-example-json.html

# Just set this in your environment or call directly:
# make WEB_EDITOR_HOST=hostname run-web-editor-yaml
export WEB_EDITOR_HOST?=0.0.0.0

.PHONY: run-web-editor-yaml
run-web-editor-yaml: ## Show a demo-web editor for the yaml format
	cd autocomplete-example-yaml && npm clean-install
	cd autocomplete-example-yaml && npm run start

.PHONY: clean
clean: ## Clean up all build artifacts
	rm -rf $(DISTDIR)
	rm -rf autocomplete-example-yaml/node_modules
	rm -rf autocomplete-example-yaml/dist
