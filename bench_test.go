package cgobench

import (
	"testing"
	"encoding/json"
	"strings"
	"io"
)
func BenchmarkJSONCall(b *testing.B) {
	msg := `1`
	b.RunParallel(func(pb *testing.PB) {
		var dst int
		r := strings.NewReader(msg)
		dec := json.NewDecoder(r)
		for pb.Next() {
			r.Seek(0, io.SeekStart)
			if err := dec.Decode(&dst); err != nil {
				panic(err)
			}
		}
	})
}
// helper to cut down on boilerplate
func pbench(b *testing.B, f func()) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f()
		}
	})
	
}
// Same as above, but explicitly calling the inlineable Call func.
func BenchmarkEmptyCallInlineable(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Call()
		}
	})
}
func BenchmarkEmptyCall(b *testing.B) {
	pbench(b, Call)
}
func BenchmarkCgoCall(b *testing.B) {
	pbench(b, CgoCall)
}
