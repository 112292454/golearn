// Lissajous generates GIF animations of random Lissajous figures.
package ch1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.White, color.RGBA{0x00, 0xFF, 0, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	f, _ := os.Create("out.gif")
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10      // number of complete x oscillator revolutions
		res     = 0.001   // angular resolution
		size    = 200     // image canvas covers [-size..+size]
		nframes = 128 * 2 // number of animation frames
		delay   = 8       // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//var ctype = uint8(rand.Intn(2)) + 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(int(t)%2)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return
	} // NOTE: ignoring encoding errors
}
