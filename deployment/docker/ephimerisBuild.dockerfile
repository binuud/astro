FROM gcc:4.9 AS builder

# LABEL maintainer="Binu Udayakumar <binu@dronasys.com>"

# COPY ephimerals/ /ephimerals/
# WORKDIR /ephimerals/src

# RUN  make astro
 
# LABEL maintainer="binu@dronasys.com"

 
# generate clean, final image for end users
# STAGE 2: build the container to run
# FROM gcr.io/distroless/static AS final
 
LABEL maintainer="binu@dronasys.com"
#USER nonroot:nonroot


COPY ./astroBin/ /astroBin/
WORKDIR /astroBin

ARG BUILD_VER
ENV TRANCE_VERSION=$BUILD_VER


 
# run binary; use vector form
ENTRYPOINT ["sleep", "12000"]