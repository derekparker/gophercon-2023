bin/traceme:
	go build -gcflags="all=-N -l" -o bin/traceme

.PHONY: trace
trace:
	dlv trace "^main.doWork$$"

.PHONY: trace-with-stack
trace-with-stack:
	dlv trace ^main.doWork --stack 10

.PHONY: trace-waitgroup
trace-waitgroup:
	dlv trace WaitGroup

.PHONY: trace-waitgroup-stack
trace-waitgroup-stack:
	dlv trace WaitGroup --stack=2

# Only works on Linux machines.
.PHONY: trace-with-ebpf
trace-with-ebpf:
	sudo env "PATH=$$PATH" dlv trace doWork --ebpf