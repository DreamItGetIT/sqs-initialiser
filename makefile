.PHONY: build test

DOCKER_IP=$(shell ip route | grep docker0 | grep src | awk '{ print $$9 }')

default: full

dev-env: start-sqs

full: start-dev-env test stop-dev-env

stop-dev-env:
	@docker rm -vf fake-sqs > /dev/null 2>&1 || true

start-dev-env:
	@docker run -d -p 4568:4568 --name fake-sqs digit/docker-fake-sqs

test:
	@docker run --rm -e AWS_ACCESS_KEY_ID=DOESNOTMATTER -e AWS_SECRET_ACCESS_KEY=doesnotmatter digit/sqs-initialiser -endpoint $(DOCKER_IP):4568 -region eu-west-1 -ssl=false -queues "test1,test2"
	@docker run --rm -e AWS_ACCESS_KEY_ID=DOESNOTMATTER -e AWS_SECRET_ACCESS_KEY=doesnotmatter digit/sqs-initialiser -endpoint $(DOCKER_IP):4568 -region eu-west-1 -ssl=false -queues "test3"
	go test *_test.go

build-in-docker:
	@docker run --rm -v "$$(pwd):/go/src/github.com/DreamItGetIT/sqs-initialiser" \
		-w /go/src/github.com/DreamItGetIT/sqs-initialiser \
		-e GO15VENDOREXPERIMENT=1 \
		digit/go-build:v1.5.1 \
		go build -o build/sqsinit create_queues.go
build:
	@go build -o build/sqsinit create_queues.go

package:
	docker build -t digit/sqs-initialiser .

push-package:
	docker push digit/sqs-initialiser
