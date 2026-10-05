build:
	go build -o bin/first_project ./cmd/first_project

run:
	go run ./cmd/first_project $(ARGS)

build-run: build
	./bin/first_project $(ARGS)