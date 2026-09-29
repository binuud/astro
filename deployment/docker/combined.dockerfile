FROM --platform=$BUILDPLATFORM golang:latest AS builder

# LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"
ARG TARGETOS
ARG TARGETARCH

RUN mkdir /build
COPY ./ /build/
WORKDIR /build

ENV GO111MODULE=auto

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build cmd/combined/main.go   
 
# generate clean, final image for end users
# STAGE 2: build the container to run
#FROM gcr.io/distroless/static AS final
FROM  ubuntu:latest 
#FROM  golang:latest 

LABEL maintainer="binu@dronasys.com"
# USER nonroot:nonroot

EXPOSE 9080
EXPOSE 9090

ARG BUILD_VER
ENV TRANCE_VERSION=$BUILD_VER

# copy compiled app
COPY --from=builder /build/main /main
COPY  ./astroBin/ /astroBin
COPY  ./data/ /data/

WORKDIR /



ENTRYPOINT ["./main"]

