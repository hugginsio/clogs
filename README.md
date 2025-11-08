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

BenchmarkLogger_MessageSizes/clogs_small-8          	15433291    77.28 ns/op    24 B/op    1 allocs/op
BenchmarkLogger_MessageSizes/standard_small-8       	 7812878    153.5 ns/op    24 B/op    2 allocs/op
BenchmarkLogger_MessageSizes/slog_small-8           	 2695798    446.3 ns/op     8 B/op    0 allocs/op

BenchmarkLogger_MessageSizes/clogs_medium-8         	14166360    85.44 ns/op    24 B/op    1 allocs/op
BenchmarkLogger_MessageSizes/standard_medium-8      	 7419481    162.9 ns/op    24 B/op    2 allocs/op
BenchmarkLogger_MessageSizes/slog_medium-8           	  837007     1435 ns/op     8 B/op    0 allocs/op

BenchmarkLogger_MessageSizes/clogs_large-8          	11016468    109.2 ns/op    24 B/op    1 allocs/op
BenchmarkLogger_MessageSizes/standard_large-8       	 5891317    204.1 ns/op    24 B/op    2 allocs/op
BenchmarkLogger_MessageSizes/slog_large-8            	  171866     7007 ns/op     8 B/op    0 allocs/op

BenchmarkLogger_MessageFormatting/clogs_single-8    	12112407    99.22 ns/op    21 B/op    2 allocs/op
BenchmarkLogger_MessageFormatting/standard_single-8 	 6090074    197.7 ns/op    24 B/op    1 allocs/op

BenchmarkLogger_MessageFormatting/clogs_double-8    	 9817215    122.9 ns/op    40 B/op    2 allocs/op
BenchmarkLogger_MessageFormatting/standard_double-8 	 4248105    283.7 ns/op    24 B/op    1 allocs/op

BenchmarkLogger_MessageFormatting/clogs_triple-8    	 8455057    141.6 ns/op    80 B/op    2 allocs/op
BenchmarkLogger_MessageFormatting/standard_triple-8 	 3386719    354.6 ns/op    24 B/op    1 allocs/op
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
