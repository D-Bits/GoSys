SHELL := /bin/bash

run:
	cd bin
	go run apiclient

compile:
	go build -o bin/
	GOOS=windows GOARCH=amd64 go build -o bin/
	GOOS=darwin GOARCH=amd64 go build -o bin/