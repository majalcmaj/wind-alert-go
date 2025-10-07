clean:
	[ ! -d bin ] || rm -rf bin

test:
	go test

build:
	go build -tags lambda.norpc -o bin/lambda

build-docker:
	docker buildx build --platform linux/amd64 --provenance=false -t wind-alert:latest .

