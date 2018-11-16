default: test

get:
	go get github.com/mitchellh/gox

test:
	go test -v -parallel=4 .

run:
	go run cli.go main.go request.go route53.go

clean:
	rm -rf bin dist

build: get clean
	mkdir bin dist
	gox -osarch="darwin/amd64" \
		-osarch="linux/amd64" \
		-osarch="windows/amd64" \
		-output="bin/{{.OS}}_{{.Arch}}/{{.Dir}}"
	zip -j dist/lacrosse_darwin_amd64.zip bin/darwin_amd64/lacrosse
	zip -j dist/lacrosse_linux_amd64.zip bin/linux_amd64/lacrosse
	zip -j dist/lacrosse_windows_amd64.zip bin/windows_amd64/lacrosse.exe
