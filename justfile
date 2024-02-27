build:
	go build -o ./bin/gochain .

run: build
	./bin/gochain

clean:
	rm -rf ./tmp/**
