package exercise_loops_and_qfunctions

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	zp := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-zp) < 0.001 {
			break
		}
		zp = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
