.PHONY: run
run:
	go run main.go < urls.txt

.PHONY: debug
debug:
	dlv debug

.PHONY: debug-with-redirect
debug-with-redirect:
	dlv debug -r stdin:urls.txt
