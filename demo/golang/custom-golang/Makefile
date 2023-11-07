HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output

MODULE:=custom-golang
SRC_FILES:=$(shell find . -type f -name '*.go')
PLATFORM_ALL:=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 linux/arm/v7 windows/amd64

GIT_TAG:=$(shell git tag --contains HEAD)
GIT_REV:=git-$(shell git rev-parse --short HEAD)
VERSION:=$(if $(GIT_TAG),$(GIT_TAG),$(GIT_REV))

GO       = go
GO_MOD   = $(GO) mod
GO_ENV   = env CGO_ENABLED=0
GO_TEST_FLAGS:=-race -short -covermode=atomic -coverprofile=coverage.txt
GO_TEST_PKGS:=$(shell go list ./...)
GO_BUILD = $(GO_ENV) $(GO) build $(GO_FLAGS)
ifndef PLATFORMS
	GO_OS:=$(shell go env GOOS)
	GO_ARCH:=$(shell go env GOARCH)
	GO_ARM:=$(shell go env GOARM)
	PLATFORMS:=$(if $(GO_ARM),$(GO_OS)/$(GO_ARCH)/$(GO_ARM),$(GO_OS)/$(GO_ARCH))
	ifeq ($(GO_OS),darwin)
		PLATFORMS+=linux/amd64
	endif
else ifeq ($(PLATFORMS),all)
	override PLATFORMS:=$(PLATFORM_ALL)
endif

REGISTRY:=
XFLAGS:=--load
XPLATFORMS:=$(shell echo $(filter-out darwin/amd64 windows/amd64,$(PLATFORMS)) | sed 's: :,:g')

OUTPUT     :=output
OUTPUT_DIRS:=$(PLATFORMS:%=$(OUTPUT)/%/$(MODULE))
OUTPUT_BINS:=$(OUTPUT_DIRS:%=%/$(MODULE))
PKG_PLATFORMS := $(shell echo $(PLATFORMS) | sed 's:/:-:g')
OUTPUT_PKGS:=$(PKG_PLATFORMS:%=$(OUTPUT)/$(MODULE)_%_$(VERSION).zip)

.PHONY: all
all: build

prepare: prepare-dep
prepare-dep:
	git config --global http.sslVerify false
	$(AGILE_BCLOUD)

set-env:
	$(GO) env -w GO111MODULE="on"
	$(GO) env -w GONOSUMDB=\*
	$(GO) env -w GOPROXY=https://goproxy.cn

$(OUTPUT_BINS): set-env $(SRC_FILES)
	@echo "BUILD $@"
	@mkdir -p $(dir $@)
	@$(shell echo $(@:$(OUTPUT)/%/$(MODULE)/$(MODULE)=%)  | sed 's:/v:/:g' | awk -F '/' '{print "GOOS="$$1" GOARCH="$$2" GOARM="$$3""}') $(GO_BUILD) -o $@ .

.PHONY: build
build: $(OUTPUT_BINS)

.PHONY: build-local
build-local: set-env $(SRC_FILES)
	@echo "BUILD $(MODULE)"
	$(GO_MOD) tidy
	$(GO_BUILD) -o $(MODULE) .
	@chmod +x $(MODULE)

.PHONY: test
test: fmt
	@go test ${GO_TEST_FLAGS} ${GO_TEST_PKGS}
	@go tool cover -func=coverage.txt | grep total

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(OUTDIR)
	rm -rf $(HOMEDIR)/$(MODULE)
