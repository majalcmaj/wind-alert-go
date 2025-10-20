docker-image=wind-alert:latest
GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0

clean:
	[ ! -d bin ] || rm -rf bin

test:
	go test -v ./...

test-coverage:
	go test -timeout=30s -cover -coverprofile test-coverage.out ./... && go tool cover -html=test-coverage.out

build:
	go build -tags lambda.norpc -o bin/ ./...

fmt: 
	go fmt ./...

lint:
	go run $(GOLANGCI_LINT_PACKAGE) run

lint-fix:
	go run $(GOLANGCI_LINT_PACKAGE) run --fix

build-docker:
	docker buildx build --platform linux/amd64 --provenance=false -t $(docker-image) .

run-docker:
	docker run --name wind-alert --rm -p 9000:8080 --env-file .env --entrypoint /usr/local/bin/aws-lambda-rie $(docker-image) ./main

run-test-request:
	curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'

