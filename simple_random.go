package main

import "fmt"
import "math/rand"

const sides = 6

func main() {
	fmt.Println(rand.Intn(sides-1) + 1)
}
