MAKEFLAGS += --jobs=2
.PHONY: yarn frontend server all

default: all

yarn:
	cd website && yarn

frontend: yarn
	cd website && yarn serve

server:
	go run cmd/chpunk/main.go server

all: frontend server
