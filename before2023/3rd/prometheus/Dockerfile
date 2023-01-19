# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.15.6 as builder

WORKDIR /app
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o metrics

ENTRYPOINT ["/app/metrics"]