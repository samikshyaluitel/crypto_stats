build_dependency:
	GOMOD="on" go build ./...

build:
	GOMOD="on" go build -o bin/crypto ./main.go

run:
	./bin/crypto