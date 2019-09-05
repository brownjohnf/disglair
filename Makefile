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
	rm -rf bin build

.PHONY: docker
docker:
	mkdir -p build
	sed 's/%%BALENA_MACHINE_NAME%%/amd64/' Dockerfile.template > build/Dockerfile
	docker build -f build/Dockerfile .

