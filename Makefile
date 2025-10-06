clean:
	rm -rf bin/

build:
	go build -tags lambda.norpc -o bin/lambda

test:
	go test
