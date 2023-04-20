# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

include ./common.mk ./build-tooling.mk

.DEFAULT_GOAL:=help

SHELL := /usr/bin/env bash

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd"

ifeq ($(GOHOSTOS), linux)
XDG_DATA_HOME := ${HOME}/.local/share
endif
ifeq ($(GOHOSTOS), darwin)
XDG_DATA_HOME := "$${HOME}/Library/Application Support"
endif

# Directories
TOOLS_DIR := $(abspath hack/tools)
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
BIN_DIR := bin
GO_MODULES=$(shell find . -path "*/go.mod" | xargs -I _ dirname _)

ifndef IS_OFFICIAL_BUILD
IS_OFFICIAL_BUILD = ""
endif

# Package tooling related variables
PACKAGE_VERSION ?= ${BUILD_VERSION}
REPO_BUNDLE_VERSION ?= ${BUILD_VERSION}

# OCI registry for hosting tanzu framework components (containers and packages)
OCI_REGISTRY ?= projects.registry.vmware.com/tanzu_framework
export OCI_REGISTRY

##
## Install targets
##

.PHONY: install
# Install CRDs into a cluster
install: manifests tools
	kubectl apply -f apis/core/config/crd/bases

.PHONY: uninstall
# Uninstall CRDs from a cluster
uninstall: manifests
	kubectl delete -f apis/core/config/crd/bases

##
## Version
##

.PHONY: version
# Show generated build version
version:
	@echo $(BUILD_VERSION)

##
## Testing, verification, formating and cleanup
##

.PHONY: lint-all
# Run linting and misspell checks
lint-all: tools go-lint doc-lint misspell yamllint
	# Check licenses in shell scripts and Makefiles
	hack/check-license.sh

.PHONY: misspell
misspell:
	hack/check/misspell.sh

.PHONY: yamllint
yamllint:
	hack/check/check-yaml.sh


# These are the modules that still contain issues reported by golangci-lint.
# Once a module is updated to be lint-free, remove from list to catch lint regressions.
MODULES_NEEDING_LINT_FIX= . ./apis/core ./capabilities/client

.PHONY: go-lint
# Run linting of go source
go-lint: tools
	@for i in $(GO_MODULES); do \
		echo "-- Linting $$i --"; \
		pushd $${i} > /dev/null; \
		if echo ${MODULES_NEEDING_LINT_FIX} | tr ' ' '\n' | grep -q -x $$i; then \
			$(GOLANGCI_LINT) run -v --timeout=10m; \
			if [ $$? -eq 0 ]; then \
				echo "****** NOTE: $$i is now lint-free. Recommend it be excluded from MODULES_NEEDING_LINT_FIX in Makefile"; \
			else \
				echo; echo "****** WARNING: module $$i needs lint fixes"; echo; \
			fi; \
		else \
			$(GOLANGCI_LINT) run -v --timeout=10m || exit 1; \
		fi; \
		popd > /dev/null; \
	done

	# Prevent use of deprecated ioutils module
	@CHECK=$$(grep -r --include="*.go" --exclude="zz_generated*" ioutil .); \
	if [ -n "$${CHECK}" ]; then \
		echo "ioutil is deprecated, use io or os replacements"; \
		echo "https://go.dev/doc/go1.16#ioutil"; \
		echo "$${CHECK}"; \
		exit 1; \
	fi

.PHONY: doc-lint
# Run linting checks for docs
doc-lint: tools
	$(VALE) --config=.vale/config.ini --glob='*.md' ./
	# mdlint rules with possible errors and fixes can be found here:
	# https://github.com/DavidAnson/markdownlint/blob/main/doc/Rules.md
	# Additional configuration can be found in the .markdownlintrc file.
	hack/check-mdlint.sh

.PHONY: modules
# Runs go mod tidy to ensure modules are up to date.
modules:
	@for i in $(GO_MODULES); do \
		echo "-- Tidying $$i --"; \
		pushd $${i}; \
		$(GO) mod tidy || exit 1; \
		popd; \
	done

.PHONY: verify
# Run all verification scripts
verify: tools modules
	$(MAKE) smoke-build generate-go generate
	./hack/verify-dirty.sh

.PHONY: smoke-build
# Do a quick test that all directories can be built
smoke-build:
	@for i in $(GO_MODULES); do \
		echo "-- Building $$i --"; \
		pushd $${i}; \
		$(GO) build ./... || exit 1; \
		$(GO) clean -i ./... || exit 1; \
		popd; \
	done


##
## Code scaffolding targets
##

.PHONY: manifests
# Generate manifests e.g. CRD, RBAC etc.
manifests:
	$(MAKE) -C apis/config generate-manifests
	$(MAKE) -C apis/core generate-manifests

.PHONY: generate-go
# Generate code via go generate.
generate-go: $(COUNTERFEITER)
	@for i in $(GO_MODULES); do \
		echo "-- Running go generate ./... $$i --"; \
		pushd $${i}; \
		PATH=$(abspath hack/tools/bin):"$(PATH)" go generate ./...; \
		popd; \
	done

.PHONY: generate
# Generate code (legacy)
generate: tools
	$(MAKE) generate-controller-code
	@for i in $(GO_MODULES); do \
		echo "-- Running $(MAKE) -C $$i generate-controller-code --"; \
		pushd $${i}; \
		$(MAKE) -C $${i} generate-controller-code; \
		popd; \
	done

.PHONY: generate-fakes
# Generate fakes for writing unit tests
generate-fakes:
	$(GO) generate ./...
	$(MAKE) fmt

.PHONY: clean-generated-conversions
# Remove files generated by conversion-gen from the mentioned dirs. Example SRC_DIRS="./api/core/v1alpha1"
clean-generated-conversions:
	(IFS=','; for i in $(SRC_DIRS); do find $$i -type f -name 'zz_generated.conversion*' -exec rm -f {} \;; done)

.PHONY: generate-go-conversions
# Generate conversions go code
generate-go-conversions: $(CONVERSION_GEN)
	$(CONVERSION_GEN) \
		-v 3 --logtostderr \
		--input-dirs="./apis/core/v1alpha1,./apis/core/v1alpha2" \
		--build-tag=ignore_autogenerated_core \
		--output-base ./ \
		--output-file-base=zz_generated.conversion \
		--go-header-file=./hack/boilerplate.go.txt

.PHONY: generate-package-secret
# Generate the default package values secret. Usage: make generate-package-secret PACKAGE=capabilities tkr=v1.23.3---vmware.1-tkg.1 iaas=vsphere
generate-package-secret:
	@if [ -z "$(PACKAGE)" ]; then \
		echo "PACKAGE argument required"; \
		exit 1 ;\
	fi

	@if [ $(PACKAGE) == 'capabilities' ]; then \
	  ./capabilities/hack/generate-package-secret.sh -v tkr=${tkr} --data-value-yaml 'rbac.podSecurityPolicyNames=[${psp}]';\
	else \
	  echo "invalid PACKAGE: $(PACKAGE)" ;\
	  exit 1 ;\
	fi

.PHONY: create-package
# Stub out new package directories and manifests. Usage: make create-package PACKAGE_NAME=foobar
create-package:
	@hack/packages/scripts/create-package.sh $(PACKAGE_NAME)
