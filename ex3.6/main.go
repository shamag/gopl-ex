package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x04, 0x0f, 0xff},
	color.RGBA{0x03, 0x26, 0x28, 0xff},
	color.RGBA{0x07, 0x3e, 0x1e, 0xff},
	color.RGBA{0x18, 0x55, 0x08, 0xff},
	color.RGBA{0x5f, 0x6e, 0x0f, 0xff},
	color.RGBA{0x84, 0x50, 0x19, 0xff},
	color.RGBA{0x9b, 0x30, 0x22, 0xff},
	color.RGBA{0xb4, 0x92, 0x2f, 0xff},
	color.RGBA{0x94, 0xca, 0x3d, 0xff},
	color.RGBA{0x4f, 0xd5, 0x51, 0xff},
	color.RGBA{0x66, 0xff, 0xb3, 0xff},
	color.RGBA{0x82, 0xc9, 0xe5, 0xff},
	color.RGBA{0x9d, 0xa3, 0xeb, 0xff},
	color.RGBA{0xd7, 0xb5, 0xf3, 0xff},
	color.RGBA{0xfd, 0xd6, 0xf6, 0xff},
	color.RGBA{0xff, 0xf0, 0xf2, 0xff},
}

func getAvg(a, b, c, d uint32) uint16 {
	return uint16((a + b + c + d) / 4)
}
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	const dx = (xmax - xmin) / float64(width*2)
	const dy = (ymax - ymin) / float64(height*2)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			r1, g1, b1, a1 := mandelbrot(complex(x-dx, y-dy)).RGBA()
			r2, g2, b2, a2 := mandelbrot(complex(x+dx, y-dy)).RGBA()
			r3, g3, b3, a3 := mandelbrot(complex(x-dx, y+dy)).RGBA()
			r4, g4, b4, a4 := mandelbrot(complex(x+dx, y+dy)).RGBA()
			// fmt.Println(x, x-(xmax-xmin)/float64(width))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.RGBA64{getAvg(r1, r2, r3, r4), getAvg(g1, g2, g3, g4), getAvg(b1, b2, b3, b4), getAvg(a1, a2, a3, a4)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%16]
		}
	}
	return palette[15]
}
