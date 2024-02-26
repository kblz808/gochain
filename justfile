build:
	go build ./bin/gochain .

run: build
	./bin/gochain
