// Lissajous generates gif animatinos of reandome Lissajous figures
package main

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// array of colors
var palette = []color.Color{color.Black, color.RGBA{0x00, 0x80, 0x00, 0xFF}, color.RGBA{0xFF, 0xFF, 0x00, 0xFF}}

const (
	bgIndex  = 0
	fgIndex  = 1
	midIndex = 2
	cycles   = 5     // number of x revolutions
	res      = 0.001 // angular resolution
	size     = 100   // image canvas from -size to +size, e.g. -100..0..+100
	nframes  = 64    // # frames
	delay    = 8     // delay between frames in 10ms units
)

func main_lissajous() {

	generateLissajous(os.Args[1])
}

func generateLissajous(outFile string) {
	f, _ := os.Create(outFile)

	// ignore handling for now

	w := bufio.NewWriter(f)

	lissajous(w)

	w.Flush()
}

func lissajous(out io.Writer) {

	freq := rand.Float64() * 3.0 // frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := getPix(math.Sin(t))
			y := getPix(math.Sin(t*freq + phase))
			index := getIndex(x, y)
			img.SetColorIndex(x, y, index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)

}

func getIndex(x int, y int) uint8 {
	if math.Sqrt(float64((x-size)*(x-size)+(y-size)*(y-size))) > size {
		return midIndex
	}
	return fgIndex
}

// Turn a value from [-1,1] into a pixel x or y dimension
func getPix(frac float64) int {
	return size + int(frac*size+0.5)
}
