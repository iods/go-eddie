WHOAMI ?= $(shell whoami)
CWD := $(shell pwd)

# TODO add some tests
# TODO add default shell
# TODO add build and install scripts
# TODO add docker build and install (run in docker always)

BIN_EDDIE := eddie
BIN_EDDIECTL := eddiectl
INSTALL_LOCATION := /usr/local/bin
BUILD_OUTPUT_DIR := $(CWD)/build
BINARY_LOCATION := $(BUILD_OUTPUT_DIR)/$(BIN_NAME)
MODULE := $(shell go list -m)
CMD_MODULE := $(MODULE)/cmd/$(BIN_NAME)
GOLANGCI_LINT_VERSION = latest


test:
	@go run cmd/eddie/main.go
	@echo $(WHOAMI)
	@echo $(CWD)

testctl:
	@go run cmd/eddiectl/main.go
