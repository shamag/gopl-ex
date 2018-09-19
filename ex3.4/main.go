package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var max = 0.0

func isNan(params ...float64) bool {
	for _, param := range params {
		if math.IsNaN(param) {
			return true
		}
	}
	return false
}
func main() {
	http.HandleFunc("/", svg)
	http.ListenAndServe(":8080", nil)
}
func svg(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, za := corner(i+1, j)
			bx, by, zb := corner(i, j)
			cx, cy, zc := corner(i, j+1)
			dx, dy, zd := corner(i+1, j+1)
			maxab := math.Max(za, zb)
			maxcd := math.Max(zc, zd)
			maxz := (math.Max(maxab, maxcd) + 1) * 127.5
			color := fmt.Sprintf("rgb(%d , 0, %d)", int(maxz), int(255-maxz))
			if isNan(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(rw, "<polygon style='stroke: %s; fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(rw, "</svg>")
}
func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
