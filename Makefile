.PHONY: pkg
pkg:
	cd packages && tar -z -c -f base.tar.gz base

.PHONY: build
build:
	$$(which dem) go build -ldflags "-s -w" ./cmd/dem
	$$(which dem) go build -ldflags "-s -w" ./cmd/deu

.PHONY: install
install:
	sudo mv dem /usr/local/bin/
	sudo mv deu /usr/local/bin/

.PHONY: build-all
build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $$(which dem) go build -ldflags "-s -w" ./cmd/dem
	GOOS=darwin GOARCH=arm64 $$(which dem) go build -ldflags "-s -w" ./cmd/deu
	tar zcf darwin-arm64-${VERSION}.tar.gz dem deu
	mv darwin-arm64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $$(which dem) go build -ldflags "-s -w" ./cmd/dem
	GOOS=darwin GOARCH=amd64 $$(which dem) go build -ldflags "-s -w" ./cmd/deu
	tar zcf darwin-amd64-${VERSION}.tar.gz dem deu
	mv darwin-amd64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $$(which dem) go build -ldflags "-s -w" ./cmd/dem
	GOOS=linux GOARCH=arm64 $$(which dem) go build -ldflags "-s -w" ./cmd/deu
	tar zcf linux-arm64-${VERSION}.tar.gz dem deu
	mv linux-arm64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $$(which dem) go build -ldflags "-s -w" ./cmd/dem
	GOOS=linux GOARCH=amd64 $$(which dem) go build -ldflags "-s -w" ./cmd/deu
	tar zcf linux-amd64-${VERSION}.tar.gz dem deu
	mv linux-amd64-${VERSION}.tar.gz ~/Downloads && rm dem deu
