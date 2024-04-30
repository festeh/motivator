main:
	go run .

install:
	go build . && ln -s $(shell pwd)/motivator ~/.local/bin/motivator
