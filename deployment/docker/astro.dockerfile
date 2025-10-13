FROM --platform=$BUILDPLATFORM golang:latest AS builder

# LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"
ARG TARGETOS
ARG TARGETARCH

RUN mkdir /build
COPY ./ /build/
WORKDIR /build

ENV GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build cmd/astro/main.go   

 
# generate clean, final image for end users
# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final
 
LABEL maintainer="binu@dronasys.com"
USER nonroot:nonroot
EXPOSE 50061

# copy compiled app
COPY --from=builder --chown=nonroot:nonroot /build/main /main

ARG BUILD_VER
ENV TRANCE_VERSION=$BUILD_VER

ENTRYPOINT ["./main"]
 
# run binary; use vector form
#ENTRYPOINT ["sleep", "12000"]