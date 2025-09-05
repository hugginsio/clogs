# 👞 clogs

A simple and robust logger for Go.

## Preview

```
2025-09-04 22:36:31 DBG You should be able to see this now that DEBUG is enabled.
2025-09-04 22:36:31 INF PrintLn is an alias for INFO.
2025-09-04 22:36:31 INF This method prints INFO.
2025-09-04 22:36:31 WRN This method prints WARN.
2025-09-04 22:36:31 ERR This method prints ERROR. All methods support multiple arguments, like this!
```

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

## Example Usage

```go
package main

import (
	"errors"

	"github.com/hugginsio/clogs"
)

func main() {
	clogs.Debugln("This log line won't appear. DEBUG is off by default.")
	clogs.SetDebugMode(true)
	clogs.Debugln("You should be able to see this now that DEBUG is enabled.")

	clogs.Println("PrintLn is an alias for INFO.")
	clogs.Infoln("This method prints INFO.")
	clogs.Warnln("This method prints WARN.")
	clogs.Errorln("This method prints ERROR. All methods support multiple arguments,", errors.New("like this!"))
}
```
