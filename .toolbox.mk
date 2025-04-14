## toolbox - start
## Generated with https://github.com/bakito/toolbox

## Current working directory
TB_LOCALDIR ?= $(shell which cygpath > /dev/null 2>&1 && cygpath -m $$(pwd) || pwd)
## Location to install dependencies to
TB_LOCALBIN ?= $(TB_LOCALDIR)/bin
$(TB_LOCALBIN):
	mkdir -p $(TB_LOCALBIN)

## Tool Binaries
TB_GINKGO ?= $(TB_LOCALBIN)/ginkgo
TB_GO_BUILDER_GENERATOR ?= $(TB_LOCALBIN)/go-builder-generator
TB_GOLANGCI_LINT ?= $(TB_LOCALBIN)/golangci-lint
TB_GORELEASER ?= $(TB_LOCALBIN)/goreleaser

## Tool Versions
TB_GINKGO_VERSION ?= v2.23.4
TB_GO_BUILDER_GENERATOR_VERSION ?= v1.9.3
TB_GOLANGCI_LINT_VERSION ?= v2.1.1
TB_GORELEASER_VERSION ?= v2.8.2

## Tool Installer
.PHONY: tb.ginkgo
tb.ginkgo: $(TB_GINKGO) ## Download ginkgo locally if necessary.
$(TB_GINKGO): $(TB_LOCALBIN)
	test -s $(TB_LOCALBIN)/ginkgo || GOBIN=$(TB_LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@$(TB_GINKGO_VERSION)
.PHONY: tb.go-builder-generator
tb.go-builder-generator: $(TB_GO_BUILDER_GENERATOR) ## Download go-builder-generator locally if necessary.
$(TB_GO_BUILDER_GENERATOR): $(TB_LOCALBIN)
	test -s $(TB_LOCALBIN)/go-builder-generator || GOBIN=$(TB_LOCALBIN) go install github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@$(TB_GO_BUILDER_GENERATOR_VERSION)
.PHONY: tb.golangci-lint
tb.golangci-lint: $(TB_GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(TB_GOLANGCI_LINT): $(TB_LOCALBIN)
	test -s $(TB_LOCALBIN)/golangci-lint || GOBIN=$(TB_LOCALBIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(TB_GOLANGCI_LINT_VERSION)
.PHONY: tb.goreleaser
tb.goreleaser: $(TB_GORELEASER) ## Download goreleaser locally if necessary.
$(TB_GORELEASER): $(TB_LOCALBIN)
	test -s $(TB_LOCALBIN)/goreleaser || GOBIN=$(TB_LOCALBIN) go install github.com/goreleaser/goreleaser@$(TB_GORELEASER_VERSION)

## Reset Tools
.PHONY: tb.reset
tb.reset:
	@rm -f \
		$(TB_LOCALBIN)/ginkgo \
		$(TB_LOCALBIN)/go-builder-generator \
		$(TB_LOCALBIN)/golangci-lint \
		$(TB_LOCALBIN)/goreleaser

## Update Tools
.PHONY: tb.update
tb.update: tb.reset
	toolbox makefile -f $(TB_LOCALDIR)/Makefile \
		github.com/onsi/ginkgo/v2/ginkgo \
		github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator \
		github.com/golangci/golangci-lint/v2/cmd/golangci-lint \
		github.com/goreleaser/goreleaser
## toolbox - end