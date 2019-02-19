
#VERSION=$(shell git describe --abbrev=0 --tags)
#LDFLAG=-ldflags "-X github.com/mcbernie/se/share.BuildVersion=$(VERSION)"
BUILDDIR=build
RELEASEDIR=release

.DEFAULT_GOAL := help

.PHONY: help build win32 win64 arm5 arm7



help: ## Print a description of all available targets
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

create-folder:
	mkdir -p $(BUILDDIR)
	mkdir -p $(RELEASEDIR)


clean: ## Cleans build adn release directory
	rm -rf $(BUILDDIR)/*
	rm -rf $(RELEASEDIR)/*

win32: create-folder ## Builds binary for windows 32b
	env CGO_ENABLED=1 CC="/usr/local/bin/x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=386 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se.exe

win64: ## Builds binary for windows 64b
	env CGO_ENABLED=1 CC="/usr/local/bin/x86_64-w64-mingw32-gcc" GOOS=windows \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se.exe

linux64: ## Builds binary for linux
	env CGO_ENABLED=1 GOOS=linux GOARCH=386 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se-l64

linux32: ## Builds binary for linux
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se-l32

arm5: ## builds for raspberry arm
	env CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=5 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se-arm5

arm7: ## build for rpi 3
	env CC=arm-none-eabi-gcc CXX=arm-none-eabi-g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se-arm7

arm8: ## build arm8
	env GOOS=linux GOARCH=arm64 \
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se-arm8

build: ## Builds binary
	go build $(LDFLAG) \
	-o $(BUILDDIR)/se

all: linux32 linux64 win32 win64 arm build ## Builds all

release: clean all ## Build all and create archiv with version number
	tar cfzv $(RELEASEDIR)/release-$(VERSION).tar.gz $(BUILDDIR)/se32.exe $(BUILDDIR)/se64.exe $(BUILDDIR)/se
	zip $(RELEASEDIR)/release-$(VERSION).zip $(BUILDDIR)/se32.exe $(BUILDDIR)/se64.exe $(BUILDDIR)/se
