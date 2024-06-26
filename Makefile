VERSION = v1.1.2

SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
PROTO_IN := proto
PROTO_OUT := pb
PROTO_DEPS := $(wildcard $(PROTO_IN)/*.proto)

.PHONY: all
all: protos
	go build .

.PHONY: protos
protos: $(PROTO_DEPS)
	protoc -I=./$(PROTO_IN) --go_out=paths=source_relative:$(PROTO_OUT) \
		--go-grpc_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DEPS)

.PHONY: clean
clean:
	rm -f pb/*
	rm -f api-client

.PHONY: docker
docker: all
	docker build . -t rocketrescuenode/rescue-proxy:$(VERSION)
	docker tag rocketrescuenode/rescue-proxy:$(VERSION) rocketrescuenode/rescue-proxy:latest
	docker tag rocketrescuenode/rescue-proxy:$(VERSION) rescue-proxy:latest

.PHONY: publish
publish:
	docker push rocketrescuenode/rescue-proxy:latest
	docker push rocketrescuenode/rescue-proxy:$(VERSION)

.DELETE_ON_ERROR: cov.out
cov.out: $(SOURCES)
	go test -coverprofile=cov.out ./...

.PHONY: testcov
testcov: cov.out
	go tool cover -html=cov.out

./api-client: protos
	go build -o api-client api/client/main.go
