githash = $(shell git rev-parse --short HEAD)
ldflags = -X main.BuildTag=$(githash)

build: clean buildengine builddocker
run: buildengine builddocker up

buildengine:
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o ./dist/engine ../../

builddocker:
	docker build -f Dockerfile -t example-nodejs .

up:
	docker run --network host -t example-nodejs

clean:
	rm -rf ./dist