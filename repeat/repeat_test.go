package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	reps := Repeat("Hi", 4)
	fmt.Println(reps)
	//Output: HiHiHiHi
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("Hi", 5)
	}
}

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	got := Repeat("a", 5)
	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)
	}
}
