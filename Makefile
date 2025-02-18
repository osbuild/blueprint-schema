.PHONY: help
help: ## print this help
	@echo "make [TARGETS...]"
	@echo
	@echo 'Targets:'
	@awk 'match($$0, /^([a-zA-Z_\/-]+):.*?## (.*)$$/, m) {printf "  \033[36m%-30s\033[0m %s\n", m[1], m[2]}' $(MAKEFILE_LIST) | sort

.PHONY: generate-schema
generate-schema: ## Generate schema
	@go run ./cmd/generate-schema > blueprint-schema.json
