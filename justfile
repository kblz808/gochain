build:
	go build -o ./bin/gochain .

run: build
	./bin/gochain
