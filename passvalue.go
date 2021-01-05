package main

import (
	"fmt"
	"time"
)

type vertex struct {
	X int
	Y int
}

func (v vertex) mutate(x, y int) vertex {
	v.X = x
	v.Y = y
	return v
}

func (v *vertex) mutateSelf(x, y int) {
	(*v).X = x
	(*v).Y = y
}

func main() {
	limit := 100

	avg := 0
	for i := 0; i < limit; i++ {
		t := time.Now()
		v := vertex{X: 1, Y: 2}
		v = v.mutate(10, 100)
		avg += int(time.Since(t))
	}
	fmt.Println("Return value test average:", avg/limit)

	avg = 0
	for i := 0; i < limit; i++ {
		t := time.Now()
		v := vertex{X: 1, Y: 2}
		v.mutateSelf(10, 100)
		avg += int(time.Since(t))
	}
	fmt.Println("Pointer value test average:", avg/limit)
}
