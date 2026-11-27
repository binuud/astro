FROM --platform=$BUILDPLATFORM golang:latest AS builder

# LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"
ARG TARGETOS
ARG TARGETARCH

RUN mkdir /build
COPY ./references/swisseph/ /build/
WORKDIR /build

RUN make