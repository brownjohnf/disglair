.PHONY: build clean run
DEFAULT_GOAL: build

bin:
	mkdir bin

build: bin
	CGO_ENABLED=0 go build \
		-tags netgo \
		-o bin/disglair

run:
	source .env \
	&& CGO_ENABLED=0 go run main.go

clean:
	rm -rf bin

.PHONY: docker
docker:
	sed 's/%%BALENA_MACHINE_NAME%%/amd64/' Dockerfile \
	| docker build -

