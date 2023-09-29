# Dockerfile

## Images
Docker Official Image
- https://hub.docker.com/search?image_filter=official&q=&type=image

## Multi-stage builds
- https://docs.docker.com/build/building/multi-stage/

Multi-stage builds are useful to anyone who has struggled to optimize Dockerfiles while keeping them easy to read and maintain.

```
FROM golang:1.21 AS builder
WORKDIR /src
COPY main.go ./main.go
RUN go build -o /bin/hello ./main.go

FROM alpine
COPY --from=builder /bin/hello /bin/hello
CMD ["/bin/hello"]
```


## 命令
### FROM
基础镜像 base image
- scratch
  - https://hub.docker.com/_/scratch
  - an explicitly empty image, especially for building images "FROM scratch"
- busybox
  - https://en.wikipedia.org/wiki/BusyBox
  - https://hub.docker.com/_/busybox
  - Busybox base image.
- alpine
  - https://en.wikipedia.org/wiki/Alpine_Linux
  - https://hub.docker.com/_/alpine
  - A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size!

语言镜像
- Go
  - golang:1.21-alpine
  - golang:1.21


### CMD
Running multiple CMD
- https://serverfault.com/questions/685697/multiple-commands-in-docker-cmd-directive
- https://stackoverflow.com/questions/49630960/dockerfile-running-multiple-cmd-starting-nginx-and-php

