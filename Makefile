.PHONY: run
run:
	go run main.go < urls.txt

.PHONY: debug
debug:
	dlv debug

