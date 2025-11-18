# go-map-keys

![Alt text](go-map-keys-bench.png "Optional title text")

Prefer using a loop with a pre-allocated array. In this example, I use [lo](https://github.com/samber/lo) to get the keys. The performance is comparable to using a pre-allocated array.

In upcoming posts, I will explore more features of the [lo](https://github.com/samber/lo) package.

## Why have use pre-allocated array?

Pre-allocated array, avoid repeated allocations and copies when tou build a slice. That reduce CPU time, memory usage and GC pressure.

## Run

```bash
docker build -t go-map-keys-bench:latest .
docker run --rm go-map-keys-bench:latest
```
