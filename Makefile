export GO111MODULE=off

GO ?= go

BUILD_PATH := $(shell pwd)/build
BUILD_BIN_PATH := ${BUILD_PATH}/bin
COVERAGE_PATH := ${BUILD_PATH}/coverage
JUNIT_PATH := ${BUILD_PATH}/junit

GOLANGCI_LINT := ${BUILD_BIN_PATH}/golangci-lint
GINKGO := ${BUILD_BIN_PATH}/ginkgo

define go-build
	$(shell cd `pwd` && ${GO} build -o $(BUILD_BIN_PATH)/$(shell basename $(1)) $(1))
	@echo > /dev/null
endef

all: build test

.PHONY: build
build:
	${GO} build -ldflags '-s -w' ./pkg/...

${GINKGO}:
	$(call go-build,./vendor/github.com/onsi/ginkgo/ginkgo)

${GOLANGCI_LINT}:
	$(call go-build,./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint)

.PHONY: lint
lint: ${GOLANGCI_LINT}
	${GOLANGCI_LINT} run

.PHONY: test
test: ${GINKGO}
	rm -rf ${COVERAGE_PATH} && mkdir -p ${COVERAGE_PATH}
	rm -rf ${JUNIT_PATH} && mkdir -p ${JUNIT_PATH}
	${BUILD_BIN_PATH}/ginkgo \
		${TESTFLAGS} \
		-r -p \
		--cover \
		--randomizeAllSpecs \
		--randomizeSuites \
		--covermode atomic \
		--outputdir ${COVERAGE_PATH} \
		--coverprofile coverprofile \
		--succinct
	# fixes https://github.com/onsi/ginkgo/issues/518
	sed -i '2,$${/mode: atomic/d;}' ${COVERAGE_PATH}/coverprofile
	${GO} tool cover -html=${COVERAGE_PATH}/coverprofile -o ${COVERAGE_PATH}/coverage.html
	${GO} tool cover -func=${COVERAGE_PATH}/coverprofile | sed -n 's/\(total:\).*\([0-9][0-9].[0-9]\)/\1 \2/p'
	find . -name '*_junit.xml' -exec mv -t ${JUNIT_PATH} {} +

.PHONY: vendor
vendor:
	export GO111MODULE=on \
		${GO} mod tidy && \
		${GO} mod vendor && \
		${GO} mod verify
