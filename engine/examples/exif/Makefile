githash = $(shell git rev-parse --short HEAD)
ldflags = -X main.BuildTag=$(githash)

build: clean buildengine builddocker
run: buildengine builddocker up

buildengine:
	GOOS=linux GOARCH=amd64 go build -ldflags "${ldflags}" -o ./dist/engine ../../
	GOOS=linux GOARCH=amd64 go build -o ./dist/exif-extraction-engine

builddocker:
	docker build -f Dockerfile -t exif-extraction-engine .

up:
	docker run -t exif-extraction-engine

clean:
	rm -rf ./dist

push:
	docker tag exif-extraction-engine docker.veritone.com/15170/exif-extraction-engine:beta
	docker push docker.veritone.com/15170/exif-extraction-engine:beta
