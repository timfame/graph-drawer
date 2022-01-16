package draw

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type std struct {
	image *image.RGBA
	color color.Color
}

func NewSTD() *std {
	log.Println("Text drawing is not possible in std visualiser")
	return &std{
		color: color.Black,
	}
}

func (s *std) NewImage(width, height int) {
	s.image = image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s.image.Set(x, y, color.White)
		}
	}
}

func (s *std) DrawCircle(x0, y0, r int) error {
	if err := s.checkImage(); err != nil {
		return err
	}
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)
	for x > y {
		s.image.Set(x0+x, y0+y, s.color)
		s.image.Set(x0+y, y0+x, s.color)
		s.image.Set(x0-y, y0+x, s.color)
		s.image.Set(x0-x, y0+y, s.color)
		s.image.Set(x0-x, y0-y, s.color)
		s.image.Set(x0-y, y0-x, s.color)
		s.image.Set(x0+y, y0-x, s.color)
		s.image.Set(x0+x, y0-y, s.color)
		if err <= 0 {
			y++
			err += dy
			dy++
		}
		if err > 0 {
			x--
			dx++
			err += dx - (r * 2)
		}
	}
	return nil
}

func (s *std) DrawLine(x1, y1, x2, y2 int) error {
	if err := s.checkImage(); err != nil {
		return err
	}
	if y1 > y2 {
		y1, y2 = y2, y1
		x1, x2 = x2, x1
	}
	for y := y1; y <= y2; y++ {
		x := int(float64(x1) + float64(x2 - x1) * (float64(y - y1) / float64(y2 - y1)))
		s.image.Set(x, y, s.color)
	}
	return nil
}

func (s *std) DrawText(x, y int, text string) error {
	if err := s.checkImage(); err != nil {
		return err
	}
	return nil
}

func (s *std) SaveImage(path string) error {
	if err := s.checkImage(); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, s.image)
}

func (s *std) checkImage() error {
	if s.image == nil {
		return fmt.Errorf("new image is not created")
	}
	return nil
}
