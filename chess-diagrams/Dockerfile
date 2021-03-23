FROM golang:1.15-alpine3.12 as builder

RUN apk add --update make
RUN apk add --update git
RUN apk add --update gcc
RUN apk add --update musl-dev

RUN mkdir /build
WORKDIR /build
COPY Makefile /build/
RUN make deps

ADD . /build/
RUN make test
RUN make build-docker


FROM scratch

COPY images /app/images/
COPY fonts /app/fonts/
COPY html /app/html/

COPY --from=builder /build/chess-diagrams /app/
WORKDIR /app

CMD ["./chess-diagrams"]

EXPOSE 8080/tcp