package clock

import (
	"fmt"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/presenter"
	"log/slog"
	"net/http"
)

const svgStringFormat = `<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
	<style>
		text { fill: %s; }
		.imageColor { fill: %s; }
	</style>

	<title>Kindle Clock</title>
	<g transform="%s" font-family="Dosis">
		<text font-size="140px" y="300" x="380" text-anchor="middle">
			%s
		</text>

		<text font-size="470" y="790" x="380" text-anchor="middle">
			<tspan>%s</tspan>
			<tspan dy="-0.15em">:</tspan>
			<tspan dy="0.15em">%s</tspan>
		</text>
	</g>
</svg>`

type ClockHandler struct {
	clock domain.Clock
}

func NewClockHandler(clock domain.Clock) *ClockHandler {
	return &ClockHandler{
		clock: clock,
	}
}

func (ch *ClockHandler) Handle(w http.ResponseWriter, r *http.Request) {
	now := ch.clock.Now()
	colors := presenter.CalculateColors(now)
	transform := fmt.Sprintf("rotate(-90, %d, %d)", presenter.Width/2, presenter.Height/2)
	svgString := fmt.Sprintf(
		svgStringFormat,
		presenter.Width,
		presenter.Height,
		colors.Text,
		colors.Text,
		transform,
		now.Format("Mon, 02 Jan 2006"),
		now.Format("15"),
		now.Format("04"),
	)
	buf, err := presenter.ConvertSVGToPNG(svgString, colors.Bg, presenter.Width, presenter.Height)
	if err != nil {
		slog.Error("failed to convert SVG to PNG", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}
