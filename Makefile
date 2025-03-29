clean: 
	rm ./bin

build_win: clean
	go build -o ./bin/app

run_win: build_win
	./bin/app.exe

build_linux: clean
	go build -o ./bin/app

run_linux: build_linux
	./bin/app

build_darwin: clean
	go build -o ./bin/app

run_darwin: build_darwin
	./bin/app

PHONY: clean build_win run_win build_linux run_linux build_darwin run_darwin
# This Makefile is used to build and run a Go application for different operating systems.