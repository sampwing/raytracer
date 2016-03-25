package raytracer

import (
	"math"
)

type Vec3 struct {
	Elements []float64
}

func NewVec3(point1, point2, point3 float64) *Vec3 {
	vec3 := &Vec3{
		Elements: []float64{point1, point2, point3},
	}
	return vec3
}

//getters
func (v Vec3) X() float64 {
	return v.Elements[0]
}

func (v Vec3) Y() float64 {
	return v.Elements[1]
}

func (v Vec3) Z() float64 {
	return v.Elements[2]
}

// operators
func VectorsMatch(v1 Vec3, v2 Vec3) bool {
	return v1.X() == v2.X() &&
		v1.Y() == v2.Y() &&
		v1.Z() == v2.Z()
}

func AddVectors(v1 Vec3, v2 Vec3) *Vec3 {
	u := v1.Elements
	w := v2.Elements
	return NewVec3(u[0]+w[0], u[1]+w[1], u[2]+w[2])
}

func SubtractVectors(v1 Vec3, v2 Vec3) *Vec3 {
	u := v1.Elements
	w := v2.Elements
	return NewVec3(u[0]-w[0], u[1]-w[1], u[2]-w[2])
}

func MultiplyVectors(v1 Vec3, v2 Vec3) *Vec3 {
	u := v1.Elements
	w := v2.Elements
	return NewVec3(u[0]*w[0], u[1]*w[1], u[2]*w[2])
}

func DivideVectors(v1 Vec3, v2 Vec3) *Vec3 {
	u := v1.Elements
	w := v2.Elements
	return NewVec3(u[0]/w[0], u[1]/w[1], u[2]/w[2])
}

func MultiplyVectorBy(v Vec3, f float64) *Vec3 {
	e := v.Elements
	return NewVec3(e[0]*f, e[1]*f, e[2]*f)
}

func DivideVectorBy(v Vec3, f float64) *Vec3 {
	e := v.Elements
	return NewVec3(e[0]/f, e[1]/f, e[2]/f)
}

func Dot(v1 Vec3, v2 Vec3) float64 {
	u := v1.Elements
	w := v2.Elements
	return u[0]*w[0] + u[1]*w[1] + u[2]*w[2]
}

func Cross(v1 Vec3, v2 Vec3) *Vec3 {
	u := v1.Elements
	w := v2.Elements
	return NewVec3(u[1]*w[2]-u[2]*w[1],
		-(u[0]*w[2] - u[2]*w[0]),
		u[0]*w[1]-u[1]*w[0])

}

func (v1 *Vec3) Add(v2 Vec3) {
	v1.Elements[0] += v2.Elements[0]
	v1.Elements[1] += v2.Elements[1]
	v1.Elements[2] += v2.Elements[2]
}

func (v1 *Vec3) Subtract(v2 Vec3) {
	v1.Elements[0] -= v2.Elements[0]
	v1.Elements[1] -= v2.Elements[1]
	v1.Elements[2] -= v2.Elements[2]
}

func (v1 *Vec3) Multiply(v2 Vec3) {
	v1.Elements[0] *= v2.Elements[0]
	v1.Elements[1] *= v2.Elements[1]
	v1.Elements[2] *= v2.Elements[2]
}

func (v1 *Vec3) Divide(v2 Vec3) {
	v1.Elements[0] /= v2.Elements[0]
	v1.Elements[1] /= v2.Elements[1]
	v1.Elements[2] /= v2.Elements[2]
}

func (v1 *Vec3) MultiplyBy(t float64) {
	v1.Elements[0] *= t
	v1.Elements[1] *= t
	v1.Elements[2] *= t
}

func (v1 *Vec3) DivideBy(t float64) {
	k := 1.0 / t
	v1.Elements[0] *= k
	v1.Elements[1] *= k
	v1.Elements[2] *= k
}

func UnitVector(v Vec3) *Vec3 {
	return DivideVectorBy(v, v.Length())
}

// func Perform(Operation, v1, v2) ...

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vec3) SquaredLength() float64 {
	elements := v.Elements
	var sum float64
	for i := 0; i < len(elements); i++ {
		sum += elements[i] * elements[i]
	}
	return sum
}

func (v *Vec3) MakeUnitVector() {
	k := 1.0 / v.Length()
	v.Elements[0] *= k
	v.Elements[1] *= k
	v.Elements[2] *= k
}
