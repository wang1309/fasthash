ROOT_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOBIN := go
BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD)
LAST_COMMIT_ID=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date)
MO_VERSION=$(shell git describe --always --tags $(shell git rev-list --tags --max-count=1))
GO_VERSION=$(shell go version)
CGO_DEBUG_OPT :=
CGO_OPTS=CGO_CFLAGS="-I$(ROOT_DIR)/cgo" CGO_LDFLAGS="-L$(ROOT_DIR)/cgo -lmo -lm"
GO=$(CGO_OPTS) $(GOBIN)
GOLDFLAGS=-ldflags="-X 'main.GoVersion=$(GO_VERSION)' -X 'main.BranchName=$(BRANCH_NAME)' -X 'main.CommitID=$(LAST_COMMIT_ID)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.Version=$(MO_VERSION)'"

.PHONY: cgo
cgo:
	@(cd cgo; make ${CGO_DEBUG_OPT})

.PHONY: clean
clean:
	$(info [Clean up])
	$(info Clean go test cache)
	@go clean -testcache
	$(MAKE) -C cgo clean

.PHONY: clean
test:
	$(GO) test -v