GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = simplehttpserver


default: build

build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
