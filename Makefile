OS ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)

# Image URL to use all building/pushing image targets
REGISTRY ?= registry.gitlab.com/whizus/customer/pinto
IMAGE ?= cert-manager-webhook-pinto
FULL_IMAGE ?= $(REGISTRY)/$(IMAGE)

IMAGE_TAG ?= $(shell git rev-parse --short HEAD)

# Openapi generator
PACKAGE ?= "gopinto"
SOURCE_URL ?= "https://pinto.irgendwo.co/api/swagger/dns-swagger.json"

GIT_HOST ?= "gitlab.com"
GIT_REPOSITORY_NAMESPACE ?= "whizus/customer/pinto"
GIT_REPOSITORY_NAME ?= "cert-manager-webhook-pinto"

# Kubebuilder
KUBEBUILDER_VERSION=2.3.1

TEST_ZONE_NAME ?= example.com.

# Run tests
test: unit-test tests/kubebuilder
	TEST_ZONE_NAME=$(TEST_ZONE_NAME) go test -v ./... -coverprofile cover.out

.PHONY: unit-test
unit-test:
	go test -v -race ./pkg/...

tests/kubebuilder:
	curl -fsSL https://github.com/kubernetes-sigs/kubebuilder/releases/download/v$(KUBEBUILDER_VERSION)/kubebuilder_$(KUBEBUILDER_VERSION)_$(OS)_$(ARCH).tar.gz -o kubebuilder-tools.tar.gz
	mkdir tests/kubebuilder
	tar -xvf kubebuilder-tools.tar.gz
	mv kubebuilder_$(KUBEBUILDER_VERSION)_$(OS)_$(ARCH)/bin tests/kubebuilder/
	rm kubebuilder-tools.tar.gz
	rm -R kubebuilder_$(KUBEBUILDER_VERSION)_$(OS)_$(ARCH)

clean-kubebuilder:
	rm -Rf tests/kubebuilder

build:
	CGO_ENABLED=0 go build -v -o target/cert-manager-webhook-pinto main.go

docker-build:
	docker build . --platform=$(OS)/$(ARCH) -t $(FULL_IMAGE):$(IMAGE_TAG)-$(ARCH)

.PHONY: get-dependencies
get-dependencies:
	go get -v -t -d ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt-fix
fmt-fix:
	gofmt -s -w .

.PHONY: generate
generate: generate-openapi fmt-fix

.PHONY: generate-openapi
generate-openapi:
	openapi-generator-cli generate -g go \
		-i $(SOURCE_URL) \
		-o internal/gopinto/ \
		--package-name $(PACKAGE) \
		--git-repo-id $(GIT_REPOSITORY_NAME) \
		--git-user-id $(GIT_REPOSITORY_NAMESPACE) \
		--git-host $(GIT_HOST) \
		--additional-properties=generateInterfaces=true,isGoSubmodule=true
	rm internal/gopinto/go.mod internal/gopinto/go.sum
