.PHONY: build
build:
	go build -ldflags "-s -w" -o dem ./bin/dem.go
	go build -ldflags "-s -w" -o deu ./bin/deu.go

.PHONY: install
install:
	sudo mv dem /usr/local/bin/
	sudo mv deu /usr/local/bin/
