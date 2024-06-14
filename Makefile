.PHONY: build
build:
	go build -o dem ./bin/dem.go
	go build -o deu ./bin/deu.go

.PHONY: install
install:
	sudo mv dem /usr/local/bin/
	sudo mv deu /usr/local/bin/
