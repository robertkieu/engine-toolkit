FROM alpine:latest
RUN apk --no-cache add ca-certificates
ADD dist/exif-engine /app/exif-engine

ADD manifest.json /var/manifest.json
ADD ./dist/engine /app/engine
ENV VERITONE_WEBHOOK_READY="http://0.0.0.0:8080/readyz"
ENV VERITONE_WEBHOOK_PROCESS="http://0.0.0.0:8080/process"
ENTRYPOINT [ "/app/engine", "/app/exif-engine" ]
