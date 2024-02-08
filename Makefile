BIN_DIR=bin
BIN=beget
PATH_TO_ENTER_POINT=github.com/ThCompiler/go.beget.api/src

.PHONY: local-build
local-build: bin-dir
	go build -o $(BIN_DIR)/$(BIN) $(PATH_TO_ENTER_POINT)

.PHONY: build
build: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		go build -o $(BIN_DIR)/$(BIN) $(PATH_TO_ENTER_POINT) ; \
	else \
		echo Working directory not clean, commit changes; \
	fi

.PHONY: bin-dir
bin-dir:
	mkdir -p $(BIN_DIR)

.PHONY: release
release:
	git tag $(VERSION); \
	git push origin $(VERSION)

.PHONY: clean
changelog:
	sh ./workflow/changes.sh $(VERSION) > CURRENT-CHANGELOG.md

.PHONY: clean
clean:
	echo "Cleaning..."; \
	rm -Rf $(BIN_DIR)

.PHONY: fmt
fmt:
	gofumpt -e -w -d -extra .
