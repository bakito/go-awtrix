package types

type Drawable interface {
	Draw() []interface{}
}

var (
	_ Drawable = &DrawPixel{}
	_ Drawable = &DrawLine{}
	_ Drawable = &DrawRectangle{}
	_ Drawable = &DrawCircle{}
	_ Drawable = &DrawText{}
	_ Drawable = &DrawBitmap{}
)

type DrawPixel struct {
	X     int
	Y     int
	Color string
}

func (d *DrawPixel) Draw() []interface{} {
	return []interface{}{d.X, d.Y, d.Color}
}

type DrawLine struct {
	X1    int
	Y1    int
	X2    int
	Y2    int
	Color string
}

func (d *DrawLine) Draw() []interface{} {
	return []interface{}{d.X1, d.Y1, d.X2, d.Y2, d.Color}
}

type DrawRectangle struct {
	X      int
	Y      int
	Width  int
	Height int
	Color  string
}

func (d *DrawRectangle) Draw() []interface{} {
	return []interface{}{d.X, d.Y, d.Width, d.Height, d.Color}
}

type DrawCircle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

func (d *DrawCircle) Draw() []interface{} {
	return []interface{}{d.X, d.Y, d.Radius, d.Color}
}

type DrawText struct {
	X     int
	Y     int
	Text  string
	Color string
}

func (d *DrawText) Draw() []interface{} {
	return []interface{}{d.X, d.Y, d.Text, d.Color}
}

type DrawBitmap struct {
	X      int
	Y      int
	Width  int
	Height int
	Bitmap []string
}

func (d *DrawBitmap) Draw() []interface{} {
	return []interface{}{d.X, d.Y, d.Width, d.Height, d.Bitmap}
}
