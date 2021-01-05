package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width, height = 600, 320               // 픽셀 단위 캔버스 크기
var cells = 100                            // 격자 셀의 숫자
var xyrange = 30.0                         // 축 범위 (-xyrange..+xyrange)
var xyscale = float64(width) / 2 / xyrange // x나 y의 단위 픽셀
var zscale = float64(height) * 0.4         // z 단위 픽셀
var angle = math.Pi / 6                    // x, y축의 각도 (=30도)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		if v, ok := r.Form["width"]; ok {
			if w, err := strconv.Atoi(v[0]); err == nil {
				width = w
				xyscale = float64(width) / 2 / xyrange
			}
		}

		if v, ok := r.Form["height"]; ok {
			if h, err := strconv.Atoi(v[0]); err == nil {
				height = h
				zscale = float64(height) * 0.4
			}
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintf(w, "</svg>")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64) {
	// (i, j) 셀 코너에서 (x, y) 점 찾기
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// 표면 높이 z 연산
	z := f(x, y)

	// (x,y,z)를 3차원 SVG 평면 (sx, sy)에 등각 투영시킴
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)에서의 거리
	return math.Sin(r) / r
}
