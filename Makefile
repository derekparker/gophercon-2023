.PHONY: build
build:
	go build

.PHONY: run
run: build
	./gc2023 < urls.txt

.PHONY: debug
debug:
	dlv debug

.PHONY: attach
attach:
	dlv attach $$(pgrep gc2023)