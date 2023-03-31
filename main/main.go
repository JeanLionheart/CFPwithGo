package main

import (
	"fmt"
	"gcm/pipeline"
	"math"
	"math/rand"
)

func FindPrime() {
	src := func() any {
		return rand.Int()
	}
	hdl1 := func(n any) any {
		a := n.(int)
		for i := 2; i < int(math.Sqrt(float64(a))); i++ {
			if a%i == 0 {
				return -1
			}
		}
		fmt.Println(a)
		return a
	}
	done := make(chan any)
	g := pipeline.Generate(done, src)
	h := pipeline.Handle(done, g, hdl1)
	for true {
		n := <-h
		if n.(int) == -1 {
			continue
		}
		fmt.Println(n)
	}
}

func main() {
	FindPrime()
}
