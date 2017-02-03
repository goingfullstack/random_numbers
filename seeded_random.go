package main

import "fmt"
import "math/rand"
import "time"

const sides = 6

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(sides) + 1)
}
