package main

import (
	"math"

	"fmt"
)
import "math/rand"
import "time"

const sides = 6
const sampleInterval = 10000

func main() {
	counter := make(map[int]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < math.MaxInt32; i++ {
		num := rand.Intn(sides) + 1
		count := counter[num]
		count++
		counter[num] = count

		if i%sampleInterval == 0 {
			fmt.Printf("%d,", i/sampleInterval)
			fmt.Printf("%f,", float64(counter[1])/float64(i))
			fmt.Printf("%f,", float64(counter[2])/float64(i))
			fmt.Printf("%f,", float64(counter[3])/float64(i))
			fmt.Printf("%f,", float64(counter[4])/float64(i))
			fmt.Printf("%f,", float64(counter[5])/float64(i))
			fmt.Printf("%f\n", float64(counter[6])/float64(i))
		}
	}
}
