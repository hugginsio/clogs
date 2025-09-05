# 👞 clogs

A simple and robust logger for Go.

## Performance

`clogs` is simpler than the standard library's `log` package and has some mild performance optimizations:

- Utilization of dedicated `append` methods for common types
- Cached time format for logs printed in the same second
- Attempts to reduce allocations by estimating required buffer size

The following benchmark results compare its performance to the standard library's `log` and `slog` packages.

```
goos: darwin
goarch: arm64
pkg: github.com/hugginsio/clogs
cpu: Apple M2
BenchmarkLogger_MessageSizes/clogs_small-8         	 8775430	       131.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_small-8      	 4398867	       273.1 ns/op	      24 B/op	       2 allocs/op
BenchmarkLogger_MessageSizes/slog_small-8          	 1559587	       769.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkLogger_MessageSizes/clogs_medium-8        	 8233572	       145.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_medium-8     	 4181818	       287.1 ns/op	      24 B/op	       2 allocs/op
BenchmarkLogger_MessageSizes/slog_medium-8         	  504582	      2379 ns/op	       8 B/op	       0 allocs/op
BenchmarkLogger_MessageSizes/clogs_large-8         	 6328874	       189.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/standard_large-8      	 3342482	       358.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkLogger_MessageSizes/slog_large-8          	  105213	     11391 ns/op	       8 B/op	       0 allocs/op
PASS
ok  	github.com/hugginsio/clogs	11.021s
```
