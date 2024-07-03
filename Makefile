.PHONY: pkg
pkg:
	cd packages && tar -z -c -f base.tar.gz base
.PHONY: build
build:
	go build -ldflags "-s -w" -o dem ./bin/dem.go
	go build -ldflags "-s -w" -o deu ./bin/deu.go

.PHONY: install
install:
	sudo mv dem /usr/local/bin/
	sudo mv deu /usr/local/bin/

.PHONY: build-all
build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dem ./bin/dem.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o deu ./bin/deu.go
	tar zcf darwin-arm64-${VERSION}.tar.gz dem deu
	mv darwin-arm64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dem ./bin/dem.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o deu ./bin/deu.go
	tar zcf darwin-amd64-${VERSION}.tar.gz dem deu
	mv darwin-amd64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dem ./bin/dem.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o deu ./bin/deu.go
	tar zcf linux-arm64-${VERSION}.tar.gz dem deu
	mv linux-arm64-${VERSION}.tar.gz ~/Downloads && rm dem deu

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dem ./bin/dem.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o deu ./bin/deu.go
	tar zcf linux-amd64-${VERSION}.tar.gz dem deu
	mv linux-amd64-${VERSION}.tar.gz ~/Downloads && rm dem deu
