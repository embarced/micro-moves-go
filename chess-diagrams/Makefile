# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=chess-diagrams
SRC_FILE=cmd/chess-diagrams/main.go


all: test build
build:
		$(GOBUILD) -o $(BINARY_NAME) $(SRC_FILE)
build-docker:
		CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)

docker-image:
	docker build . -t chess-diagrams
