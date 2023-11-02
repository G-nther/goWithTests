package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeatStd(t *testing.T) {
	repeated := RepeatStd("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStd("a", 5)
	}
}

func TestRepeatBuilder(t *testing.T) {
	repeated := RepeatBuilder("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatBuilder("a", 5)
	}
}
