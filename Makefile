SHELL := /bin/bash

run:
	go run main.go files.go misc.go system.go

compile:
	go build -o bin/
	GOOS=windows GOARCH=amd64 go build -o bin/
	GOOS=darwin GOARCH=amd64 go build -o bin/