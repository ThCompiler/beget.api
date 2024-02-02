BIN_DIR=bin
BIN=you_app  # TODO: set you app executable name
PATH_TO_ENTER_POINT=go-project/src   # TODO: set path to main file of you repository

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