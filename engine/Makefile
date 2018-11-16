githash = $(shell git rev-parse --short HEAD)
ldflags = -X github.com/veritone/engine-tool/engine.BuildTag=$(githash)

buildgo:
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o ./dist/engine

release:
	mkdir -p ./release/engine-toolkit-sdk
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o release/engine-toolkit-sdk/engine
