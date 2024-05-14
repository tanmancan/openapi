#!/bin/bash
GIT_REMOTE_URL=$(git config --get remote.origin.url)

docker build \
    --build-arg GIT_REMOTE_URL=${GIT_REMOTE_URL} \
    -t tanmancan/openapi:latest \
    .

docker run --rm \
    -v .:/app \
    -t tanmancan/openapi:latest \