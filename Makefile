.PHONY: install build generate vet test docker-build docker-dev-build docs

### Version info
VERSION_GIT_REPO=$(shell git config --get remote.origin.url)
VERSION_GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
VERSION_GIT_HASH=$(shell git rev-parse HEAD)
VERSION_GIT_TAG=$(shell tag=$(git tag -l --contains HEAD);[ -n "$(tag)" ] && echo "$(tag)" || echo "v0.0.0")
BUILD_TAG=latest

install:
	go mod tidy && go mod vendor
	go install github.com/Pluto-tv/acyl

build:
	go build -mod tidy -mod vendor

generate:
	go generate ./...

vet:
	go vet $(shell go list ./... |grep pkg/)

test:
	go test -cover $(shell go list ./... |grep -v pkg/persistence  |grep -v pkg/locker |grep -v pkg/api)
	go test -cover github.com/Pluto-tv/acyl/pkg/persistence
	go test -cover github.com/Pluto-tv/acyl/pkg/locker
	go test -cover github.com/Pluto-tv/acyl/pkg/api

docker-build: install vet test
	DOCKER_BUILDKIT=1 docker buildx build \
	--platform linux/amd64 \
	--build-arg GIT_REPO="$(VERSION_GIT_REPO)" \
	--build-arg GIT_BRANCH="$(VERSION_GIT_BRANCH)" \
	--build-arg APP_VERSION="$(VERSION_GIT_TAG)" \
	--build-arg GIT_HASH="$(VERSION_GIT_HASH)" \
	-t acyl:$(BUILD_TAG) .

docker-dev-build:
	$(MAKE) BUILD_TAG=dev docker-build

docs:
	./openapi.sh
