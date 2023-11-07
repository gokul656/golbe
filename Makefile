BINARY_NAME = golbe

run:
	go run cmd/*

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/*
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/*
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/*