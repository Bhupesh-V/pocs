## Pre-requisites

- Go
- Make

## Run Benchmark

```bash
make bench
```

## Benchmark Result

```
goos: darwin
goarch: arm64
pkg: benchmark-null-check
cpu: Apple M1
BenchmarkSumWithContinue-8           338           3434846 ns/op          355038 B/op      14792 allocs/op
BenchmarkSumWithIfCheck-8            340           3327715 ns/op          352950 B/op      14705 allocs/op
PASS
ok      benchmark-null-check    4.095s
```

The `BenchmarkSumWithIfCheck` is faster (3327715 ns/op vs 3434846 ns/op), uses slightly less memory (352950 B/op vs 355038 B/op), and has fewer allocations (14705 vs 14792).

```
goos: darwin
goarch: arm64
pkg: benchmark-null-check
cpu: Apple M1
BenchmarkSumWithContinue-8           337           3298837 ns/op          356092 B/op      14836 allocs/op
BenchmarkSumWithIfCheck-8            348           3306556 ns/op          344836 B/op      14367 allocs/op
PASS
ok      benchmark-null-check    4.035s
```

`BenchmarkSumWithContinue` is slightly faster in execution time.
`BenchmarkSumWithIfCheck` is better in memory efficiency (fewer bytes and allocations).

## Conclusion

Both `if` and `continue` are almost same in terms of performance. So, it's better to use `if` as it's more readable.