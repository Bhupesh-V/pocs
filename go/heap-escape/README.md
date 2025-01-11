## Pre-requisites

- Go
- Make
- Hyperfine

## Build

```bash
make build
```

## Run Benchmark

```bash
make bench
```

## Benchmark Results

## Effect of stack vs heap allocation of integer by returning pointer v/s non-pointer

```
hyperfine --warmup 5 --shell=none './simple_stack_heap stack 10000000' './simple_stack_heap heap 10000000'
Benchmark 1: ./simple_stack_heap stack 10000000
  Time (mean ± σ):      10.4 ms ±   1.9 ms    [User: 4.6 ms, System: 0.9 ms]
  Range (min … max):     6.2 ms …  21.8 ms    255 runs
 
Benchmark 2: ./simple_stack_heap heap 10000000
  Time (mean ± σ):      11.5 ms ±   2.4 ms    [User: 5.5 ms, System: 0.9 ms]
  Range (min … max):     7.1 ms …  21.3 ms    259 runs
 
Summary
  ./simple_stack_heap stack 10000000 ran
    1.11 ± 0.31 times faster than ./simple_stack_heap heap 10000000
```

### Effect of stack vs heap allocation of struct by returning pointer v/s non-pointer

```
hyperfine --warmup 5 --shell=none './struct_stack_heap stack 10000000' './struct_stack_heap heap 10000000'
Benchmark 1: ./struct_stack_heap stack 10000000
  Time (mean ± σ):       5.6 ms ±   1.5 ms    [User: 1.0 ms, System: 1.1 ms]
  Range (min … max):     2.8 ms …  29.0 ms    932 runs
 
  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet system without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.
 
Benchmark 2: ./struct_stack_heap heap 10000000
  Time (mean ± σ):       5.5 ms ±   1.5 ms    [User: 1.0 ms, System: 1.1 ms]
  Range (min … max):     2.8 ms …  32.9 ms    965 runs
 
  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet system without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.
 
Summary
  ./struct_stack_heap heap 10000000 ran
    1.03 ± 0.39 times faster than ./struct_stack_heap stack 10000000
```

## Conclusion

For shortlived values, stack allocation is faster than heap allocation.

## Resources

- https://stackoverflow.com/questions/10866195/stack-vs-heap-allocation-of-structs-in-go-and-how-they-relate-to-garbage-collec