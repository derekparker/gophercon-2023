.PHONY: test-count
test-count:
	go test ./count

.PHONY: debug-test-count
debug-test-count:
	dlv test ./count
