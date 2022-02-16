#!/bin/bash
#export DOCKER_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')

export DOCKER_HOST="host.docker.internal"

ping -q -c1 $DOCKER_HOST > /dev/null 2>&1
if [ $? -ne 0 ]; then
  HOST_IP=$(ip route | awk 'NR==1 {print $3}')
  echo "${HOST_IP}	${DOCKER_HOST}" | sudo tee -a /etc/hosts > /dev/null
fi

sudo chown -R efishery:efishery /go/pkg
sudo chown -R efishery:efishery ./vendor

echo "Vendoring, please wait..."
go mod vendor

CompileDaemon \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -mod=vendor -o $SERVICE_NAME ./cmd/$SERVICE_NAME/..." \
    -command="./${SERVICE_NAME}"
    