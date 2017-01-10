default: test

test:
	go test

run:
	go run *.go

build:
	go build -o dist/lacrosse

deps:
	glide i
