package echo_test

import (
	"strings"
	"testing"

	"github.com/yujin3i/transcription/gobook/ch01/ex03/echo"
)

func BenchmarkEcho1(b *testing.B) {
	data := strings.Split("test1 test2 test3 test4 test5", " ")

	for i := 0; i < b.N; i++ {
		echo.Echo1(data)
	}
}

func BenchmarkEcho2(b *testing.B) {
	data := strings.Split("test1 test2 test3 test4 test5", " ")

	for i := 0; i < b.N; i++ {
		echo.Echo2(data)
	}
}

func BenchmarkJoin(b *testing.B) {
	data := strings.Split("test1 test2 test3 test4 test5", " ")

	for i := 0; i < b.N; i++ {
		strings.Join(data, " ")
	}
}

// go test -bench=.                                                         
// goos: darwin
// goarch: amd64
// pkg: github.com/yujin3i/transcription/gobook/ch01/ex03/echo
// BenchmarkEcho1-4   	10000000	       127 ns/op
// BenchmarkEcho2-4   	10000000	       166 ns/op
// BenchmarkJoin-4    	20000000	        69.7 ns/op
// PASS
// ok  	github.com/yujin3i/transcription/gobook/ch01/ex03/echo	4.740s
