package logger_test

import (
	"testing"

	logger "github.com/flarehotspot/core/internal/utils/logger"
)

func TestGetCallerFileLine(t *testing.T) {
	t.Run("getting file and line number of caller", func(t *testing.T) {
		callDepth := 0
		gotFile, gotLine := logger.GetCallerFileLine(callDepth)
		wantFile, wantLine := "/home/marcvincentbentoy/dev/flarehotspot/core/internal/utils/logger/logger_test.go", 14 // TODO : refactor to proper retrieving of filenames

		assertCorrectFile(t, gotFile, wantFile)
		assertCorrectLine(t, gotLine, wantLine)
	})
}

func assertCorrectFile(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertCorrectLine(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLineCounter(t *testing.T) {
	// how can I make this test
	// requirements:
	// - file with content
	// - lineCounter should return the number of lines inside the file
	// steps:
	// 1. create a dummy file
	// 2. insert data
}

// func TestGetLogLines(t *testing.T) {
// 	want := 299
// 	lines := logger.GetLogLines()

// 	if want != lines {
// 		t.Fatalf(`Expected: %d, actual: %d`, want, lines)
// 	}
// }

// func TestReadLogs(t *testing.T) {
// 	// read logs should have end - start lines
// 	start := 0
// 	end := 99

// 	want := end - start + 1
// 	result, err := logger.ReadLogs(start, end)
// 	if err != nil {
// 		log.Println("Error reading logs", err)
// 	}

// 	if len(result) != want {
// 		t.Errorf("Result logs read was incorrect, got: %d, want: %d", len(result), want)
// 	}
// }

// func BenchmarkReadLogs(b *testing.B) {
// 	logger.ReadLogs(0, 29170)
// }

// func BenchmarkGetLines(b *testing.B) {
// 	logger.GetLogLines()
// }
