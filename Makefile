docker-image=wind-alert:latest

clean:
	[ ! -d bin ] || rm -rf bin

test:
	go test

build:
	go build -tags lambda.norpc -o bin/lambda

build-docker:
	docker buildx build --platform linux/amd64 --provenance=false -t $(docker-image) .

run-docker:
	docker run --name wind-alert --rm -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie $(docker-image)


