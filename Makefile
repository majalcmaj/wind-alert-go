docker-image=wind-alert:latest

clean:
	[ ! -d bin ] || rm -rf bin

test:
	go test -v ./...

build:
	go build -tags lambda.norpc -o bin/ ./...

format: 
	go fmt ./...

build-docker:
	docker buildx build --platform linux/amd64 --provenance=false -t $(docker-image) .

run-docker:
	docker run --name wind-alert --rm -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie $(docker-image) ./main

run-test-request:
	curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'


