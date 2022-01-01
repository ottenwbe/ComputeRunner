APP				?= ComputeMiddleware
TMPAPP			?= ComputeMiddleware-TMP
SNAPSHOT		?= ComputeMiddlewaere-SNAPSHOT

APP_CODE		= cmd/runner/*.go

APP_REPO		?= github.com/ottenwbe/recipes-manager
VERSIONPKG 		?= "$(APP_REPO)/appVersionString"
APP_VERSION		?= $(shell git describe --tags --always --match=v* 2> /dev/null || echo v0.0.0)
APP_GIT_HASH	= $(shell git rev-parse --short HEAD)

GO      		= go
GOFMT			= $(GO) fmt

.PHONY: release
release: ; $(info $(M) building executable…) @ ## Build the app's binary release version
	@$(GO) build \
		-tags release \
		-ldflags "-s -w" \
		-ldflags "-X $(VERSIONPKG)=$(APP_VERSION)" \
		-o $(APP)-$(APP_VERSION) \
		$(APP_CODE)

.PHONY: snapshot
snapshot:  ; $(info $(M) building snapshot…) @ ## Build the app's snapshot version
		@$(GO) build \
		-o $(SNAPSHOT) \
		-ldflags "-X $(VERSIONPKG)=$(APP_VERSION)" \
		$(APP_CODE)

.PHONY: start
start: ; $(info $(M) running the app locally…) @ ## Run the program's snapshot version
	@$(GO) build \
	    -o $(TMPAPP) \
    	-ldflags "-X $(VERSIONPKG)=$(APP_VERSION)" \
    	$(APP_CODE) && ./$(TMPAPP)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@for d in $$($(GO) list -f '{{.Dir}}' ./...); do \
		$(GOFMT) $$d/*.go  ; \
	 done

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
