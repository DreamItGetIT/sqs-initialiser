.PHONY: build

dev-env: start-sqs

stop-dev-env:
	@docker rm -vf fake-sqs > /dev/null 2>&1 || true

start-sqs:
	@docker run -d -p 4568:4568 --name fake-sqs digit/docker-fake-sqs

build:
	@docker run --rm -v "$$(pwd):/go/src/github.com/DreamItGetIT/sqs-initialiser" \
		-w /go/src/github.com/DreamItGetIT/sqs-initialiser \
		-e GO15VENDOREXPERIMENT=1 \
		digit/go-build:v1.5.1 \
		make .build-in-docker

.build-in-docker:
	@go build -o build/sqsinit create_queues.go

package: build
	docker build -t docker.dreamitget.it/sqs-initialiser build

push-package:
	docker push docker.dreamitget.it/sqs-initialiser	
