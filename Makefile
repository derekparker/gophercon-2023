# ===================
# Go commands
# ====================

.PHONY: run
run:
	go run main.go -path ./count/testdata/urls.txt

.PHONY: test-count
test-count:
	go test ./count

# ===================
# Delve commands
# ====================

.PHONY: debug-test-count
debug-test-count:
	dlv test ./count

.PHONY: debug
debug:
	dlv debug -- -path=./count/testdata/urls.txt

# ===================
# Docker commands
# ====================

.PHONY: build-image
build-image:
	docker build -t count:latest -f deploy/Dockerfile .

.PHONY: build-debug-image
build-debug-image:
	docker build -t debug:latest -f deploy/Dockerfile-debug .

.PHONY: run-image
run-image:
	docker run --rm --name=count -d -p 3333:3333 count:latest

.PHONY: debug-image
debug-image:
	docker run -it --rm --privileged --pid="container:count" debug:latest

# ===================
# Misc commands
# ====================

.PHONY: curl-app
curl-app:
	curl localhost:3333/count?urls=http://google.com,http://yahoo.com,http://go.dev

