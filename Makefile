GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

PROJECT_ROOT := github.com/suutaku/protoc-gen-gorm

DOCKERFILE_PATH := $(CURDIR)/docker
IMAGE_REGISTRY ?= infoblox
IMAGE_VERSION  ?= dev-gengorm

OS         := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH       := $(shell uname -m )
OSOPER     := $(shell uname -s | tr '[:upper:]' '[:lower:]' | sed 's/darwin/apple-darwin/' | sed 's/linux/linux-gnu/')
ARCHOPER   := $(shell uname -m )
PROTOC_VER := 3.13.0

export PATH := $(shell pwd)/bin:$(PATH)

bin/protoc-${PROTOC_VER}.zip:
	mkdir -p bin
	curl -L -o bin/protoc-${PROTOC_VER}.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-${OS}-${ARCH}.zip

bin/protoc: bin/protoc-${PROTOC_VER}.zip
	unzip -o -d bin $^
	mv bin/bin/protoc bin/protoc-${PROTOC_VER}
	chmod +x bin/protoc-${PROTOC_VER}
	ln -sf protoc-${PROTOC_VER} $@
	touch $@

# configuration for the protobuf gentool
SRCROOT_ON_HOST      := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER := /go/src/$(PROJECT_ROOT)
DOCKERPATH           := /go/src
DOCKER_RUNNER        := docker run --rm
DOCKER_RUNNER        += -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER) -w $(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR     := infoblox/docker-protobuf:latest
PROTOC_FLAGS         := -I. -Ivendor \
		-Iproto \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/v2 \
		--go_out="Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options:$(shell go env GOPATH)/src" \
		--gorm_out="engine=postgres,enums=string,gateway,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options:$(shell go env GOPATH)/src"

GENTOOL_FLAGS         := -Ivendor -Iexample -Iproto \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/v2 \
		--gorm_out="Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options,engine=postgres,enums=string,gateway:/go" \
		--go_out="Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options:/go"
GENERATOR            := $(DOCKER_RUNNER) $(DOCKER_GENERATOR) $(PROTOC_FLAGS)

.PHONY: default
default: vendor install

.PHONY: vendor
vendor:
	@dep ensure -vendor-only

.PHONY: vendor-update
vendor-update:
	@dep ensure

protos: options/gorm.pb.go types/types.proto \
	example/postgres_arrays/postgres_arrays.pb.go \
	example/user/user.pb.go \
	example/feature_demo/demo_types.pb.go \
	example/feature_demo/demo_service.pb.go \
	example/feature_demo/demo_multi_file.pb.go \
	example/feature_demo/demo_multi_file_service.pb.go

# FIXME: match these with patterns

options/gorm.pb.go: options/gorm.proto
	protoc $(PROTOC_FLAGS) $^

types/types.pb.go: types/types.proto
	protoc $(PROTOC_FLAGS) $^

example/user/user.pb.go: example/user/user.proto
	protoc  $(PROTOC_FLAGS) $^

example/postgres_arrays/postgres_arrays.pb.go: example/postgres_arrays/postgres_arrays.proto
	protoc $(PROTOC_FLAGS) $^

example/feature_demo/demo_multi_file_service.pb.go: example/feature_demo/demo_multi_file_service.proto
	protoc $(PROTOC_FLAGS) $^

example/feature_demo/demo_multi_file.pb.go: example/feature_demo/demo_multi_file.proto
	protoc $(PROTOC_FLAGS) $^

example/feature_demo/demo_service.pb.go: example/feature_demo/demo_service.proto
	protoc $(PROTOC_FLAGS) $^

example/feature_demo/demo_types.pb.go: example/feature_demo/demo_types.proto
	protoc $(PROTOC_FLAGS) $^

build: bin/protoc-gen-gorm

test: bin/protoc protos
	go test -v ./...
	go build ./example/user
	go build ./example/feature_demo

.PHONY: bin/protoc-gen-gorm
bin/protoc-gen-gorm: $(shell find plugin/)
	go build -o bin/protoc-gen-gorm

.PHONY: install
install:
	go install

options: options-proto
	go build ./options

.PHONY: example
example:

	protoc -I. $(PROTOC_FLAGS) \
		example/feature_demo/demo_types.proto \
		example/feature_demo/demo_multi_file_service.proto \
		example/feature_demo/demo_multi_file.proto \
		example/feature_demo/demo_types.proto \
		example/feature_demo/demo_service.proto \
		example/feature_demo/demo_multi_file_service.proto

	protoc -I. -I$(SRCPATH) -I./vendor -I./vendor -I./vendor/github.com/grpc-ecosystem/grpc-gateway \
		--go_out="plugins=grpc:$(SRCPATH)" --gorm_out="$(SRCPATH)" \
		example/user/user.proto

.PHONY: test
test: example run-tests

.PHONY: gentool
gentool: vendor
	@docker build -f $(GENGORM_DOCKERFILE) -t $(GENGORM_IMAGE):$(IMAGE_VERSION) .
	@docker tag $(GENGORM_IMAGE):$(IMAGE_VERSION) $(GENGORM_IMAGE):latest
	@docker image prune -f --filter label=stage=server-intermediate

.PHONY: gentool-example
gentool-example: gentool
	@$(GENERATOR) \
		--go_out="plugins=grpc:$(DOCKERPATH)" \
		--gorm_out="engine=postgres,enums=string,gateway:$(DOCKERPATH)" \
			example/feature_demo/demo_multi_file.proto \
			example/feature_demo/demo_types.proto \
			example/feature_demo/demo_service.proto \
			example/feature_demo/demo_multi_file_service.proto

.PHONY: gentool-test
gentool-test: gentool-example run-tests

.PHONY: gentool-types
gentool-types:
	@$(GENERATOR) --go_out=$(DOCKERPATH) types/types.proto

.PHONY: gentool-options
gentool-options:
	@$(GENERATOR) \
                --gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:$(DOCKERPATH)" \
                options/gorm.proto
