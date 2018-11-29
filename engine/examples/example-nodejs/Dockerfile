FROM node:10.2-alpine as alpine

ARG GITHUB_ACCESS_TOKEN
ADD . /app
COPY manifest.json /var/
WORKDIR /app

RUN apk add --update git
RUN apk add --no-cache --virtual .gyp \
        -U --no-cache ca-certificates \
        python \
        make \
        g++
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN npm install
RUN npm rebuild sleep --force
ENV VERITONE_WEBHOOK_READY="http://0.0.0.0:8080/readyz"
ENV VERITONE_WEBHOOK_PROCESS="http://0.0.0.0:8080/process"
RUN chmod +x "/app/run.sh"

ENV KAFKA_BROKERS="localhost:9092"
ENV KAFKA_INPUT_TOPIC="input_topic"
ENV KAFKA_CONSUMER_GROUP="consumer_group"
ENV END_IF_IDLE_SECS=600
ENV KAFKA_CHUNK_TOPIC="kafka_chunk_topic"

ADD ./dist/engine /app/engine

ENTRYPOINT ["/app/engine", "/app/run.sh"]
