APP_NAME = go_server
BUILD_DIR = ${PWD}/build

# Go Paramaters
clean-deps:
	go mod tidy

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) server.go

run:
	$(BUILD_DIR)/$(APP_NAME)

clean:
	rm server
