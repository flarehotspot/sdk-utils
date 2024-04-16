package logger

import "testing"

func BenchmarkReadLogs(b *testing.B) {
	ReadLogs(0, 29170)
}

func BenchmarkGetLines(b *testing.B) {
	GetLogLines()
}
