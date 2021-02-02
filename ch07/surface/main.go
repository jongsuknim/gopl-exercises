package main

import (
	"fmt"
	"github/jongsuknim/gopl-exercises/ch07/eval"
	"io"
	"log"
	"math"
	"net/http"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env(eval.Env{"x": x, "y": y, "r": r}))
	})
}

const (
	width, height = 600, 320            // 픽셀 단위 캔버스 크기
	cells         = 100                 // 격자 셀의 숫자
	xyrange       = 30.0                // 축 범위 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x나 y의 단위 픽셀
	zscale        = height * 0.4        // z 단위 픽셀
	angle         = math.Pi / 6         // x, y축의 각도 (=30도)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	// (i, j) 셀 코너에서 (x, y) 점 찾기
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 표면 높이 z 연산
	z := f(x, y)

	// (x,y,z)를 3차원 SVG 평면 (sx, sy)에 등각 투영시킴
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func main() {
	http.HandleFunc("/", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
