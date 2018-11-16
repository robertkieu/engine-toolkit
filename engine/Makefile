githash = $(shell git rev-parse --short HEAD)
ldflags = -X github.com/veritone/engine-tool/engine.BuildTag=$(githash)

.PHONY: release buildgo clean test dist

dist: clean test buildgo release

buildgo:
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o ./dist/engine

release:
	mkdir -p ./release/engine-toolkit-sdk-v0.1.0
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o release/engine-toolkit-sdk-v0.1.0/bin/engine
	cd release && tar cvvzf engine-toolkit-sdk-v0.1.0.tar.gz engine-toolkit-sdk-v0.1.0

clean:
	rm -rf release

test:
	go test -v ./...
