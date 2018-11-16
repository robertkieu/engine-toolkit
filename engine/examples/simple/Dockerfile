FROM scratch

ADD ./dist/simple-engine /app/simple-engine

# Add and configure the engine
ADD manifest.json /var/manifest.json
ADD ./dist/engine /app/engine
ENV VERITONE_WEBHOOK_READY="http://0.0.0.0:8888/readyz"
ENV VERITONE_WEBHOOK_PROCESS="http://0.0.0.0:8888/process"
ENTRYPOINT [ "/app/engine", "/app/simple-engine" ]
