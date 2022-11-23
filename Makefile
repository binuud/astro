# Maintainer Binu Udayakumar <binu@portworx.com>

# Optional environmental variables
REPO ?= dronasys-com
APP_NAME ?= astro-stack
AWS_ECR ?= 919350043131.dkr.ecr.ap-southeast-1.amazonaws.com

BUILD_VER ?= a1.0.0

DOCKER_HUB_TAG ?= a1.0.0

TAGGED_NAME = $(REPO)/$(APP_NAME)

BASE_DIR	:=	$(shell git rev-parse --show-toplevel)
BIN			:=	$(BASE_DIR)/bin


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help vendor vendor-update
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

ifndef PROTOC
PROTOC = protoc
endif

ifndef PROTOWEB
PROTOWEB = protoc-gen-grpc-web
endif

protos-web-js: ## WebJs Proto - working for streams too
	$(PROTOC) -I=./proto/v1/. \
	--js_out=import_style=commonjs:./proto/v1/astro/webjs/ \
	--grpc-web_out=import_style=typescript,mode=grpcwebtext:./proto/v1/astro/webjs/ \
	proto/v1/astro/astro.proto \
	proto/v1/astro/dasha.proto \
	proto/v1/astro/panchang.proto \
	proto/v1/astro/metas.proto 

protos: ## Build Go and Web Proto
	# ****** Astro Calculations Related Proto 
	$(PROTOC) -I=./proto/v1/.  \
	--go_out=./proto/v1/astro/ \
    --go-grpc_out=./proto/v1/astro/ \
	proto/v1/astro/astro.proto \
	proto/v1/astro/dasha.proto \
	proto/v1/astro/panchang.proto \
	proto/v1/astro/metas.proto 
	make protos-web-js

build: ## Build docker image
# docker system prune -f
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) -t $(REPO)/$(APP_NAME) -f astro.dockerfile .
	

win: ## Build windows docker image
# docker system prune -f
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) -t $(REPO)/$(APP_NAME) -f windows.dockerfile .
	
run: ## Run astro microservice
	cd services/astrocalc/cmd && ./runAstro.sh	

runContainer: ## Start container
	./startAstroContainer.sh	