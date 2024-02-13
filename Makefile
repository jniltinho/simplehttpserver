GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOFLAGS = -v -ldflags="-s -w"
GOGET = $(GOCMD) get
BINARY_NAME = simplehttpserver


default: get build


get:
	@go mod tidy

build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) $(GOFLAGS)


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
