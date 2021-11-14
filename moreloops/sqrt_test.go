package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	got := Sqrt(2.0)
	want := 1.414213
	epsilon := 0.000001

	if math.Abs(got-want) > epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}
