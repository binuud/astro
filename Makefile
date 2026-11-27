# Maintainer Binu Udayakumar <binu@portworx.com>

# Optional environmental variables
REPO ?= dronasys-com
ASTRO_APP ?= astro-backend
COMBINED_APP ?= astro-combined
EPHIMERIS_APP ?= astro-ephimeris
AWS_ECR ?= 919350043131.dkr.ecr.ap-southeast-1.amazonaws.com

BUILD_VER ?= a1.0.0

DOCKER_HUB_TAG ?= a1.0.0
GOLANG_IMAGE=golang:1.27.1

TAGGED_NAME = $(REPO)/$(APP_NAME)

BASE_DIR	:=	$(shell git rev-parse --show-toplevel)
BIN			:=	$(BASE_DIR)/bin


MULTI_PLATFORM_DOCKER = --platform=linux/amd64,linux/arm64

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

build: ## Build docker image
# docker system prune -f
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(ASTRO_APP) -f ./deployment/docker/astro.dockerfile .
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(EPHIMERIS_APP) -f ./deployment/docker/ephimeris.dockerfile .
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(COMBINED_APP) -f ./deployment/docker/combined.dockerfile .

build-ephimeris: ## Build the ephimeris docker image
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(EPHIMERIS_APP) -f ./deployment/docker/ephimeris.dockerfile .

build-combined: ## Build astro combined image astro + ephimeris
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(COMBINED_APP) -f ./deployment/docker/combined.dockerfile .

build-mcp-server: ## Build astro mcp server
	go build -o $(GOBIN)/mcp-astro-server ./cmd/mcpServer/.

deploy: ## Deploy all images to docker hub
	@echo "====================> Pushing image to dockerhub$(AWS_ECR) ."
	## @docker login --username=$(DOCKER_HUB_USER) --password="$(DOCKER_HUB_PASSWD)"
	aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 919350043131.dkr.ecr.ap-southeast-1.amazonaws.com
	docker tag $(REPO)/${ASTRO_APP} $(AWS_ECR)/${ASTRO_APP}:$(DOCKER_HUB_TAG)
	docker tag $(REPO)/${EPHIMERIS_APP} $(AWS_ECR)/${EPHIMERIS_APP}:$(DOCKER_HUB_TAG)

	docker push $(AWS_ECR)/$(ASTRO_APP):$(DOCKER_HUB_TAG)
	docker push $(AWS_ECR)/$(EPHIMERIS_APP):$(DOCKER_HUB_TAG)

vendor: ## go vendor and tidy
	echo "tidy + vendor"
	go mod tidy
	go mod tidy

protos-web-js: ## WebJs Proto - working for streams too
	$(PROTOC) -I=./proto/. \
	--go_opt paths=source_relative \
	--js_out=import_style=commonjs:./proto/v1/astro/webjs/ \
	--grpc-web_out=import_style=typescript,mode=grpcwebtext:./proto/v1/astro/webjs/ \
	proto/v1/astro/astro.proto \
	proto/v1/astro/dasha.proto \
	proto/v1/astro/panchang.proto \
	proto/v1/astro/metas.proto \
	proto/v1/astro/cities.proto 

protos: ## Build Go and Web Proto
# ****** Astro Calculations Related Proto 
# 	--swift_out=./proto/v1/astro/swift/ 
# 	--grpc-swift_out=Client=true,Server=false:./proto/v1/astro/swift/ 
	$(PROTOC) -I=./proto/.  \
	--go_out=./gen/go/ \
	--go_opt paths=source_relative \
  --go-grpc_out=./gen/go/ \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out=./gen/go/ \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--grpc-gateway-ts_out=./gen/web/ \
	--grpc-gateway-ts_opt paths=source_relative \
	--grpc-gateway-ts_opt generate_unbound_methods=true \
	--oas_out ./gen/web/ \
	proto/v1/astro/metas.proto \
	proto/v1/astro/ephimeris.proto \
	proto/v1/astro/cities.proto \
	proto/v1/astro/astro.proto \
	proto/v1/astro/dasha.proto \
	proto/v1/astro/panchang.proto  
	#make protos-web-js
	yq eval ./gen/web/v1/astro/openapi.yaml -o=json -P > ./gen/web/v1/astro/openapi.json

compose-up: ## docker compose up
	docker-compose up  --remove-orphans 

start-container: ## run this to start a golang container, with this code base mounted
	docker run --rm -it -v .:/app -w /app dronasys-com/golang:1.27.1

clean-docker: ## Clean existing containers
	docker rm -f $(EPHIMERIS_APP)
	docker rm -f $(ASTRO_APP)

run-docker-astro: ## run container image
	docker rm -f $(ASTRO_APP)
	docker run	\
	-e MONGO_ENDPOINT="host.docker.internal:27017"	\
	-e MONGO_DB="astro"	\
	-e MONGO_USER=""	\
	-e MONGO_PASSWORD="" \
	-e MONGO_ENCRYPTION_KEY="" \
	-e ASTRO_EPHIMERIS_ENDPOINT="host.docker.internal:50062" \
	-p self.local:50061:50061 \
	--name $(ASTRO_APP) $(REPO)/$(ASTRO_APP)

run-docker-combined: ## run container for astro + ephimeris + cities
	docker rm -f $(COMBINED_APP)
	docker run	\
	-e MONGO_ENDPOINT="host.docker.internal:27017"	\
	-e MONGO_DB="astro"	\
	-e MONGO_USER=""	\
	-e MONGO_PASSWORD="" \
	-e MONGO_ENCRYPTION_KEY="" \
	-e ASTRO_EPHIMERIS_ENDPOINT="host.docker.internal:9090" \
	-p 9090:9090 \
	-p 9080:9080 \
	--platform linux/amd64 \
	--name $(COMBINED_APP) $(REPO)/$(COMBINED_APP)

run-docker-ephimeris: ## run container image
	docker rm -f $(EPHIMERIS_APP)
	docker run   \
	-e LOG_VERBOSITY="1" \
	-p 50062:50062 \
	--platform linux/amd64 \
	--name $(EPHIMERIS_APP) $(REPO)/$(EPHIMERIS_APP)

run-container: ## Start container
	./scripts/startAstroContainer.sh	

run-astro: ## Start astro
	./scripts/runAstro.sh

run: ## Start astro, cities, panchang, ephimeris server
	./scripts/runCombined.sh	

run-mcp-server: ## run the mcp server
	./scripts/runMcpServer.sh	

run-ephimeris: ## Run ephimeris microservice
	./scripts/runEphimeris.sh	

run-test-data: ## Creates test data charts
	./scripts/runTestData.sh	

run-scratch: ## Dev programs, to check functionality
	./scripts/runScratch.sh	

run-panchang: ## Create panchang json files ./scripts/panchangJson/panchang/
	./scripts/runPanchang.sh	

run-life-span: ## run life span and chart test
	./scripts/runLifeSpan.sh

test: ## Run tests
	export ASTRO_DATA_PATH=`pwd`/data/; \
	export ASTRO_BIN=`pwd`/astroBin/astroBin; \
	go test ./pkg/...

test-panchang: ## Run panchang
	export ASTRO_DATA_PATH=`pwd`/data/; \
	export ASTRO_BIN=`pwd`/astroBin/astroBin; \
	export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"; \
	go test -v --run Test_Pacnhang_Month ./pkg/astro/

test-cities: ## Run tests
	export ASTRO_DATA_PATH=`pwd`/data/; \
	export ASTRO_BIN=`pwd`/astroBin/astroBin; \
	go test ./pkg/servers/citiesServer/...	

test-astro-chart: ## Run tests on astro chart
	export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"; \  
	export ASTRO_DATA_PATH=`pwd`/data/; \
	export ASTRO_BIN=`pwd`/astroBin/astroBin; \
	go test -v ./pkg/astro/...	

test-coverage: ## Run test coverage
	export ASTRO_DATA_PATH=`pwd`/data/; \
	export ASTRO_BIN=`pwd`/astroBin/astroBin; \
	go test ./pkg/... -coverprofile=coverage.out

clone-submodules: ## clone submodules
	git submodule update --init --recursive	

build-astrobin-amd: ## build swiss ephimeris for amd64
	cd ./references/swisseph/ && make clean
	cd ./references/swisseph/ && docker run --rm --platform linux/amd64 -v  `pwd`:/src   -w /src gcc:latest make
	cd ./references/swisseph/bin && file *

build-astrobin-arm: ## build swiss ephimeris for arm64
	cd ./references/swisseph/ && make clean
	cd ./references/swisseph/ && docker run --rm --platform linux/arm64 -v `pwd`:/src -w /src gcc:latest make
	cd ./references/swisseph/bin && file *