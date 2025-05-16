TINYGO?=tinygo
SOURCES=$(shell find . -name '*.go' -name 'go.mod' -name 'go.sum' -name 'Makefile')
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
	@WRITE_FIXTURES=1 go test -count=1 ./pkg/blueprint/

.PHONY: pkg-go-dev-update
pkg-go-dev-update: ## Schedule https://pkg.go.dev/github.com/osbuild/blueprint-schema for update
	GOPROXY=https://proxy.golang.org go get github.com/osbuild/blueprint-schema

.PHONY: test
test: ## Run all tests
	@go test -count=1 . ./pkg/blueprint

blueprint-oas3.yaml: $(SCHEMA_SRC) $(SOURCES)
	go run ./cmd/image-builder-blueprint -print-yaml-schema > blueprint-oas3.yaml

blueprint-oas3.json: $(SCHEMA_SRC) $(SOURCES)
	go run ./cmd/image-builder-blueprint -print-json-schema > blueprint-oas3.json

blueprint-oas3-ext.json: $(SCHEMA_SRC) $(SOURCES)
	go run ./cmd/image-builder-blueprint -print-json-extended-schema > blueprint-oas3-ext.json

.PHONY: schema-bundle
schema-bundle: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json ## Bundle OpenAPI schema

pkg/blueprint/types.gen.go: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json oapi-codegen.cfg.yml
	oapi-codegen -config oapi-codegen.cfg.yml -generate types -o pkg/blueprint/types.gen.go blueprint-oas3.json

pkg/blueprint/http.gen.go: blueprint-oas3.yaml blueprint-oas3.json blueprint-oas3-ext.json oapi-codegen.cfg.yml
	oapi-codegen -config oapi-codegen.cfg.yml -generate std-http -o pkg/blueprint/http.gen.go blueprint-oas3.json

schema: pkg/blueprint/types.gen.go pkg/blueprint/http.gen.go ## Generate bundled schema and Go code

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
	rm -rf $(DISTDIR) blueprint-oas3*.{yaml,json} pkg/blueprint/*.gen.go
	rm -f ./testdata/*.out.yaml ./testdata/*.validator.out ./testdata/*.validator.out
	rm -rf web/node_modules web/dist
	@echo -e '// allow go run\npackage blueprint\ntype Blueprint struct {}' > pkg/blueprint/types.gen.go
