.PHONY: shell-docker build-docker-image test

shell-docker:
	@docker run --rm -it -v '$(shell pwd):/go/src/project' dns-challenge:0.0.1 /bin/sh

build-docker-image:
	@docker build -t dns-challenge:0.0.1 .

test:
	@docker run --rm -it -v '$(shell pwd):/go/src/project' dns-challenge:0.0.1 go test
