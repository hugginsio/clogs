# 👞 clogs

A simple and robust logging library for Go.

## Performance

`clogs` is simpler than the standard library's `log` package and has some mild performance optimizations:

- Utilization of dedicated `append` methods for common types
- Cached time format for logs printed in the same second
- Attempts to reduce allocations by estimating required buffer size

The following benchmark results compare its performance to the standard library's `log` package.

```
❯ go test -bench=BenchmarkLogger -benchmem
goos: darwin
goarch: arm64
pkg: github.com/hugginsio/clogs
cpu: Apple M2
BenchmarkLogger_MessageSizes/clogs_small-8         	 9295224	       129.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_small-8      	 4399618	       274.8 ns/op	      24 B/op	       2 allocs/op
BenchmarkLogger_MessageSizes/clogs_medium-8        	 8371639	       143.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_medium-8     	 4195374	       286.0 ns/op	      24 B/op	       2 allocs/op
BenchmarkLogger_MessageSizes/clogs_large-8         	 6381142	       189.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_large-8      	 3339766	       357.5 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	github.com/hugginsio/clogs	7.484s
```
