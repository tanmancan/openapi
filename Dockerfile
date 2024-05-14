FROM openapitools/openapi-generator-cli

RUN apt-get update -y && \  
    apt-get install -y \
    make

WORKDIR /app

COPY ./spec ./spec
COPY ./Makefile ./Makefile

ARG GIT_REMOTE_URL
ENV GIT_REMOTE_URL=${GIT_REMOTE_URL}

RUN make

ENTRYPOINT [ "make" ]