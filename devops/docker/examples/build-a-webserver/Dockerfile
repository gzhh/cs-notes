FROM golang:1.18 AS build-stage

WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /webserver

FROM scratch AS run-stage

WORKDIR /

COPY --from=build-stage /webserver /webserver

EXPOSE 8080

ENTRYPOINT ["/webserver"]
