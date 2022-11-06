#!/bin/bash

local_workdir=$(cd $(dirname $(dirname "${BASH_SOURCE[0]}")) >/dev/null 2>&1 && pwd)

main() {
  local container_workdir=/go/src/github.com/dronasys-com/astro
  local container_name=trance-astro   

  docker run --rm -it \
    --name $container_name \
    --volume $local_workdir:$container_workdir \
    --workdir $container_workdir \
    -p 127.0.0.1:50061:50061 \
    golang
}

main