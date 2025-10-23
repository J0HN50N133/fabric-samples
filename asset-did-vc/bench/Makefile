SHELL=/bin/bash

.PHONY: all install test

all: install test

# Installs caliper-cli
install:
	@npm install --prefix ../../caliper-benchmarks @hyperledger/caliper-cli@0.5.0

# Runs the benchmark
test:
	@cd ../../ && \
	npx caliper launch manager \
	--caliper-workspace . \
	--caliper-networkconfig test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml \
	--caliper-benchconfig asset-did-vc/bench/config.yaml \
	--caliper-flow-only-test \
	--caliper-fabric-gateway-enabled
