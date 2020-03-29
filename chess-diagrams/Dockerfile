#
# Minimal Microservice Go Container using Docker Builder Pattern.
#
FROM golang:alpine as builder
RUN apk add --update make
RUN apk add --update git
RUN apk add --update gcc

RUN mkdir /build
ADD . /build/
WORKDIR /build
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN make deps
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o chess-diagrams .

FROM scratch
COPY images /app/images/
COPY fonts /app/fonts/
COPY --from=builder /build/chess-diagrams /app/
WORKDIR /app

CMD ["./chess-diagrams"]