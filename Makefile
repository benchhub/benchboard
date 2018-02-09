PKG = github.com/benchhub/benchboard/pkg
VERSION = 0.0.1
FLAGS = -X $(PKG)/common.version=$(VERSION) -X $(PKG)/common.gitCommit=$(BUILD_COMMIT) -X $(PKG)/common.buildTime=$(BUILD_TIME) -X $(PKG)/common.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/benchboard

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./lib ./pkg ./plugin