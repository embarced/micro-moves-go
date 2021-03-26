FROM golang:1.15-alpine3.12 as builder

RUN apk add --update make
RUN apk add --update git
RUN apk add --update gcc
RUN apk add --update musl-dev

RUN mkdir /build
WORKDIR /build
COPY Makefile /build/
# RUN make deps

ADD . /build/
RUN make test
RUN make build-docker


FROM scratch

COPY assets/images /app/images/
COPY assets/fonts /app/fonts/
COPY assets/html /app/html/

COPY --from=builder /build/chess-diagrams /app/
WORKDIR /app

CMD ["./chess-diagrams"]

EXPOSE 8080/tcp