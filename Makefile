# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

.PHONY: all test coverage
all: test coverage

checkfmt:
	@echo 'Checking gofmt';\
 	bash -c "diff -u <(echo -n) <(gofmt -d .)";\
	EXIT_CODE=$$?;\
	if [ "$$EXIT_CODE"  -ne 0 ]; then \
		echo '$@: Go files must be formatted with gofmt'; \
	fi && \
	exit $$EXIT_CODE

lint:
	golangci-lint run

benchmark:
	$(GOTEST) ./benchmarks/bench_realtime_ops/. -run=XXX -bench=.  -benchtime=100000000x
	$(GOTEST) ./benchmarks/bench_deferrable_ops/. -run=XXX -bench=.  -benchtime=100000000x

get:
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GOGET) -t -v ./...

fmt:
	$(GOFMT) ./...

test: get fmt lint
	$(GOTEST) -race -covermode=atomic ./...

coverage: get test
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic .

