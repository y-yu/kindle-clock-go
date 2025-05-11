package presenter

import (
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"log/slog"
)

const (
	Width  = 758
	Height = 1024
)

func DrawStringCentering(
	d *font.Drawer,
	width int,
	str string,
) {
	length := d.MeasureString(str)
	// Int26_6 is fixed point number, 26 bits for integer part and 6 bits for fraction part
	// so in division we have to multiply 2^6 (= 64) to left hand side because Int26_6 is shifted 6 bits for fraction part.
	x := (fixed.I(width) - length) * 64 / fixed.I(2)
	slog.Info("DrawStringCentering", "x", x, "int(fixed.I(2))", int(fixed.I(2)), "int(I(width))", int(fixed.I(width)), "I(width)", fixed.I(width), "x", int(x), "length", length, "int(length)", int(length), "str", str)
	d.Dot.X += x
	d.DrawString(str)
}
