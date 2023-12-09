PACKAGES=\
  github.com/tksasha/utils/strings \

.PHONY: all
all: vet fix fmt test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet $(PACKAGES)

.PHONY: fix
fix:
	@echo "go fix"
	@go fix $(PACKAGES)

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go fmt $(PACKAGES)

.PHONY: test
test:
	@echo "go test"
	@go test $(PACKAGES)
