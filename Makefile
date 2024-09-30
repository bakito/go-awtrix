# Include toolbox tasks
include ./.toolbox.mk

# Run go golanci-lint
lint: generate tb.golangci-lint
	$(TB_GOLANGCI_LINT) run --fix

# Run go mod tidy
tidy:
	go mod tidy

# Run tests
test: tb.ginkgo tidy lint
	$(TB_GINKGO) -r --cover --coverprofile=coverage.out

generate: tb.go-builder-generator
	$(TB_GO_BUILDER_GENERATOR) generate --file client/types/app.go --structs App --dest client/types/
	$(TB_GO_BUILDER_GENERATOR) generate --file client/types/settings.go --structs Settings --dest client/types/

