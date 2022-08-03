lint:
	@sh -c "'$(CURDIR)/scripts/lint.sh'"

coverage:
	@sh -c "'$(CURDIR)/scripts/coverage.sh'"

build:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

pipeline:
	make lint && make coverage && make build

run:
	go run ./cmd/main.go

.PHONY: lint, coverage, build, pipeline

.DEFAULT_GOAL := pipeline
