FROM golang:latest as builder

LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"

RUN mkdir /build
COPY ./ /build/
WORKDIR /build
ENV GO111MODULE auto
RUN CGO_ENABLED=0 GOOS=linux go build services/ephimeris/cmd/main.go   

# generate clean, final image for end users
# STAGE 2: build the container to run
FROM gcr.io/distroless/base-debian11
## Need to run dynamically linked bin, so using base-debian11

LABEL maintainer="binu@dronasys.com"

EXPOSE 50062

ARG BUILD_VER
ENV TRANCE_VERSION $BUILD_VER

# copy compiled app
COPY --chown=nonroot:nonroot services/ephimeris/cmd/astroBin/ /astroBin
COPY --chown=nonroot:nonroot --from=builder  /build/main /main
COPY --chown=nonroot:nonroot ./data/ /data/
# COPY services/ephimeris/cmd/astroBin/ /astroBin
# COPY  --from=builder  /build/main /main

WORKDIR /
USER nonroot:nonroot

# run binary; use vector form
ENTRYPOINT ["/main"]
# ENTRYPOINT ["sleep", "12000"]