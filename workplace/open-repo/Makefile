TARGET=thienht
DIR=build
$(warning *******************************************************)
help:
	@echo "Available targets:"
	@echo "  docker-compose-up   - Run Docker Compose services in the background."
	@echo "  docker-compose-down - Stop and remove Docker Compose containers."
	@echo "  docker-compose-build - Build Docker Compose services."
	@echo "  docker-compose-ps   - Show the status of Docker Compose services."
	@echo "  docker-compose-logs - View logs of Docker Compose services."
	@echo "  docker-compose-custom - Run a custom Docker Compose command."
	
all: clean build
init:
	export PATH=$PATH:/usr/local/go/bin
doc:
	./swag init -g pkg/controller/*
run:
	go run cmd/dev/main.go
clean: 
	rm -rf $(DIR)
build_pro:
	go vet
	go fmt
	go build -o $(DIR)/$(TARGET) main.go
	scp $(DIR)/$(TARGET) :'/home/thienhang/gochat/core'
run-doc:
	go run cmd/shop.thienhang.com/main.go server


