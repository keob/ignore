# Binary name
BINARY=ignore

build:
	go build -o ${BINARY}

release:
	go clean
	rm -rf *.gz *.zip
	# build for linux
	go clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	tar czvf ignore_linux_x64_${VERSION}.tar.gz ./ignore
	# build for mac
	go clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
	tar czvf ignore_darwin_x64_${VERSION}.tar.gz ./ignore
	# build for windows
	go clean
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
	zip -r ignore_windows_x64_${VERSION}.zip ./ignore.exe

clean:
	go clean
	rm -rf *.zip *.gz

.PHONY: clean build
