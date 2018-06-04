VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

GOOS ?= $(shell go env GOOS)

ORG := github.com
OWNER := inwinstack
REPOPATH ?= $(ORG)/$(OWNER)/kubeconfig-generator

$(shell mkdir -p ./out)

.PHONY: build
build: out/kg out/kgctl

.PHONY: kg
out/kg:
	GOOS=$(GOOS) go build \
	  -ldflags="-X $(REPOPATH)/pkg/version.version=$(VERSION)" -a -o $@ apps/server/main.go

.PHONY: kgctl
out/kgctl:
	GOOS=$(GOOS) go build \
	  -ldflags="-X $(REPOPATH)/pkg/version.version=$(VERSION)" -a -o $@ apps/cli/main.go

.PHONY: build_images
build_images: image-kg-server image-kg-ui

.PHONY: image-kg-server
image-kg-server:
	docker build -t $(OWNER)/kg-server:$(VERSION) ./apps/server/

.PHONY: image-kg-ui
image-kg-ui:
	docker build -t $(OWNER)/kg-ui:$(VERSION) ./apps/ui/

.PHONY: clean
clean:
	rm -rf out/
