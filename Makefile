TINYGO?=tinygo
SOURCES=$(shell find . -name '*.go' -name 'blueprint-schema.json' -name 'go.mod' -name 'go.sum' -name 'Makefile')
SCHEMA_SRC=$(shell find ./schema -name '*schema.yaml')
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

.PHONY: generate-schema
generate-schema: ## Generate schema
	@go run ./cmd/generate-schema > blueprint-schema.json

.PHONY: write-fixtures
write-fixtures: ## Write new test fixtures
	@rm -f ./fixtures/*.out.yaml ./fixtures/*.validator.json
	@WRITE_FIXTURES=1 go test -count=1 ./pkg/blueprint/

.PHONY: pkg-go-dev-update
pkg-go-dev-update: ## Schedule https://pkg.go.dev/github.com/osbuild/blueprint-schema for update
	GOPROXY=https://proxy.golang.org go get github.com/osbuild/blueprint-schema

.PHONY: test
test: ## Run all tests
	@go test -count=1 ./...

# Option --without-id is a workaround for VSCode: https://github.com/sourcemeta/jsonschema/blob/main/docs/bundle.markdown
$(SCHEMA_DST): $(SCHEMA_SRC) Makefile ## Build the schema from schema/*.schema.yaml files
	jsonschema bundle schema/blueprint.schema.yaml --verbose --resolve schema/ --extension schema.yaml --without-id > $@

.PHONY: schema-check
schema-check: ## Check input schema files against JSON Metaschema
	jsonschema metaschema -e schema.yaml schema/

.PHONY: schema-fmt
schema-fmt: $(SCHEMA_DST) ## Lint and format the bundled schema
	jsonschema lint --fix $(SCHEMA_DST)
	jsonschema fmt $(SCHEMA_DST)

.PHONY: schema-meta
schema-meta: $(SCHEMA_DST) ## Validate the schema against JSON Metaschema
	jsonschema metaschema $(SCHEMA_DST)

.PHONY: schema
schema: $(SCHEMA_DST) schema-meta schema-fmt ## Build, format, lint and validate the JSON schema

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
	@(type "wasm-objdump" &> /dev/null && wasm-objdump -j Export -x dist/blueprint_go.wasm dist/blueprint_tgo.wasm) || true

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o $(DISTDIR)/blueconv_$(os)_$(arch) ./cmd/blueconv/

$(DISTDIR)/blueprint_go.wasm: $(SOURCES) $(DISTDIR) ## Builds wasm via go
	GOOS=js GOARCH=wasm go build -o $(DISTDIR)/blueprint_go.wasm ./cmd/wasm/

$(DISTDIR)/blueprint_tgo.wasm: $(SOURCES) $(DISTDIR) ## Builds wasm via tinygo - GOROOT and GOPATH must be set to compatible Go
	GOOS=js GOARCH=wasm $(TINYGO) build -scheduler=none -o $(DISTDIR)/blueprint_tgo.wasm ./conv/wasm/

.PHONY: run-web-editor-json
run-web-editor-json: ## show a demo-web editor for the json format
	xdg-open ./web/src/json.html

# Just set this in your environment or call directly:
# make WEB_EDITOR_HOST=hostname run-editor
export WEB_EDITOR_HOST?=localhost

.PHONY: run-editor
run-editor: ## Build, run webserver and open a demo-web editor
	cd web && npm clean-install
	cd web && npm run start

.PHONY: build-editor
build-editor: ## Build the demo-web editor
	cd web && npm clean-install
	cd web && npm run build

.PHONY: clean
clean: ## Clean up all build artifacts
	rm -rf $(DISTDIR)
	rm -rf web/node_modules
	rm -rf web/dist
