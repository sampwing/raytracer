package main

import (
	"fmt"
	"log"
	"os"
)

func drawPPM() {
	filename := "output.ppm"
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
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			line = fmt.Sprintf("%d %d %d\n", ir, ig, ib)
			f.WriteString(line)
		}
	}

	log.Printf("%d %d", nx, ny)
	f.Sync()
}

func main() {
	drawPPM()
}
