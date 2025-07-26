VERSION = v2.1.1

ABIGEN_CMD := go run github.com/ethereum/go-ethereum/cmd/abigen@v1.16.1 --v2

SOURCEDIR := .
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
PROTO_IN := proto
PROTO_OUT := pb
PROTO_DEPS := $(wildcard $(PROTO_IN)/*.proto)

SW_DIR := executionlayer/stakewise
SW_ABI_DIR := $(SW_DIR)/abis
MULTICALL_ABI_DIR := executionlayer/dataprovider/abis
ABI_ENCODINGS = $(MULTICALL_ABI_DIR)/multicall_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketstorage_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketnodemanager_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketnodedistributorfactory_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketminipoolmanager_encoding.go \
	$(MULTICALL_ABI_DIR)/rocketdaonodetrusted_encoding.go \
	$(SW_DIR)/vaults-registry-encoding.go \
	$(SW_DIR)/eth-priv-vault-encoding.go

.PHONY: all
all: protos $(ABI_ENCODINGS)
	go build .

executionlayer/dataprovider/abis/multicall_encoding.go: executionlayer/dataprovider/abis/multicall_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type Multicall3 --out $@
executionlayer/dataprovider/abis/rocketstorage_encoding.go: executionlayer/dataprovider/abis/rocketstorage_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketStorage --out $@
executionlayer/dataprovider/abis/rocketnodemanager_encoding.go: executionlayer/dataprovider/abis/rocketnodemanager_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketNodeManager --out $@
executionlayer/dataprovider/abis/rocketnodedistributorfactory_encoding.go: executionlayer/dataprovider/abis/rocketnodedistributorfactory_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketNodeDistributorFactory --out $@
executionlayer/dataprovider/abis/rocketminipoolmanager_encoding.go: executionlayer/dataprovider/abis/rocketminipoolmanager_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketMinipoolManager --out $@
executionlayer/dataprovider/abis/rocketdaonodetrusted_encoding.go: executionlayer/dataprovider/abis/rocketdaonodetrusted_abi.json
	$(ABIGEN_CMD) --abi $< --pkg abis --type RocketDaoNodeTrusted --out $@

.PHONY: protos
protos: $(PROTO_DEPS)
	protoc -I=./$(PROTO_IN) --go_out=paths=source_relative:$(PROTO_OUT) \
		--go-grpc_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DEPS)

$(SW_DIR)/vaults-registry-encoding.go: $(SW_ABI_DIR)/vaults-registry.json
	$(ABIGEN_CMD) --abi $< --pkg stakewise --type vaultsRegistry --out $@
$(SW_DIR)/eth-priv-vault-encoding.go: $(SW_ABI_DIR)/eth-priv-vault.json
	$(ABIGEN_CMD) --abi $< --pkg stakewise --type ethPrivVault --out $@

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

./api-client: protos
	go build -o api-client api/client/main.go
