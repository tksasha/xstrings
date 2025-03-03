MODFILE=-modfile go.tool.mod

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go tool $(MODFILE) gofumpt -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@go tool $(MODFILE) golangci-lint run

.PHONY: test
test:
	@echo "go test"
	@go test ./...

.PHONY: prepare
prepare:
	go get -tool $(MODFILE) mvdan.cc/gofumpt@latest
	go get -tool $(MODFILE) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
