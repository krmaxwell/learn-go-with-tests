package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeating 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Repeating 'a' 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	repeated := Repeat("k", 6)
	fmt.Println(repeated)
	// Output: kkkkkk
}
