.PHONY: shell-docker build-docker-image test

shell-docker:
	@docker run --rm -it -v '$(shell pwd):/go/src/project' dns-challenge:dev /bin/sh

build-docker-image-dev:
	@docker build -f Dockerfile.dev -t dns-challenge:dev .

build-docker-image:
	@docker build -t dns-challenge:0.0.1 .

run: build-docker-image
	@docker run -p 53:53 -p 9000:9000 dns-challange:0.0.1

test:
	@docker run --rm -it -v '$(shell pwd):/go/src/project' dns-challenge:0.0.1 go test

