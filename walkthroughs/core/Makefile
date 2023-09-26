.PHONY: build-image
build-image:
	docker build -t coredebug:latest .

.PHONY: docker
docker: build-image
	docker run -it -w /src -v $$(pwd):/src coredebug:latest /bin/bash

.PHONY: build-crash
build-crash:
	go build -o bin/crash

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash

.PHONY: set-core-pattern
set-core-pattern:
	echo core | sudo tee /proc/sys/kernel/core_pattern

.PHONY: dlv-core
dlv-core:
	dlv core ./bin/crash core