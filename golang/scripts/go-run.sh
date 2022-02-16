#!/bin/bash

SERVICE_NAME='efishery-be-task'

export LOAD_DOT_ENV=true

go mod vendor

CompileDaemon \
    -exclude-dir="build vendor deployments scripts" \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -mod=vendor -o $SERVICE_NAME ./cmd/$SERVICE_NAME/..." \
    -command="./$SERVICE_NAME"