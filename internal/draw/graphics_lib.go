package draw

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
)

type graphicsLib struct {
	ctx  *gg.Context
}

func NewGraphicsLib() *graphicsLib {
	return &graphicsLib{}
}

func (g *graphicsLib) NewImage(width, height int) {
	g.ctx = gg.NewContext(width, height)
	g.ctx.SetColor(color.White)
	g.ctx.DrawRectangle(0, 0, float64(width), float64(height))
	g.ctx.Fill()
	g.ctx.SetColor(color.Black)
}

func (g *graphicsLib) DrawCircle(x, y, r int) error {
	if err := g.ctxCheck(); err != nil {
		return err
	}
	g.ctx.DrawCircle(float64(x), float64(y), float64(r))
	g.ctx.Fill()
	return nil
}

func (g *graphicsLib) DrawLine(x1, y1, x2, y2 int) error {
	if err := g.ctxCheck(); err != nil {
		return err
	}
	g.ctx.DrawLine(float64(x1), float64(y1), float64(x2), float64(y2))
	g.ctx.Stroke()
	return nil
}

func (g *graphicsLib) DrawText(x, y int, text string) error {
	if err := g.ctxCheck(); err != nil {
		return err
	}
	g.ctx.SetColor(color.White)
	g.ctx.DrawString(text, float64(x), float64(y))
	g.ctx.Stroke()
	g.ctx.SetColor(color.Black)
	return nil
}

func (g *graphicsLib) SaveImage(path string) error {
	if err := g.ctxCheck(); err != nil {
		return err
	}
	return g.ctx.SavePNG(path)
}

func (g *graphicsLib) ctxCheck() error {
	if g.ctx == nil {
		return fmt.Errorf("new image is not created")
	}
	return nil
}
