VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

ORG := github.com
OWNER := inwinstack
REPOPATH ?= $(ORG)/$(OWNER)/kubeconfig-generator

$(shell mkdir -p ./out)

.PHONY: build
build: out/kg out/kgctl

.PHONY: kg
out/kg:
	go build -ldflags="-X $(REPOPATH)/pkg/version.version=$(VERSION)" -a -o $@ apps/server/main.go

.PHONY: kgctl
out/kgctl:
	go build -ldflags="-X $(REPOPATH)/pkg/version.version=$(VERSION)" -a -o $@ apps/cli/main.go

.PHONY: build_image
build_image:
	docker build -t $(OWNER)/kubeconfig-generator:$(VERSION) .

.PHONY: clean
clean:
	rm -rf out/
