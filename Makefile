all: run

run:
	go run .

build:
	go build -o bin/oj .

release:
	go build -ldflags "-s -w" -o bin/oj .
	chmod u+x oj
