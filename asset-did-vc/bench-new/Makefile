SHELL=/bin/bash

.PHONY: all install test

all: install test

# Installs caliper-cli
install:
	@pnpm install

# Runs the benchmark
test:
	pnpm exec caliper launch manager \
	--caliper-workspace ../../ \
	--caliper-networkconfig asset-did-vc/bench-new/networkConfig.yaml \
	--caliper-benchconfig asset-did-vc/bench-new/config.yaml \
	--caliper-flow-only-test \
	--caliper-fabric-gateway-enabled
