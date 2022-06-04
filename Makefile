.PHONY: build
build: build-image
	@mkdir -p ./build; \
	docker run --rm -it -e HOME=/tmp \
		--entrypoint=/bin/bash -v "$${PWD}:/go/src/" \
			-w "/go/src/" emoji-picker-builder "-c" "make build-local USERID=$${UID}"


.PHONY: build-local
build-local:
# @go mod tidy
	@packr2 -v
	@go build -o ./build/emoji-picker
	@chown -R $(USERID):$(USERID) ./build
	@chown -R $(USERID):$(USERID) ./packrd

.PHONY: build-gui
build-gui:
	docker run --rm -it -u $$UID -e HOME=/tmp \
		--entrypoint=/bin/sh -v "$${PWD}/gui:/app" \
			-w "/app" node:slim "-c" "npm install && npm run build"

.PHONY: build-all
build-all: build-gui build


.PHONY: run
run:
	./build/emoji-picker

.PHONY: build-image
build-image:
	docker build -t emoji-picker-builder .;

.PHONY: clean
clean:
	docker rmi emoji-picker-builder;
	rm -rf ./build;
	rm -rf ./packrd;


.PHONY: deploy
deploy:
	sudo cp ./build/emoji-picker /usr/local/bin/emoji-picker;
