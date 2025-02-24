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

# Just set this in your environment or call directly:
# make YAML_DEMO_BIND_IP=your_ip run-web-editor-yaml
export YAML_DEMO_BIND_IP?=0.0.0.0

.PHONY: run-web-editor-yaml
run-web-editor-yaml: ## Show a demo-web editor for the yaml format
	cd autocomplete-example-yaml && npm clean-install
	cd autocomplete-example-yaml && npm run start

.PHONY: clean
clean: ## Clean up all build artifacts
	rm -rf autocomplete-example-yaml/node_modules
	rm -rf autocomplete-example-yaml/dist
