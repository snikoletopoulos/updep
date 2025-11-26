# Strip debug info
GO_FLAGS += "-ldflags=-w -s"
# Avoid embedding the build path in the executable for more reproducible builds
GO_FLAGS += -trimpath

PROJECT_FILES = cmd/$(APP_NAME)/*.go pkg/*/*.go internal/*/*.go go.mod go.sum

OS = darwin linux win32
ARCH = arm64 x64

PLATFORMS = $(foreach os, $(OS), $(foreach arch, $(ARCH), $(os)-$(arch)))
PLATFORM_DIRS = $(foreach platform, $(PLATFORMS), npm/platforms/$(platform))
PLATFORM_FILES = $(foreach platform, $(PLATFORMS), npm/platforms/$(platform)/$(call binaryname,$(platform)))

APP_NAME = updep
binaryname = $(if $(findstring win32, $(1)),$(APP_NAME).exe,$(APP_NAME))

getos = $(subst win32,windows,$(firstword $(subst -, ,$(1))))
getarch = $(subst x64,amd64,$(lastword $(subst -, ,$(1))))

.PHONY: all
all: build

.PHONY: build
build: $(PLATFORM_FILES)

.PHONY: format
format:
	@go fmt ./...

npm/platforms/%/$(APP_NAME) npm/platforms/%/$(APP_NAME).exe: $(PROJECT_FILES)
	$(info Building $*)
	GOOS=$(call getos,$*) GOARCH=$(call getarch,$*) go build $(GO_FLAGS) -o $@ ./cmd/$(APP_NAME)

define npm-command
	pushd $(1);
	npm pkg get version;
	popd;
endef

.PHONY: npm
.ONESHELL:
npm:
	cd npm/platforms
	pwd

.PHONY: run
run:
	@go run ./cmd/$(APP_NAME)

.PHONY: install
install: build
	@go install $(GO_FLAGS) ./cmd/$(APP_NAME)

.PHONY: check
check: 
	echo todo

.PHONY: clean
clean:
	@rm -rf $(PLATFORM_FILES)

