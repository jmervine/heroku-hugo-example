HEROKU_APP_URI?=

default: build server run

build:
	-hugo --baseURL="$(HEROKU_APP_URI)" --theme=bootswatch

server:
	godep save -r server/server.go
	cd server; go build

run:
	server/server
