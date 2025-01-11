.PHONY: build bench

build:
	go build -gcflags="-m" -o simple_stack_heap simple.go
	go build -gcflags="-m -m" -o struct_stack_heap struct.go
bench:
	hyperfine --warmup 5 --shell=none './simple_stack_heap stack 10000000' './simple_stack_heap heap 10000000'
	hyperfine --warmup 5 --shell=none './struct_stack_heap stack 10000000' './struct_stack_heap heap 10000000'