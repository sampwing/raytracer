package raytracer

import (
	"math"
)

func (r *Ray) HasHitSphere(center *Vec3, radius float64) float64 {
	objectCenter := SubtractVectors(r.Origin(), center)
	a := Dot(r.Direction(), r.Direction())
	b := 2.0 * Dot(objectCenter, r.Direction())
	c := Dot(objectCenter, objectCenter) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func (r *Ray) Color() *Vec3 {
	t := r.HasHitSphere(NewVec3(0.0, 0.0, -1.0), 0.5)
	if t > 0.0 {
		nv := SubtractVectors(r.PointAtParameter(t), NewVec3(0.0, 0.0, -1.0))
		n := UnitVector(nv)
		v := NewVec3(n.X()+1.0, n.Y()+1.0, n.Z()+1.0)
		return MultiplyVectorBy(v, 0.5)
	}
	unitDirection := UnitVector(r.Direction())
	t = 0.5 * (unitDirection.Y() + 1.0)
	x := NewVec3(1.0, 1.0, 1.0)
	x.MultiplyBy(1.0 - t)
	y := NewVec3(0.5, 0.7, 1.0)
	y.MultiplyBy(t)
	return AddVectors(x, y)
}
