## Pre-requisites

- Go
- Make

## Run Benchmark

```bash
make bench
```

## Benchmark Result

```
Running benchmarks, please wait for atleast 20 seconds.
go test -bench=. -benchmem -benchtime=10s
goos: darwin
goarch: arm64
pkg: buffer-pool
cpu: Apple M1
BenchmarkWithBuffer-8             788197             12907 ns/op               0 B/op          0 allocs/op
BenchmarkWithoutBuffer-8          510500             22668 ns/op           64000 B/op       1000 allocs/op
PASS
ok      buffer-pool     22.519s
```

## Reading the Result

- `BenchmarkWithBuffer` is faster (12907 ns/op vs 22668 ns/op), uses less memory (0 B/op vs 64000 B/op), and has fewer allocations (0 vs 1000).

## Conclusion

- Reusing objects from the pool dramatically improves performance and reduces memory usage. Avoids frequent heap allocations, reducing garbage collection overhead.
- Use `sync.Pool` for scenarios where objects like buffers are reused frequently. This leads to significant performance gains and reduced memory overhead.
