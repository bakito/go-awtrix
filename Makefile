# Run go golanci-lint
lint: generate golangci-lint
	$(GOLANGCI_LINT) run --fix

# Run go mod tidy
tidy:
	go mod tidy

# Run tests
test: ginkgo tidy lint
	$(GINKGO) -r --cover --coverprofile=coverage.out

generate: go-builder-generator
	$(GO_BUILDER_GENERATOR) generate --file client/types/app.go --structs App --dest client/types/
	$(GO_BUILDER_GENERATOR) generate --file client/types/settings.go --structs Settings --dest client/types/

## toolbox - start
## Current working directory
LOCALDIR ?= $(shell which cygpath > /dev/null 2>&1 && cygpath -m $$(pwd) || pwd)
## Location to install dependencies to
LOCALBIN ?= $(LOCALDIR)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
GINKGO ?= $(LOCALBIN)/ginkgo
GO_BUILDER_GENERATOR ?= $(LOCALBIN)/go-builder-generator
GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint
GORELEASER ?= $(LOCALBIN)/goreleaser

## Tool Versions
GINKGO_VERSION ?= v2.17.2
GO_BUILDER_GENERATOR_VERSION ?= v1.3.1
GOLANGCI_LINT_VERSION ?= v1.58.0
GORELEASER_VERSION ?= v1.25.1

## Tool Installer
.PHONY: ginkgo
ginkgo: $(GINKGO) ## Download ginkgo locally if necessary.
$(GINKGO): $(LOCALBIN)
	test -s $(LOCALBIN)/ginkgo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)
.PHONY: go-builder-generator
go-builder-generator: $(GO_BUILDER_GENERATOR) ## Download go-builder-generator locally if necessary.
$(GO_BUILDER_GENERATOR): $(LOCALBIN)
	test -s $(LOCALBIN)/go-builder-generator || GOBIN=$(LOCALBIN) go install github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@$(GO_BUILDER_GENERATOR_VERSION)
.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(GOLANGCI_LINT): $(LOCALBIN)
	test -s $(LOCALBIN)/golangci-lint || GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
.PHONY: goreleaser
goreleaser: $(GORELEASER) ## Download goreleaser locally if necessary.
$(GORELEASER): $(LOCALBIN)
	test -s $(LOCALBIN)/goreleaser || GOBIN=$(LOCALBIN) go install github.com/goreleaser/goreleaser@$(GORELEASER_VERSION)

## Update Tools
.PHONY: update-toolbox-tools
update-toolbox-tools:
	@rm -f \
		$(LOCALBIN)/ginkgo \
		$(LOCALBIN)/go-builder-generator \
		$(LOCALBIN)/golangci-lint \
		$(LOCALBIN)/goreleaser
	toolbox makefile -f $(LOCALDIR)/Makefile \
		github.com/onsi/ginkgo/v2/ginkgo \
		github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/goreleaser/goreleaser
## toolbox - end
