package main

import (
	"fmt"
	"math"
	"sync"
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

func main() {
	draw1()
	draw2()
	draw3()
}

func draw1() string {
	r := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			r += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	r += fmt.Sprintf("</svg>")

	return r
}

type draw struct {
	i    int
	j    int
	line string
}

func draw2() string {
	r := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	ch := make(chan draw, cells)
	var wg sync.WaitGroup

	for i := 0; i < cells; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				var dr draw
				dr.i, dr.j, dr.line = i, j, fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				ch <- dr
			}
		}(i)
	}

	go func() {
		wg.Wait() // wait for every worker to finish
		close(ch) // close the channel
	}()

	lines := make([][]string, cells)
	for i := range lines {
		lines[i] = make([]string, cells)
	}

	for d := range ch {
		lines[d.i][d.j] = d.line
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			r += lines[i][j]
		}
	}

	r += "</svg>"

	return r
}

func draw3() string {
	r := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	lines := make([][]string, cells)
	for i := range lines {
		lines[i] = make([]string, cells)
	}

	goroutine := 10
	var wg sync.WaitGroup

	for gr := 0; gr < goroutine; gr++ {
		i := gr * 10
		wg.Add(1)
		go func(i, limit int) {
			defer wg.Done()
			for ; i < limit; i++ {
				for j := 0; j < cells; j++ {
					ax, ay := corner(i+1, j)
					bx, by := corner(i, j)
					cx, cy := corner(i, j+1)
					dx, dy := corner(i+1, j+1)
					lines[i][j] = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
		}(i, (cells/goroutine)+i)
	}

	wg.Wait()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			r += lines[i][j]
		}
	}

	r += "</svg>"

	return r
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
