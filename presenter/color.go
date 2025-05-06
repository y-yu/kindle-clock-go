package presenter

import "time"

type Colors struct {
	Text string
	Bg   string
}

func CalculateColors(now time.Time) Colors {
	colors := Colors{
		Text: "black",
		Bg:   "white",
	}
	if now.Hour() <= 6 || now.Hour() >= 17 {
		colors = Colors{
			Text: "white",
			Bg:   "black",
		}
	}
	return colors
}
