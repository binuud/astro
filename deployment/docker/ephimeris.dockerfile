FROM --platform=$BUILDPLATFORM  golang:latest AS builder

LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"

ARG TARGETOS
ARG TARGETARCH

RUN mkdir /build
COPY ./ /build/
WORKDIR /build

ENV GO111MODULE=auto

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build cmd/ephimeris/main.go   

# generate clean, final image for end users
# STAGE 2: build the container to run
FROM  golang:latest 
#FROM gcr.io/distroless/base-debian11 ## for production
## Need to run dynamically linked bin, so using base-debian11

LABEL maintainer="binu@dronasys.com"

EXPOSE 50062

ARG BUILD_VER
ENV TRANCE_VERSION=$BUILD_VER

# copy compiled app
# COPY --chown=nonroot:nonroot ./astroBin/ /astroBin
# COPY --chown=nonroot:nonroot --from=builder  /build/main /main
# COPY --chown=nonroot:nonroot ./data/ /data/

COPY  ./astroBin/ /astroBin
COPY --from=builder  /build/main /main
COPY  ./data/ /data/

# COPY services/ephimeris/cmd/astroBin/ /astroBin
# COPY  --from=builder  /build/main /main

WORKDIR /
#USER nonroot:nonroot

# run binary; use vector form
ENTRYPOINT ["./main"]
