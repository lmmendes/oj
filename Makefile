all: run

run:
	go run .

build:
	go build

release:
	go build -ldflags "-s -w" .
	chmod u+x oj
