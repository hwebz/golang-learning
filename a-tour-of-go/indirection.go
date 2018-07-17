package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	v.Scale(2) // Auto convert: v = (&v).Scale(2) (Scale required pointer param)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())     // Auto convert: p = (*p).Abs() (Abs required value param)
	fmt.Println(AbsFunc(*p)) // Required: p have to be a value (not a pointer)
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
