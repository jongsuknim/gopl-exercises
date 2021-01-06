package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := y1
		if py != height-1 {
			y2 = float64(py+1)/height*(ymax-ymin) + ymin
		}

		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := x1
			if px != width-1 {
				x2 = float64(px+1)/width*(xmax-xmin) + xmin
			}
			z := (complex(x1, y1) + complex(x1, y2) + complex(x2, y1) + +complex(x2, y2)) / 4
			// Image의 점 (px, py)는 복소수 값 z를 나타낸다.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
