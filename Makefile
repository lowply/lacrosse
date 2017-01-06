default: test

run:
	go run *.go

build:
	go build -o bin/lacrosse

deps:
	glide i
