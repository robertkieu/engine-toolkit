githash = $(shell git rev-parse --short HEAD)
ldflags = -X main.BuildTag=$(githash)

build: clean buildgo builddocker

buildgo:
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o ./dist/engine ../../
	GOOS=linux GOARCH=amd64 go build -o ./dist/simple-engine

builddocker:
	docker build -f Dockerfile -t simple-engine .

clean:
	rm -rf ./dist
