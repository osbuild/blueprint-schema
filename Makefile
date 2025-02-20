.PHONY: help
help: ## print this help
	@echo "make [TARGETS...]"
	@echo
	@echo 'Targets:'
	@awk 'match($$0, /^([a-zA-Z_\/-]+):.*?## (.*)$$/, m) {printf "  \033[36m%-30s\033[0m %s\n", m[1], m[2]}' $(MAKEFILE_LIST) | sort

.PHONY: generate-schema
generate-schema: ## Generate schema
	@go run ./cmd/generate-schema > blueprint-schema.json

.PHONY: write-fixtures
write-fixtures: ## Write new test fixtures
	@rm -f ./validate/fixtures/*.out.yaml ./validate/fixtures/*.valid.json
	@WRITE_FIXTURES=1 go test ./validate

.PHONY: test
test: ## Run all tests
	@go test .
	@go test ./validate

.PHONY: run-web-editor-json
run-web-editor-json: ## show a demo-web editor for the json format
	xdg-open ./autocomplete-example-json.html
