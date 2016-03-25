package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sampwing/raytracer/raytracer"
)

func ch5() {
	filename := "ch5.ppm"
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
			ho := raytracer.MultiplyVectorBy(horizontal, u)
			ve := raytracer.MultiplyVectorBy(vertical, v)
			lhv := raytracer.AddVectors(lowerLeftCorner, ho)
			lhv = raytracer.AddVectors(lhv, ve)

			r := raytracer.NewRay(origin, lhv)

			col := r.Color()
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
	ch5()
}
