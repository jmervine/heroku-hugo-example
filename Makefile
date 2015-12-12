default: build build/server

setup:
	go get -u -v github.com/spf13/hugo
	go get -u -v github.com/tools/godep

build:
	hugo --theme bootswatch

build/server:
	godep save -r server/server.go
	cd server; go build

