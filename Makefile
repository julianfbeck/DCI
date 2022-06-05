.PHONY: clean
clean: ## remove files created during build pipeline
	rm -rf dist
	rm -f coverage.*

.PHONY: install
install:
	go get github.com/stretchr/testify

.PHONY: test
test: ## go test with race detector and code covarage
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: run
run: ## go run
	@go run -race .

.PHONY: go-clean
go-clean: ## go clean build, test and modules caches
	go clean -r -i -cache -testcache -modcache
