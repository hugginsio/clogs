// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package clogs_test

import (
	"bytes"
	"log"
	"log/slog"
	"strings"
	"testing"

	"github.com/hugginsio/clogs"
)

func TestLogger_Println(t *testing.T) {
	tests := []struct {
		name     string
		args     []any
		contains []string
	}{
		{
			name:     "simple message",
			args:     []any{"Hello, World!"},
			contains: []string{"INF", "Hello, World!"},
		},
		{
			name:     "multiple arguments",
			args:     []any{"User", "logged in", 123},
			contains: []string{"INF", "User", "logged in", "123"},
		},
		{
			name:     "empty message",
			args:     []any{},
			contains: []string{"INF"},
		},
		{
			name:     "numbers and strings",
			args:     []any{42, "is the answer"},
			contains: []string{"INF", "42", "is the answer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := clogs.New(&buf)

			logger.Println(tt.args...)

			output := buf.String()

			for _, expected := range tt.contains {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain %q, got: %q", expected, output)
				}
			}

			if !strings.HasSuffix(output, "\n") {
				t.Errorf("Expected output to end with newline, got: %q", output)
			}

			// RemindMe! 1800 years
			if len(output) < 4 || (output[0] != '2') {
				t.Errorf("Expected output to start with timestamp, got: %q", output)
			}
		})
	}
}

func TestLogger_Println_OutputFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := clogs.New(&buf)

	logger.Println("test message")
	output := buf.String()

	// Should have format: "YYYY-MM-DD HH:MM:SS INF test message\n"
	parts := strings.SplitN(output, " ", 3)
	if len(parts) != 3 {
		t.Fatalf("Expected 3 parts (date, time, rest), got %d: %v", len(parts), parts)
	}

	// Check date format (YYYY-MM-DD)
	if len(parts[0]) != 10 || parts[0][4] != '-' || parts[0][7] != '-' {
		t.Errorf("Expected date format YYYY-MM-DD, got: %s", parts[0])
	}

	// Check time format (HH:MM:SS)
	if len(parts[1]) != 8 || parts[1][2] != ':' || parts[1][5] != ':' {
		t.Errorf("Expected time format HH:MM:SS, got: %s", parts[1])
	}

	// Check log level and message
	if !strings.HasPrefix(parts[2], "INF test message") {
		t.Errorf("Expected message to start with 'INF test message', got: %s", parts[2])
	}
}

func TestLogger_Printf_OutputFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := clogs.New(&buf)

	logger.Printf("digit: %d, value: %v", 123, "abc")
	output := buf.String()

	// Should have format: "YYYY-MM-DD HH:MM:SS INF digit: 123, value: abc\n"
	parts := strings.SplitN(output, " ", 3)
	if len(parts) != 3 {
		t.Fatalf("Expected 3 parts (date, time, rest), got %d: %v", len(parts), parts)
	}

	// Check date format (YYYY-MM-DD)
	if len(parts[0]) != 10 || parts[0][4] != '-' || parts[0][7] != '-' {
		t.Errorf("Expected date format YYYY-MM-DD, got: %s", parts[0])
	}

	// Check time format (HH:MM:SS)
	if len(parts[1]) != 8 || parts[1][2] != ':' || parts[1][5] != ':' {
		t.Errorf("Expected time format HH:MM:SS, got: %s", parts[1])
	}

	// Check log level and message
	if !strings.HasPrefix(parts[2], "INF digit: 123, value: abc") {
		t.Errorf("Expected message to start with 'INF digit: 123, value: abc', got: %s", parts[2])
	}
}

func TestLogger_ConcurrentAccess(t *testing.T) {
	var buf bytes.Buffer
	logger := clogs.New(&buf)

	const numGoroutines = 10
	const numMessages = 100

	done := make(chan bool, numGoroutines)

	for i := range numGoroutines {
		go func(id int) {
			for j := range numMessages {
				logger.Println("goroutine", id, "message", j)
			}
			done <- true
		}(i)
	}

	for range numGoroutines {
		<-done
	}

	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	expectedLines := numGoroutines * numMessages
	if len(lines) != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, len(lines))
	}

	for i, line := range lines {
		if !strings.Contains(line, "INF") {
			t.Errorf("Line %d missing INF: %s", i, line)
		}
	}
}

func BenchmarkLogger_MessageSizes(b *testing.B) {
	var buf bytes.Buffer
	logger := clogs.New(&buf)

	sizes := []struct {
		name string
		msg  string
	}{
		{"small", "short"},
		{"medium", strings.Repeat("medium length message ", 10)},
		{"large", strings.Repeat("this is a much longer message that will test buffer reallocation behavior ", 20)},
	}

	for _, size := range sizes {
		b.Run("clogs_"+size.name, func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				buf.Reset()
				logger.Println(size.msg, i)
			}
		})

		b.Run("standard_"+size.name, func(b *testing.B) {
			stdLogger := log.New(&buf, "", log.LstdFlags)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				buf.Reset()
				stdLogger.Println(size.msg, i)
			}
		})

		b.Run("slog_"+size.name, func(b *testing.B) {
			stdSlog := slog.New(slog.NewTextHandler(&buf, nil))

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				buf.Reset()
				stdSlog.Info(size.msg, "i", i)
			}
		})
	}
}

func BenchmarkLogger_MessageFormatting(b *testing.B) {
	var buf bytes.Buffer
	logger := clogs.New(&buf)

	formats := []struct {
		name   string
		format string
		values []any
	}{
		{"single", "id: %d", []any{1}},
		{"double", "id: %d, message: %s", []any{2, "string"}},
		{"triple", "id: %d, message: %s, extra: %s", []any{3, "another string message", "additional info"}},
	}

	for _, format := range formats {
		b.Run("clogs_"+format.name, func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				buf.Reset()
				logger.Printf(format.format, format.values...)
			}
		})

		b.Run("standard_"+format.name, func(b *testing.B) {
			stdLogger := log.New(&buf, "", log.LstdFlags)

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; b.Loop(); i++ {
				buf.Reset()
				stdLogger.Printf(format.format, format.values)
			}
		})
	}
}
