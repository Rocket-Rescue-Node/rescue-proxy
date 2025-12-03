VERSION = v3.1.1

ABIGEN_CMD := go run github.com/ethereum/go-ethereum/cmd/abigen@v1.16.1 --v2

SOURCEDIR := .
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
PROTO_IN := proto
PROTO_OUT := pb
PROTOS := $(PROTO_OUT)/api.pb.go $(PROTO_OUT)/api_grpc.pb.go
PROTO_DEPS := $(wildcard $(PROTO_IN)/*.proto)

MULTICALL_ABI_DIR := executionlayer/dataprovider/abis
MULTICALL_ABI_JSON_DIR := $(MULTICALL_ABI_DIR)/json
ABI_ENCODINGS = $(MULTICALL_ABI_DIR)/multicall_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketstorage_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketnodemanager_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketnodedistributorfactory_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketminipoolmanager_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketdaonodetrusted_encoding.go \
	$(MULTICALL_ABI_DIR)/eip1271_encoding.go \
	$(MULTICALL_ABI_DIR)/vaultsregistry_encoding.go \
	$(MULTICALL_ABI_DIR)/ethprivvault_encoding.go

.PHONY: all
all: $(PROTOS) $(ABI_ENCODINGS)
	go build .

executionlayer/dataprovider/abis/multicall_encoding.go: $(MULTICALL_ABI_JSON_DIR)/multicall_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type Multicall3 --out $@
executionlayer/dataprovider/abis/rocketstorage_encoding.go: $(MULTICALL_ABI_JSON_DIR)/rocketstorage_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketStorage --out $@
executionlayer/dataprovider/abis/rocketnodemanager_encoding.go: $(MULTICALL_ABI_JSON_DIR)/rocketnodemanager_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketNodeManager --out $@
executionlayer/dataprovider/abis/rocketnodedistributorfactory_encoding.go: $(MULTICALL_ABI_JSON_DIR)/rocketnodedistributorfactory_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketNodeDistributorFactory --out $@
executionlayer/dataprovider/abis/rocketminipoolmanager_encoding.go: $(MULTICALL_ABI_JSON_DIR)/rocketminipoolmanager_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketMinipoolManager --out $@
executionlayer/dataprovider/abis/rocketdaonodetrusted_encoding.go: $(MULTICALL_ABI_JSON_DIR)/rocketdaonodetrusted_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketDaoNodeTrusted --out $@
executionlayer/dataprovider/abis/eip1271_encoding.go: $(MULTICALL_ABI_JSON_DIR)/eip1271_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type EIP1271 --out $@
executionlayer/dataprovider/abis/vaultsregistry_encoding.go: $(MULTICALL_ABI_JSON_DIR)/vaults-registry.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type VaultsRegistry --out $@
executionlayer/dataprovider/abis/ethprivvault_encoding.go: $(MULTICALL_ABI_JSON_DIR)/eth-priv-vault.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type EthPrivVault --out $@

$(PROTO_OUT):
	mkdir -p $@

$(PROTOS): $(PROTO_DEPS) $(PROTO_OUT)
	protoc -I=./$(PROTO_IN) --go_out=paths=source_relative:$(PROTO_OUT) \
		--go-grpc_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DEPS)

.PHONY: clean
clean:
	rm -f pb/*
	rm -f api-client
	rm -f $(ABI_ENCODINGS)

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

./api-client: $(PROTOS)
	go build -o api-client api/client/main.go
