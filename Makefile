help: ## Show this help
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | sed 's/Makefile://' | awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-z0-9A-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 }'


.PHONY: install-tools
install-tools: ## Install dev tools locally using 'go install'. It's expected that GOPATH is in the PATH
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI {} go install {}


.PHONY: dev
dev:
	export DASH0_HOMEEXERCISE_ENV=dev && \
	go run ./...
