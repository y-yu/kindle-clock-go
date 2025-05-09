package presenter

import (
	"image/color"
	"time"
)

type Colors struct {
	Text color.Color
	Bg   color.Color
}

func CalculateColors(now time.Time) Colors {
	colors := Colors{
		Text: color.Black,
		Bg:   color.White,
	}
	if now.Hour() <= 6 || now.Hour() >= 17 {
		colors = Colors{
			Text: color.White,
			Bg:   color.Black,
		}
	}
	return colors
}
