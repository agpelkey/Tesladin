build:
	go build -o ./bin/tesladin

run: build
	./bin/tesladin