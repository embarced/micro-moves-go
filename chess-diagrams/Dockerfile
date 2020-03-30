FROM golang:alpine as builder
RUN apk add --update make
RUN apk add --update git

RUN mkdir /build
WORKDIR /build
COPY Makefile /build/
RUN make deps

ADD . /build/
RUN make build-docker

FROM scratch
COPY images /app/images/
COPY fonts /app/fonts/
COPY --from=builder /build/chess-diagrams /app/
WORKDIR /app

CMD ["./chess-diagrams"]