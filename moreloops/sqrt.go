package sqrt

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 10
	z_prev := z
	delta := 1.0
	epsilon := 0.000001
	i := 0
	for delta > epsilon {
		z_prev = z
		z -= (z*z - x) / (2 * x)
		delta = math.Abs(z - z_prev)
		i++
		fmt.Println(z)
	}

	fmt.Printf("took %d tries\n", i)

	return z
}
