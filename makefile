VERSION?=0.0.1

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}
BINARY=fomo-proxy
TARFILE=fomo-proxy.tar.bz2
CONTRIBFILES=contrib/systemd/fomo-proxy.service contrib/script/install.sh
CONFIG=config.json
RMBIN=config.json fomo-proxy.service install.sh

VET_REPORT=vet.report
LINT_REPORT=lint.report
TEST_REPORT=test.report
TEST_XUNIT_REPORT=test.report.xml

OUTPUT=${CURRENT_DIR}/bin

OS := $(shell uname -s)
ifeq ($(OS),Darwin)

else

endif

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.compileAt=`date +%s`"

# Build the project
.PHONY: build clean dep lint run test vet package

all: clean vet fmt lint build test package

dep:
	dep ensure -v

build:
	cd cmd/fomo-proxy; go build $(LDFLAGS) -o $(OUTPUT)/$(BINARY)

test:
	env GOCACHE=off go test ./... 2>&1 | tee $(TEST_REPORT); go2xunit -fail -input $(TEST_REPORT) -output $(TEST_XUNIT_REPORT)

vet:
	go vet $$(go list ./...) 2>&1 | tee $(VET_REPORT)

fmt:
	goimports -w $$(go list -f "{{.Dir}}" ./... | grep -v /vendor/)

lint:
#	golint $$(go list ./...) | sed "s:^$(BUILD_DIR)/::" | tee $(LINT_REPORT)

package:
	for i in ${CONTRIBFILES} $(CONFIG); do cp $$i $(OUTPUT); done
	cd $(OUTPUT) && tar -jcf $(TARFILE) $(BINARY) ${RMBIN}
	cd $(OUTPUT) && for i in ${RMBIN}; do rm -f $$i; done

clean:
	-rm -f $(VET_REPORT)
	-rm -f $(LINT_REPORT)
	-rm -f $(TEST_REPORT)
	-rm -f $(TEST_XUNIT_REPORT)
	-rm -f $(BINARY)
	-rm -f $(OUTPUT)/*
