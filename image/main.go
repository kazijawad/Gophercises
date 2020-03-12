package main

import (
	"math/rand"
	"os"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 171},
		{"Feb", 180},
		{"Mar", 100},
		{"Apr", 87},
		{"May", 66},
		{"Jun", 40},
		{"Jul", 32},
		{"Aug", 55},
	}
	w := len(data)*60 + 10
	h := 300
	threshold := 120
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(w, h)
	for i, val := range data {
		p := val.Usage * (h - 50) / max
		canvas.Rect(i*60+10, (h-50)-p, 50, p, "fill:rgb(77,117,232)")
		canvas.Text(i*60+35, h-25, val.Month, "font-size:14pt;fill:rgb(150,150,150);text-anchor:middle")
	}
	threshP := threshold * (h - 50) / max
	canvas.Line(0, h-threshP, w, h-threshP, "stroke:rgb(255,100,100);opacity:0.8;stroke-width:2")
	canvas.Rect(0, 0, w, h-threshP, "fill:rgb(255,100,100);opacity:0.2")
	canvas.Line(0, h-50, w, h-50, "stroke:rgb(150,150,150);stroke-width:2")
	canvas.End()
}
