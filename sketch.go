// Package sketch draws sketchy shapes.
package sketch

import (
	"math"

	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	"github.com/bit101/blgg"
)

// Sketch is a drawing context that draws sketchy shapes.
type Sketch struct {
	*blgg.Context
	SegmentSize float64
	Shake       float64
}

// NewSketch creates a new Sketch drawing context.
func NewSketch(w, h float64) *Sketch {
	c := blgg.NewContextF(w, h)
	return FromContext(c)
}

// FromContext creates a sketch from a blgg context.
func FromContext(context *blgg.Context) *Sketch {
	return &Sketch{
		context,
		15,
		5,
	}
}

// MoveTo moves to the given x, y location, sketchily.
func (s *Sketch) MoveTo(x, y float64) {
	d := s.Shake / 2
	s.Context.MoveTo(x+random.FloatRange(-d, d), y+random.FloatRange(-d, d))
}

// LineTo draws a sketchy line to the given x, y location.
func (s *Sketch) LineTo(x, y float64) {
	p, _ := s.GetCurrentPoint()
	d := s.Shake / 2
	endX := x + random.FloatRange(-d, d)
	endY := y + random.FloatRange(-d, d)
	dx := endX - p.X
	dy := endY - p.Y
	dist := math.Hypot(dx, dy)
	steps := math.Floor(dist / s.SegmentSize)
	resX := dx / steps
	resY := dy / steps
	for i := 1.0; i < steps; i++ {
		s.Context.LineTo(p.X+resX*i+random.FloatRange(-d, d), p.Y+resY*i+random.FloatRange(-d, d))
	}
	s.Context.LineTo(endX, endY)
}

// StrokeMultiLine draws multiple, offset stroked lines between two points.
func (s *Sketch) StrokeMultiLine(x0, y0, x1, y1, separation float64, iter int) {
	for i := 0; i < iter; i++ {
		s.MoveTo(x0+random.FloatRange(-separation, separation), y0+random.FloatRange(-separation, separation))
		s.LineTo(x1+random.FloatRange(-separation, separation), y1+random.FloatRange(-separation, separation))
		s.Stroke()
	}
}

// Circle draws a sketchy circle.
func (s *Sketch) Circle(xc, yc, r float64) {
	d := s.Shake / 2
	circ := r * math.Pi * 2
	steps := circ / s.SegmentSize
	res := math.Pi * 2 / steps
	for a := 0.0; a < math.Pi*2; a += res {
		x := xc + math.Cos(a)*r + random.FloatRange(-d, d)
		y := yc + math.Sin(a)*r + random.FloatRange(-d, d)
		s.Context.LineTo(x, y)
	}
	s.ClosePath()
}

// StrokeCircle strokes a sketchy circle.
func (s *Sketch) StrokeCircle(xc, yc, r float64) {
	s.Circle(xc, yc, r)
	s.Stroke()
}

// FillCircle fills a sketchy circle.
func (s *Sketch) FillCircle(xc, yc, r float64) {
	s.Circle(xc, yc, r)
	s.Fill()
}

// Rectangle draws a sketchy rectangle.
func (s *Sketch) Rectangle(x, y, w, h float64) {
	s.MoveTo(x, y)
	s.LineTo(x+w, y)
	s.LineTo(x+w, y+h)
	s.LineTo(x, y+h)
	s.LineTo(x, y)
}

// StrokeRectangle strokes a sketchy rectangle.
func (s *Sketch) StrokeRectangle(x, y, w, h float64) {
	s.Rectangle(x, y, w, h)
	s.Stroke()
}

// FillRectangle fills a sketchy rectangle.
func (s *Sketch) FillRectangle(x, y, w, h float64) {
	s.Rectangle(x, y, w, h)
	s.Fill()
}

// StrokeMultiRect strokes multiple sketchy rectangles
func (s *Sketch) StrokeMultiRect(x, y, w, h, separation float64, iter int) {
	s.StrokeMultiLine(x, y, x+w, y, separation, iter)
	s.StrokeMultiLine(x+w, y, x+w, y+h, separation, iter)
	s.StrokeMultiLine(x+w, y+h, x, y+h, separation, iter)
	s.StrokeMultiLine(x, y+h, x, y, separation, iter)
}

// DrawString draws a sketchy string
func (s *Sketch) DrawString(str string, x, y float64) {
	d := s.Shake / 2
	rot := d * 0.05
	for _, c := range str {
		s.Push()
		s.Translate(x, y)
		s.Rotate(random.FloatRange(-rot, rot))
		s.Context.DrawString(string(c), random.FloatRange(-d, d), random.FloatRange(-d, d))
		w, _ := s.MeasureString(string(c))
		x += w
		s.Pop()
	}
}

// Path draws a sketchy path through a set of points.
func (s *Sketch) Path(points []*geom.Point, closed bool) {
	s.MoveTo(points[0].X, points[0].Y)
	for i := 1; i < len(points); i++ {
		s.LineTo(points[i].X, points[i].Y)
	}
	if closed {
		s.LineTo(points[0].X, points[0].Y)
	}
}

// StrokePath strokes a sketchy path through a set of points.
func (s *Sketch) StrokePath(points []*geom.Point, closed bool) {
	s.Path(points, closed)
	s.Stroke()
}

// FillPath fills a sketchy path through a set of points.
func (s *Sketch) FillPath(points []*geom.Point, closed bool) {
	s.Path(points, closed)
	s.Fill()
}

// StrokeMultiPath strokes a sketchy path through a set of points.
func (s *Sketch) StrokeMultiPath(points []*geom.Point, closed bool, separation float64, iter int) {
	for i := 0; i < iter; i++ {
		dx := random.FloatRange(-separation, separation)
		dy := random.FloatRange(-separation, separation)

		p2 := make([]*geom.Point, len(points))
		for i, p := range points {
			p2[i] = geom.NewPoint(p.X+dx, p.Y+dy)
		}

		s.Push()
		s.StrokePath(p2, closed)
		s.Pop()
	}
}
