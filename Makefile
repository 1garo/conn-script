include .env
build:
	go build -o cmd/ssh-conn main.go
install: build
	chmod +x cmd/ssh-conn
	@echo "alias ssh-conn=$(shell pwd)/cmd/ssh-conn" >> ${SH}
	@echo "source ${SH}"
