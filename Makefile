VERSION = v0.3.7 

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


./api-client: protos
	go build -o api-client api/client/main.go
