FROM golang:1.16-alpine

ENV PROFILE=release
ENV APP_PORT=5005

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

EXPOSE $PORT
CMD ["/build/customer-service"]