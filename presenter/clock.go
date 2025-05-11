package presenter

import (
	"bytes"
	"github.com/golang/freetype/truetype"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"image/png"
	"log/slog"
	"net/http"
	"os"
)

type ClockHandler struct {
	config config.FontConfiguration
	font   *truetype.Font
	clock  domain.Clock
}

func NewClockHandler(
	c *config.FontConfiguration,
	clock domain.Clock,
) *ClockHandler {
	fontFile, err := os.ReadFile(c.DosisFontPath)
	if err != nil {
		slog.Error("NewClockHandler font loading error", "err", err)
		panic(err)
	}
	f, err := truetype.Parse(fontFile)

	return &ClockHandler{
		font:  f,
		clock: clock,
	}
}

func (ch *ClockHandler) Handle(w http.ResponseWriter, r *http.Request) {
	buf, err := ch.generatePNG()
	if err != nil {
		slog.Error("failed to create PNG", "err", err)
		http.Error(w, "ServerError", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/png")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		slog.Error("[ClockHandler.Handle] failed to write image to output", "err", err)
	}
}

func (ch *ClockHandler) generatePNG() (bytes.Buffer, error) {
	now := ch.clock.Now()
	colors := CalculateColors(now)

	img := image.NewGray(image.Rect(0, 0, Height, Width))
	draw.Draw(img, img.Bounds(), &image.Uniform{colors.Bg}, image.Point{}, draw.Src)

	face := truetype.NewFace(ch.font, &truetype.Options{
		Size: 140,
	})
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(colors.Text),
		Face: face,
		Dot:  fixed.P(0, 150),
	}
	DrawStringCentering(d, Height, now.Format("Mon, 02 Jan 2006"))

	face = truetype.NewFace(ch.font, &truetype.Options{
		Size: 470,
	})
	d = &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(colors.Text),
		Face: face,
		Dot:  fixed.P(0, 650),
	}
	DrawStringCentering(d, Height, now.Format("15:04"))

	result := rotate90(img)

	var buf bytes.Buffer
	err := png.Encode(&buf, result)
	return buf, err
}

func rotate90(src image.Image) image.Image {
	srcBounds := src.Bounds()
	w, h := srcBounds.Dx(), srcBounds.Dy()
	dst := image.NewGray(image.Rect(0, 0, h, w))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dst.Set(y, w-1-x, src.At(srcBounds.Min.X+x, srcBounds.Min.Y+y))
		}
	}
	return dst
}
