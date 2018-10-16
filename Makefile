
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -parallel 5
GOGET=$(GOCMD) get
MAIN_FILE=app/server.go
BINARY_NAME=bin/server
BINARY_UNIX=$(BINARY_NAME)_unix

ci: deps short-test build
deploy: clean deps test build
all: clean deps test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)
test:
	$(GOTEST) -v ./...
short-test:
	$(GOTEST) -short -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)
	./$(BINARY_NAME)
deps:
	# Secrets
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/gin-gonic/gin
	# Weather forecast
	$(GOGET) github.com/mlbright/forecast/v2

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(MAIN_FILE)
