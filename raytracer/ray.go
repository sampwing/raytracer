package raytracer

type Ray struct {
	u *Vec3
	w *Vec3
}

func NewRay(u *Vec3, w *Vec3) *Ray {
	return &Ray{
		u: u,
		w: w,
	}
}

func (r Ray) Origin() *Vec3 {
	return r.u
}

func (r Ray) Direction() *Vec3 {
	return r.w
}

func (r Ray) PointAtParameter(t float64) *Vec3 {
	return MultiplyVectorBy(*r.u, t)
}
