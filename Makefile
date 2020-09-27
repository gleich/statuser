##########
# Building
##########

build-docker-dev:
	docker build -f dev.Dockerfile -t mattgleich/statuser:test .
build-docker-dev-lint:
	docker build -f dev.lint.Dockerfile -t mattgleich/statuser:lint .

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-hadolint:
	hadolint dev.Dockerfile
	hadolint dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/statuser:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev
	docker run mattgleich/statuser:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-dev build-docker-dev-lint
