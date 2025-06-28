# Etapa de build
FROM ubuntu:22.04 AS builder
WORKDIR /builder

ENV PATH="/usr/local/go/bin:${PATH}"

COPY . .

ADD https://golang.org/dl/go1.23.4.linux-amd64.tar.gz /go1.23.4.linux-amd64.tar.gz

RUN apt-get update -q && \
    apt-get install -y --no-install-recommends \
        build-essential \
        git \
        default-jre \
        maven \
        golang-go && \
    tar -C /usr/local -xzf /go1.23.4.linux-amd64.tar.gz && \
    make clean && \
    make build && \
    rm -rf /var/lib/apt/lists/*

#FROM cgr.dev/chainguard/go:latest