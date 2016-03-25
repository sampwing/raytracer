package raytracer

import (
	"log"
	"math"
	"testing"
)

var x = 1.0
var y = 2.0
var z = 3.0

func TestVec3(t *testing.T) {
	//x := 1.0
	//y := 2.0
	//z := 3.0
	v := NewVec3(x, y, z)
	if v.X() != x {
		t.Fatalf("x: %f != %f\n", v.X(), x)
	}
	if v.Y() != y {
		t.Fatalf("y: %f != %f\n", v.Y(), y)
	}
	if v.Z() != z {
		t.Fatalf("z: %f != %f\n", v.Z(), z)
	}
}

func TestVec3Length(t *testing.T) {
	v := NewVec3(x, y, z)
	squaredLength := v.SquaredLength()
	squaredLengthActual := 1.0 + 4.0 + 9.0
	if squaredLength != squaredLengthActual {
		t.Fatalf("squaredLength: %f != %f\n", squaredLength, squaredLengthActual)
	}

	length := v.Length()
	lengthActual := math.Sqrt(squaredLengthActual)
	if length != lengthActual {
		t.Fatalf("length: %f != %f\n", length, lengthActual)
	}
}

func TestVec3Unit(t *testing.T) {
	v := NewVec3(x, y, z)
	t.Logf("%v\n", v)
	log.Printf("%v\n", v)
	v.MakeUnitVector()
	t.Logf("%v\n", v)
	log.Printf("%v\n", v)
}
