package vec

import (
	"math"
	"strconv"
)

type Vec struct {
	X int
	Y int
}

func (v *Vec) Add(o Vec) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vec) Sub(o Vec) {
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vec) Mul(i int) {
	v.X *= i
	v.Y *= i
}

func (v *Vec) Div(i int) {
	v.X /= i
	v.Y /= i
}

func (v *Vec) SqrMag() float64 {
	return float64(v.X*v.X + v.Y*v.Y)
}

func (v *Vec) Mag() float64 {
	return math.Sqrt(v.SqrMag())
}

func (v *Vec) ToStr() string {
	return "(" + strconv.Itoa(v.X) + ", " + strconv.Itoa(v.Y) + ")"
}
