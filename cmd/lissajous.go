package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	f, _ := os.Create("out.gif")
	lissajous(f)
}
