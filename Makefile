PLATFORMS := linux/amd64 linux/arm64 linux/386 linux/arm darwin/amd64
WINPLATFORMS := windows/amd64 windows/386 windows/arm

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.PHONY: help
help:  ## 🤔 Show help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

build: ## 🚧 Build for local arch
	mkdir -p ./bin
	go build -o ./bin/leetoclock-helper ./main.go

clean: ## 🧹 Remove previously build binaries
	rm -rf ./bin

pre-release:
	mkdir -p ./bin/release

release: pre-release $(PLATFORMS) $(WINPLATFORMS) ## 📦 Build for GitHub release
$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o ./bin/release/leetoclock-helper-$(os)-$(arch) ./main.go
$(WINPLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o ./bin/release/leetoclock-helper-$(os)-$(arch).exe ./main.go
