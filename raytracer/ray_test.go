package raytracer

import (
	"testing"
)

var v1 = NewVec3(1.0, 2.0, 3.0)
var v2 = NewVec3(1.0, 2.0, 4.0)
var r = NewRay(v1, v2)

func TestRayOrigin(t *testing.T) {
	origin := r.Origin()
	if origin != v1 {
		t.Fatalf("failure w/ Origin: %v\n", origin)
	}
	direction := r.Direction()
	if direction != v2 {
		t.Fatalf("failure w/ Direction: %v\n", direction)
	}
	p := r.PointAtParameter(2.0)
	nv := NewVec3(2.0, 4.0, 6.0)
	if !VectorsMatch(p, nv) {
		t.Fatalf("failure w/ PointAtParameter: %v != %v\n", p, nv)
	}
}
