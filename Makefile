# Maintainer Binu Udayakumar <binu@portworx.com>

# Optional environmental variables
REPO ?= dronasys-com
ASTRO_APP ?= astro-backend
EPHIMERIS_APP ?= astro-ephimeris
AWS_ECR ?= 919350043131.dkr.ecr.ap-southeast-1.amazonaws.com

BUILD_VER ?= a1.0.0

DOCKER_HUB_TAG ?= a1.0.0

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
	go mod vendor	

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
	--oas_out ./gen/web/v1/astro/ \
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
	
run: ## Run astro microservice
	cd services/astro/cmd && ./run.sh	

clean-docker: ## Clean existing containers
	docker rm -f $(EPHIMERIS_APP)
	docker rm -f $(ASTRO_APP)

run-docker-astro: ## run container image
	docker rm -f $(ASTRO_APP)
	docker run	\
	-e MONGO_ENDPOINT="host.docker.internal:27027"	\
	-e MONGO_DB="astro"	\
	-e MONGO_USER="root"	\
	-e MONGO_PASSWORD="rootpassword" \
	-e MONGO_ENCRYPTION_KEY="" \
	-e ASTRO_EPHIMERIS_ENDPOINT="host.docker.internal:50062" \
	-p self.local:50061:50061 \
	--name $(ASTRO_APP) $(REPO)/$(ASTRO_APP)

run-docker-ephimeris: ## run container image
	docker rm -f $(EPHIMERIS_APP)
	docker run   \
	-e LOG_VERBOSITY="1" \
	-p 50062:50062 \
	--platform linux/amd64 \
	--name $(EPHIMERIS_APP) $(REPO)/$(EPHIMERIS_APP)

runContainer: ## Start container
	./scripts/startAstroContainer.sh	

runEphimeris: ## Run ephimeris microservice
	./scripts/runEphimeris.sh	

runTestData: ## Creates test data charts
	./scripts/runTestData.sh	

test: ## Run tests
	go test ./pkg/...

testCities: ## Run tests
	go test ./pkg/servers/citiesServer/...	

testCoverage: ## Run test coverage
	go test ./pkg/... -coverprofile=coverage.out
