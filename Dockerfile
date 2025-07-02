FROM ubuntu:22.04 AS builder
WORKDIR /builder

COPY . .

ADD https://golang.org/dl/go1.23.4.linux-amd64.tar.gz /go1.23.4.linux-amd64.tar.gz

RUN apt-get update -q && \
    apt-get install -y --no-install-recommends \
        build-essential \
        git \
        default-jre \
        maven && \
    tar -C /usr/local -xzf /go1.23.4.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

RUN make clean && \
    make build

#FROM cgr.dev/chainguard/go:latest