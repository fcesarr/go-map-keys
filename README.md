# go-map-keys

![Alt text](go-map-keys-bench.png "Optional title text")

Prefer using a loop with a pre-allocated array. In this example, I use lo to get the keys. The performance is comparable to using a pre-allocated array.

## Run

```bash
docker build -t go-map-keys-bench:latest .
docker run --rm go-map-keys-bench:latest
```
