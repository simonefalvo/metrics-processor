BINARY_NAME = metrics-processor
IMAGE = smvfal/metrics-processor
TAG = latest
NATS_PORT = 30008
NATS_URL = http://192.168.99.100:${NATS_PORT}
SUBJECT = metrics

publish: docker-build docker-push

docker-build:
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE}:${TAG} .

docker-push:
	docker push ${IMAGE}:${TAG}

docker-run:
	docker run \
	--env NATS_URL=${NATS_URL} --env SUBJECT=${SUBJECT} \
	--volume $(shell pwd)/data:/root/data \
	${IMAGE}:${TAG}

run:
	go build -o bin/${BINARY_NAME}
	bin/${BINARY_NAME} -s ${NATS_URL} ${SUBJECT}

vendor:
	go mod vendor -v

build:
	go build -o bin/${BINARY_NAME}

fmt:
	gofmt -l -d $(shell find . -type f -name '*.go' -not -path "./vendor/*")

clean:
	rm -rf vendor/ bin/ data/
