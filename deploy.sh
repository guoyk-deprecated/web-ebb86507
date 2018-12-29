#!/bin/bash

set -e
set -u

# build Echo
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web

# build Vue
cd ui && npm run build

# build docker image
docker build -t canhead/web .

# push docker image
docker push canhead/web
