// Package main creates an image, gif or video.
package main

import (
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	"github.com/bit101/blgg"
	"github.com/bit101/blgg/render"
	"github.com/bit101/sketch"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomono"
)

func main() {
	target := render.GifTarget

	switch target {
	case render.ImageTarget:
		render.Image(800, 800, "out.png", renderFrame, 0.5)
		render.ViewImage("out.png")
		break

	case render.GifTarget:
		render.Frames(800, 800, 10, "frames", renderFrame)
		render.MakeGIF("ffmpeg", "frames", "out.gif", 15)
		render.ViewImage("out.gif")
		break

	case render.VideoTarget:
		render.Frames(1280, 800, 60, "frames", renderFrame)
		render.ConvertToYoutube("frames", "out.mp4", 60)
		render.VLC("out.mp4", true)
		break
	}
}

func renderFrame(context *blgg.Context, width, height, percent float64) {
	random.RandSeed()
	context.BlackOnWhite()
	s := sketch.FromContext(context)

	// s.SegmentSize = 50
	// s.Shake = 20
	s.FillCircle(400, 400, 300)

	s.SetLineWidth(0.5)
	s.SetWhite()
	font, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	context.SetFontFace(face)
	s.DrawString("Hello world", 280, 640)

	r := 5.0
	s.SegmentSize = 30
	num := 5

	s.StrokeMultiRect(250, 350, 200, 200, r, num)
	s.StrokeMultiRect(350, 250, 200, 200, r, num)

	s.StrokeMultiLine(250, 350, 350, 250, r, num)
	s.StrokeMultiLine(450, 350, 550, 250, r, num)
	s.StrokeMultiLine(450, 550, 550, 450, r, num)
	s.StrokeMultiLine(250, 550, 350, 450, r, num)

	num = 3
	s.SetHSV(0, 0.25, 1)
	s.FillRectangle(10, 10, 100, 100)
	s.SetBlack()
	s.StrokeMultiRect(10, 10, 100, 100, r, num)

	s.SetHSV(70, 0.25, 1)
	s.FillRectangle(690, 10, 100, 100)
	s.SetBlack()
	s.StrokeMultiRect(690, 10, 100, 100, r, num)

	s.SetHSV(140, 0.25, 1)
	s.FillRectangle(690, 690, 100, 100)
	s.SetBlack()
	s.StrokeMultiRect(690, 690, 100, 100, r, num)

	s.SetHSV(210, 0.25, 1)
	s.FillRectangle(10, 690, 100, 100)
	s.SetBlack()
	s.StrokeMultiRect(10, 690, 100, 100, r, num)

	s.SetRGB(1, 0.5, 0)
	s.SegmentSize = 20
	points := []*geom.Point{
		geom.NewPoint(200, 20),
		geom.NewPoint(600, 20),
		geom.NewPoint(650, 100),
		geom.NewPoint(600, 50),
		geom.NewPoint(200, 50),
		geom.NewPoint(150, 100),
		geom.NewPoint(200, 20),
	}
	s.FillPath(points, true)

	s.SetBlack()
	s.StrokeMultiPath(points, true, 3, 3)
}
