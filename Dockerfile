ARG GO_VERSION=1.16
FROM golang:${GO_VERSION}-alpine AS builder
MAINTAINER Filipe Torqueto

ENV PROFILE=release
ENV APP_PORT=5005

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

EXPOSE $PORT
CMD ["/build/customer-service"]