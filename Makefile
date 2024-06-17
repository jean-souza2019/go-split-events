# .PHONY: default run build test docs clean
# Variables
APP_NAME=go-split-events
BASE_PATH=/Users/jeanmarcos/Downloads/_projects/esocial-group/go-split-events
BASE_FILES=/Users/jeanmarcos/Downloads/_projects/esocial-group/files

# Tasks
default: run

run:
	@go run ${BASE_PATH}/main.go ${BASE_FILES}/input-teste S-3000
build-windows:
	@GOOS=windows GOARCH=amd64 go build -o ${BASE_PATH}/builds/${APP_NAME}-windows.exe ${BASE_PATH}/main.go
build-mac:
	@go build -o ${BASE_PATH}/builds/${APP_NAME}-macOS ${BASE_PATH}/main.go
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o ${BASE_PATH}/builds/${APP_NAME}-linux ${BASE_PATH}/main.go