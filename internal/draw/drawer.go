package draw

type Drawer interface {
	NewImage(width, height int)
	DrawCircle(x, y, r int) error
	DrawLine(x1, y1, x2, y2 int) error
	DrawText(x, y int, text string) error
	SaveImage(path string) error
}
