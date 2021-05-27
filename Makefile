all: setup build

setup:
		go install -v ./...

build:
		go build -o app