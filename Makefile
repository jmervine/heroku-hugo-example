HEROKU_APP_URI?=

default: build build/server

build:
	hugo --baseURL="$(HEROKU_APP_URI)" --theme=bootswatch

build/server:
	godep save -r server/server.go
	cd server; go build

server: build/server
	server/server
