FROM golang:1.15-alpine3.12 as builder

RUN apk add --update make
RUN apk add --update git
RUN apk add --update gcc
RUN apk add --update musl-dev

RUN mkdir /build
WORKDIR /build
COPY Makefile /build/

ADD . /build/
RUN make build-docker

FROM scratch

COPY assets /app/assets/
COPY --from=builder /build/chess-diagrams /app/
WORKDIR /app

CMD ["./chess-diagrams"]

EXPOSE 8080/tcp