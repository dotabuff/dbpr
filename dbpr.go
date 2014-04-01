package main

import "math/rand"

func main() {
	for {
		print("nyx ")
		if rand.Float32() < 0.01 {
			print("I TOLD YOU TO STAY IN THE TREES ")
		}
	}
}
