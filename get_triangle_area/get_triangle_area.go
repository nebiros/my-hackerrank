package get_triangle_area

import (
	"math"
)

func getTriangleArea(x []int32, y []int32) int64 {
	r1 := x[0] * (y[1] - y[2])
	r2 := x[1] * (y[2] - y[0])
	r3 := x[2] * (y[0] - y[1])

	return int64(math.Abs(float64(r1+r2+r3)) / 2)
}
