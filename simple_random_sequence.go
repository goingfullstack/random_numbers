package main

import "fmt"
import "math/rand"

const sides = 6

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(sides) + 1)
	}
}
