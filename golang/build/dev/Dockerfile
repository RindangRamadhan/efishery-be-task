FROM golang:1.16-alpine

ENV SERVICE_NAME efishery-be-task
ENV REPOSITORY github.com/rindangramadhan
ENV GOPRIVATE github.com/rindangramadhan/*

RUN apk add --no-cache sudo

COPY build/assets/sudoers /etc/sudoers

RUN chmod 0400 /etc/sudoers && \
    addgroup efishery -g 1000 && \
    addgroup sudo && \
    adduser -h /home/efishery -s /bin/sh -u 1000 -G users -D efishery && \
    adduser efishery efishery && \
    adduser efishery sudo

USER efishery

RUN sudo chown -R efishery:efishery /go && \
    echo "Installing additional tools, please wait..." && \
    go get github.com/githubnemo/CompileDaemon && echo "- Success installing CompileDaemon" && \
    export GO111MODULE=auto

WORKDIR "/go/src/${REPOSITORY}/${SERVICE_NAME}"

ENTRYPOINT ["sh", "./build/dev/docker-entrypoint.sh"]
