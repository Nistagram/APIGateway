FROM golang:1.16.4-alpine3.13 AS builder
LABEL maintainer="hesoyam"

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main . 

WORKDIR /dist
RUN cp /build/main .

FROM alpine:3.13.5

COPY --from=builder /dist/main /

ENTRYPOINT [ "/main" ]
