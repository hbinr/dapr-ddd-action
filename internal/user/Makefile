RELEASE_VERSION  =v0.0.1
SERVICE_NAME    ?=dapr-user-service
DOCKER_USERNAME ?=$(DOCKER_USER)

.PHONY: all
all: help

.PHONY: tidy
tidy: ## Updates the go modules and vendors all dependencies
	go mod tidy
	go mod vendor

.PHONY: test
test: tidy ## Tests the entire project
	go test -count=1 -race ./...

.PHONY: prod
run: tidy ## Runs uncompiled code in Dapr, log_level: debug/info/error/warning
	dapr run \
		--app-id $(SERVICE_NAME) \
		--app-port 8090 \
		--app-protocol http \
		--app-max-concurrency  10000000 \
		--dapr-http-port 3500 \
		--components-path ../../dapr/components \
		--log-level error \
		--config ./configs/config.yaml \
        go run  ./main.go

.PHONY: dev
dev: tidy ## Runs uncompiled code in Dapr, log_level: debug/info/error/warning
	dapr run \
		--app-id $(SERVICE_NAME) \
		--app-port 8090 \
		--app-protocol http \
		--app-max-concurrency  10000000 \
		--dapr-http-port 3500 \
		--components-path ../../dapr/components \
		--log-level debug \
		--config ./configs/config.yaml \
        go run  ./main.go

.PHONY: invoke
invoke: ## Invokes service through Dapr API
	curl -d '{ "message": "ping" }' \
     -H "Content-type: application/json" \
     "http://localhost:3500/v1.0/invoke/$(SERVICE_NAME)/method/echo"

.PHONY: image
image: tidy ## Builds and publish docker image
	docker build -t "$(DOCKER_USERNAME)/$(SERVICE_NAME):$(RELEASE_VERSION)" .
	docker push "$(DOCKER_USERNAME)/$(SERVICE_NAME):$(RELEASE_VERSION)"

.PHONY: deploy
deploy: ## Deploys prebuild image to k8s using currently selected context
	kubectl apply -f deployment.yaml
	kubectl rollout restart deployment/grpc-dapr-admin-service
	kubectl rollout status deployment/grpc-dapr-admin-service

.PHONY: call
call: ## Invokes service through Dapr API
	$(eval API_TOKEN=$(shell kubectl get secret dapr-api-token -o jsonpath="{.data.token}" | base64 --decode))
	curl -d '{ "message": "ping" }' \
     -H "Content-type: application/json" \
     -H "dapr-api-token: $(API_TOKEN)" \
     "https://api.cloudylabs.dev/v1.0/invoke/$(SERVICE_NAME)/method/echo"

.PHONY: lint
lint: ## Lints the entire project
	golangci-lint run --timeout=3m

.PHONY: tag
tag: ## Creates release tag
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)

.PHONY: clean
clean: ## Cleans up generated files
	go clean
	rm -fr ./bin
	rm -fr ./vendor

.PHONY: help
help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'