// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package clogs provides a simple logging package. It defines another [Logger] type with methods
// for printing log lines at various log levels - DBG, INF, WRN, and ERR. Like the standard logger,
// a predefined Logger is available with a ISO8601-like timestamp format and debug printing
// disabled.
package clogs

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

// Logger represents an active instance of clogs.
type Logger struct {
	buf        []byte
	debugMode  bool // Flag controlling the printing of debug level log messages.
	mu         sync.Mutex
	out        io.Writer // The log destination writer.
	timeFormat string    // The datetime format used by `time.AppendFormat`.
	lastTime   int64
	timeStr    []byte
}

var std = New(os.Stdout)

// New creates a new [Logger] with the specified output writer.
func New(out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}

	return &Logger{
		buf:        make([]byte, 0, 256),
		debugMode:  false,
		mu:         sync.Mutex{},
		out:        out,
		timeFormat: "2006-01-02 15:04:05",
		timeStr:    make([]byte, 0, 32),
	}
}

func (l *Logger) output(level string, v ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Reset buffer for reuse
	l.buf = l.buf[:0]

	if cap(l.buf) < 512 {
		l.buf = make([]byte, 0, 512)
	}

	l.buf = l.appendTime(l.buf)
	l.buf = append(l.buf, ' ')

	if level != "" {
		l.buf = append(l.buf, level...)
		l.buf = append(l.buf, ' ')
	}

	for i, arg := range v {
		if i > 0 {
			l.buf = append(l.buf, ' ')
		}

		// Fast path for common types
		switch v := arg.(type) {
		case string:
			l.buf = append(l.buf, v...)
		case int:
			l.buf = strconv.AppendInt(l.buf, int64(v), 10)
		case int32:
			l.buf = strconv.AppendInt(l.buf, int64(v), 10)
		case int64:
			l.buf = strconv.AppendInt(l.buf, v, 10)
		case uint:
			l.buf = strconv.AppendUint(l.buf, uint64(v), 10)
		case uint32:
			l.buf = strconv.AppendUint(l.buf, uint64(v), 10)
		case uint64:
			l.buf = strconv.AppendUint(l.buf, v, 10)
		case float32:
			l.buf = strconv.AppendFloat(l.buf, float64(v), 'f', -1, 32)
		case float64:
			l.buf = strconv.AppendFloat(l.buf, v, 'f', -1, 64)
		case bool:
			if v {
				l.buf = append(l.buf, "true"...)
			} else {
				l.buf = append(l.buf, "false"...)
			}
		default:
			// Fallback to fmt for uncommon types
			l.buf = fmt.Append(l.buf, v)
		}
	}

	l.buf = append(l.buf, '\n')
	l.out.Write(l.buf)
}

func (l *Logger) appendTime(buf []byte) []byte {
	now := time.Now()
	unix := now.Unix()

	// Cache time string for the same second
	if l.lastTime != unix {
		l.lastTime = unix
		l.timeStr = now.AppendFormat(l.timeStr[:0], l.timeFormat)
	}

	return append(buf, l.timeStr...)
}

// SetDebugMode sets the debug mode for the logger.
func (l *Logger) SetDebugMode(debug bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debugMode = debug
}

// SetDebugMode sets the debug mode for the logger.
func SetDebugMode(debug bool) {
	std.SetDebugMode(debug)
}

// Println writes a log message at the "INF" level.
func (l *Logger) Println(v ...any) {
	l.output("INF", v...)
}

// Println writes a log message at the "INF" level.
func Println(v ...any) {
	std.Println(v...)
}

// Debugln writes a log message at the "DBG" level.
func (l *Logger) Debugln(v ...any) {
	if l.debugMode {
		l.output("DBG", v...)
	}
}

// Debugln writes a log message at the "DBG" level.
func Debugln(v ...any) {
	if std.debugMode {
		std.output("DBG", v...)
	}
}

// Infoln writes a log message at the "INF" level.
func (l *Logger) Infoln(v ...any) {
	l.output("INF", v...)
}

// Infoln writes a log message at the "INF" level.
func Infoln(v ...any) {
	std.Infoln(v...)
}

// Warnln writes a log message at the "WRN" level.
func (l *Logger) Warnln(v ...any) {
	l.output("WRN", v...)
}

// Warnln writes a log message at the "WRN" level.
func Warnln(v ...any) {
	std.Warnln(v...)
}

// Errorln writes a log message at the "ERR" level.
func (l *Logger) Errorln(v ...any) {
	l.output("ERR", v...)
}

// Errorln writes a log message at the "ERR" level.
func Errorln(v ...any) {
	std.Errorln(v...)
}
