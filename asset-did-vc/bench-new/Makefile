SHELL=/bin/bash

.PHONY: all install test

all: install test

# Installs caliper-cli
install:
	@pnpm install

# Runs the benchmark
test:
	@cd ../../ && \
	npx caliper launch manager \
	--caliper-workspace . \
	--caliper-networkconfig asset-did-vc/bench/networkConfig.yaml \
	--caliper-benchconfig asset-did-vc/bench/config.yaml \
	--caliper-flow-only-test \
	--caliper-fabric-gateway-enabled
