include .env
build:
	go build -o cmd/ssh-conn main.go
install:
	@chmod +x cmd/ssh-conn
	@echo "creating alias."
	@echo "alias ssh-conn=$(shell pwd)/cmd/ssh-conn" >> ${SH}
	sudo cp cmd/ssh-conn /usr/bin

