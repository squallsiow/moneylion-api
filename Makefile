setenv: 
	export GO111MODULE=on

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o moneylion-api

build-local:
	export GO111MODULE=on
	go build -o moneylion-api