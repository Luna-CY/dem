.PHONY: pack-all
pack-all: pack-darwin-amd pack-darwin-arm pack-linux-amd pack-linux-arm move-to-download

.PHONY: pack-darwin-amd
pack-darwin-amd:
	env GOOS=darwin GOARCH=amd64 dem go build -o dem ./cmd/dem.go
	env GOOS=darwin GOARCH=amd64 dem go build -o dem-utils ./cmd/dem-utils.go
	tar zcf dem_darwin_x86_64_${VERSION}.tar.gz dem dem-utils

.PHONY: pack-darwin-arm
pack-darwin-arm:
	env GOOS=darwin GOARCH=arm64 dem go build -o dem ./cmd/dem.go
	env GOOS=darwin GOARCH=arm64 dem go build -o dem-utils ./cmd/dem-utils.go
	tar zcf dem_darwin_arm64_${VERSION}.tar.gz dem dem-utils

.PHONY: pack-linux-amd
pack-linux-amd:
	env GOOS=linux GOARCH=amd64 dem go build -o dem ./cmd/dem.go
	env GOOS=linux GOARCH=amd64 dem go build -o dem-utils ./cmd/dem-utils.go
	tar zcf dem_linux_x86_64_${VERSION}.tar.gz dem dem-utils

.PHONY: pack-linux-arm
pack-linux-arm:
	env GOOS=linux GOARCH=arm64 dem go build -o dem ./cmd/dem.go
	env GOOS=linux GOARCH=arm64 dem go build -o dem-utils ./cmd/dem-utils.go
	tar zcf dem_linux_aarch64_${VERSION}.tar.gz dem dem-utils

.PHONY: move-to-download
move-to-download:
	mv dem_*.tar.gz ~/Downloads
