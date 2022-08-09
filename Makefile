# All targets.
.PHONY: lint test build clean help

# Current version of the project.
VersionMain ?= v0.0.1
VersionPre ?=


# Target binaries. You can build multiple binaries for a single project.
TARGETS := mysqlgate
PLAT_FROM := linux darwin

# Project main package location (can be multiple ones).
CMD_DIR := ./cmd
# Project output directory.
OUTPUT_DIR := ./bin
# Git commit sha.
COMMIT := $(shell git rev-parse --short HEAD)
# Build Date
BUILD_DATE=$(shelldate +%FT%T%z)
# Version File
VERSION_FILE=github.com/u2takey/mysqlgate/version

build: build-local

lint:  ## use golint to do lint
	golint ./...

test:  ## run unit tests
	go test -cover ./...

build-local: ## local build
	@for target in $(TARGETS); do                                                      \
	  go build -o $(OUTPUT_DIR)/$${target}                                             \
	    -ldflags "-s -w                                                                \
				-X $(VERSION_FILE).VersionMain=$(VersionMain)                          \
			    -X $(VERSION_FILE).VersionPre=$(VersionPre)                            \
	            -X $(VERSION_FILE).VersionDev=build.$(COMMIT)  "                       \
	    $(CMD_DIR)/$${target};                                                         \
	done

build-cross: ## cross build
	@for plat in $(PLAT_FROM); do                                                       \
		for target in $(TARGETS); do                                                        \
			CGO_ENABLED=0 GOOS=$${plat} GOARCH=amd64 go build -o $(OUTPUT_DIR)/$${plat}/$${target} \
				-ldflags "-s -w                                                         \
						-X $(VERSION_FILE).VersionMain=$(VersionMain)                   \
						-X $(VERSION_FILE).VersionPre=$(VersionPre)                     \
						-X $(VERSION_FILE).VersionDev=build.$(COMMIT)  "                \
				$(CMD_DIR)/$${target};                                                  \
		done                                                                            \
	done

.PHONY: clean
clean:  ## clean bin files
	-rm -vrf ${OUTPUT_DIR}

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
