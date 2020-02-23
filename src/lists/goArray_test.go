package lists

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	var testStrArray = [...]string{"a", "b", "c", "d", "e"}
	var reverseArray = []string{"e", "d", "c", "b", "a"}
	var resultArray = ReverseArray(testStrArray[:])
	if !reflect.DeepEqual(reverseArray, resultArray) {
		t.Errorf("ReverseArray '%q' but '%q'", reverseArray, resultArray)
	}
}
func BenchmarkReverseArray(b *testing.B) {
	var testStrArray = [...]string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "c", "d", "e", "f", "g"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseArray(testStrArray[:])
	}
}
