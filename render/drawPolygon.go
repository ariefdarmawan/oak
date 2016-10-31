package render

import (
	clip "github.com/akavel/polyclip-go"
)

var (
	usingDrawPolygon = false
	drawPolygon      clip.Polygon
)

func SetDrawPolygon(p clip.Polygon) {
	usingDrawPolygon = true
	drawPolygon = p
}

func InDrawPolygon(xi, yi, x2i, y2i int) bool {
	x := float64(xi)
	y := float64(yi)
	x2 := float64(x2i)
	y2 := float64(y2i)
	if usingDrawPolygon {
		p2 := clip.Polygon{{{x, y}, {x, y2}, {x2, y2}, {x2, y}}}
		intsct := drawPolygon.Construct(clip.INTERSECTION, p2)
		return len(intsct) != 0
	}
	return true
}