package _type

import "testing"

func BenchmarkReverseArray(b *testing.B) {
	var testStrArray = [...]string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		ReverseArray(testStrArray[:])
	}
}
