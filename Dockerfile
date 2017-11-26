FROM alpine

COPY bin/linux386/server /app/server
COPY build/docker /
RUN mkdir /storage-proxy-scope
VOLUME /storage-proxy-scope

WORKDIR /app

ENTRYPOINT ["/entrypoint.sh"]
