package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sampwing/raytracer/raytracer"
)

func ch1() {
	filename := "ch1.ppm"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to create %s", filename)
	}
	defer f.Close()

	nx := 200
	ny := 100

	line := fmt.Sprintf("P3\n%d %d\n255\n", nx, ny)
	f.WriteString(line)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			col := raytracer.NewVec3(r, g, b)
			ir := int(255.99 * col.X())
			ig := int(255.99 * col.Y())
			ib := int(255.99 * col.Z())
			line = fmt.Sprintf("%d %d %d\n", ir, ig, ib)
			f.WriteString(line)
		}
	}

	f.Sync()
}

func hitSphere(center raytracer.Vec3, radius float64, r raytracer.Ray) bool {
	objectCenter := raytracer.SubtractVectors(*r.Origin(), center)
	a := raytracer.Dot(*r.Direction(), *r.Direction())
	b := 2.0 * raytracer.Dot(*objectCenter, *r.Direction())
	c := raytracer.Dot(*objectCenter, *objectCenter) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

func color(r raytracer.Ray) *raytracer.Vec3 {
	if hitSphere(*raytracer.NewVec3(0.0, 0.0, -1.0), 0.5, r) {
		return raytracer.NewVec3(1.0, 0.0, 0.0)
	}
	unitDirection := raytracer.UnitVector(*r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	x := raytracer.NewVec3(1.0, 1.0, 1.0)
	x.MultiplyBy(1.0 - t)
	y := raytracer.NewVec3(0.5, 0.7, 1.0)
	y.MultiplyBy(t)
	return raytracer.AddVectors(*x, *y)
}

func ch4() {
	filename := "ch4.ppm"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to create %s", filename)
	}
	defer f.Close()

	nx := 200
	ny := 100

	line := fmt.Sprintf("P3\n%d %d\n255\n", nx, ny)
	f.WriteString(line)

	lowerLeftCorner := raytracer.NewVec3(-2.0, -1.0, -1.0)
	horizontal := raytracer.NewVec3(4.0, 0.0, 0.0)
	vertical := raytracer.NewVec3(0.0, 2.0, 0.0)
	origin := raytracer.NewVec3(0.0, 0.0, 0.0)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			// lowerLeftCorner + u * horizontal + v * vertical
			ho := raytracer.MultiplyVectorBy(*horizontal, u)
			ve := raytracer.MultiplyVectorBy(*vertical, v)
			lhv := raytracer.AddVectors(*lowerLeftCorner, *ho)
			lhv = raytracer.AddVectors(*lhv, *ve)

			r := raytracer.NewRay(origin, lhv)

			col := color(*r)
			//b := 0.2
			//col := raytracer.NewVec3(r, g, b)
			ir := int(255.99 * col.X())
			ig := int(255.99 * col.Y())
			ib := int(255.99 * col.Z())
			line = fmt.Sprintf("%d %d %d\n", ir, ig, ib)
			f.WriteString(line)
		}
	}

	f.Sync()

}

func main() {
	ch1()
	ch4()
}
