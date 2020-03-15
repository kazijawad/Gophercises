package main

import (
	"flag"
	"fmt"
	"image/color"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFOption func(*gofpdf.Fpdf)

func FillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}

func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}

type PDF struct {
	fpdf *gofpdf.Fpdf
	x, y float64
}

func (p *PDF) Move(xDelta, yDelta float64) {
	p.x, p.y = p.x+xDelta, p.y+yDelta
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) MoveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) Text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

func (p *PDF) Polygon(pts []gofpdf.PointType, opts ...PDFOption) {
	for _, opt := range opts {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {
	name := flag.String("name", "", "the name of the person who completed the course")
	flag.Parse()

	fpdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := fpdf.GetPageSize()
	fpdf.AddPage()
	pdf := PDF{
		fpdf: fpdf,
		x:    0,
		y:    0,
	}

	primary := color.RGBA{103, 60, 79, 255}
	secondary := color.RGBA{103, 60, 79, 220}

	// Top/Bottom Graphics
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{0, h / 9.0},
		{w - w/6.0, 0},
	}, FillColor(secondary))
	pdf.Polygon([]gofpdf.PointType{
		{w / 6, 0},
		{w, 0},
		{w, h / 9.0},
	}, FillColor(primary))
	pdf.Polygon([]gofpdf.PointType{
		{w, h},
		{w, h - h/8.0},
		{w / 6, h},
	}, FillColor(secondary))
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - h/8.0},
		{w - w/6, h},
	}, FillColor(primary))

	fpdf.SetFont("Times", "B", 50)
	fpdf.SetTextColor(50, 50, 50)
	pdf.MoveAbs(0, 100)
	_, lineHt := fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "Certificate of Completion", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("Arial", "", 28)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("Times", "B", 42)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, *name, gofpdf.AlignCenter)
	pdf.Move(0, lineHt*1.5)

	fpdf.SetFont("Arial", "", 22)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt*1.5, "For successfully completing all twenty programming exercises in the Gophercises programming course for budding Gophers (Go Developers)", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*4.5)

	pdf.Move(0, 65.0)
	fpdf.SetFillColor(100, 100, 100)
	fpdf.Rect(60.0, pdf.y, 240.0, 1.0, "F")
	fpdf.Rect(490.0, pdf.y, 240.0, 1.0, "F")

	fpdf.SetFont("Arial", "", 12)
	pdf.Move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.MoveAbs(60.0+105.0, pdf.y)
	pdf.Text("Date")
	pdf.MoveAbs(490.0+60.0, pdf.y)
	pdf.Text("Instructor - Jon Calhoun")
	pdf.MoveAbs(60.0, pdf.y-lineHt/1.5)
	fpdf.SetFont("Times", "", 22)
	_, lineHt = fpdf.GetFontSize()
	pdf.Move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	yr, mo, dy := time.Now().Date()
	dateStr := fmt.Sprintf("%d/%d/%d", mo, dy, yr)
	fpdf.CellFormat(240.0, lineHt, dateStr, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	pdf.MoveAbs(490.0, pdf.y)

	sig, err := gofpdf.SVGBasicFileParse("sig.svg")
	if err != nil {
		panic(err)
	}
	pdf.Move(0, -(0.45*sig.Ht - lineHt))
	fpdf.SVGBasicWrite(&sig, 0.5)

	err = fpdf.OutputFileAndClose("certificate.pdf")
	if err != nil {
		panic(err)
	}
}
