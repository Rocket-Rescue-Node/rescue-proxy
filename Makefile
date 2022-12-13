
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

.PHONE: clean
clean:
	rm -f pb/*
	rm -f api-client

./api-client: protos
	go build -o api-client api/client/main.go
