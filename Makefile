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
