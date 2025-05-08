package presenter

import (
	"bytes"
	"context"
	"github.com/golang/freetype/truetype"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/domain/usecase"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log/slog"
	"net/http"
	"os"
)

type RoomInfoHandler struct {
	authConfig      config.AuthenticationConfiguration
	font            *truetype.Font
	roomInfoUsecase usecase.GetRoomInfoUsecase
	clock           domain.Clock
}

func NewRoomInfoHandler(
	ctx context.Context,
	roomInfoUsecase usecase.GetRoomInfoUsecase,
	clock domain.Clock,
) *RoomInfoHandler {
	var c config.AuthenticationConfiguration
	if err := envconfig.Process(ctx, &c); err != nil {
		slog.Error("failed to process configuration for NewRoomInfoHandler", "err", err)
	}

	var f config.FontConfiguration
	if err := envconfig.Process(ctx, &f); err != nil {
		slog.Error("failed to process configuration for NewRoomInfoHandler", "err", err)
	}
	fontFile, err := os.ReadFile(f.RobotoSlabPath)
	if err != nil {
		slog.Error("NewRoomInfoHandler failed to open font file", "f", f, "file", f.RobotoSlabPath, "err", err)
		panic(err)
	}
	ft, err := truetype.Parse(fontFile)
	if err != nil {
		slog.Error("NewRoomInfoHandler font loading error", "err", err)
		panic(err)
	}

	return &RoomInfoHandler{
		authConfig:      c,
		font:            ft,
		roomInfoUsecase: roomInfoUsecase,
		clock:           clock,
	}
}

func (h *RoomInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if h.authConfig.Token != "" {
		if token := r.URL.Query().Get(h.authConfig.QueryKeyName); token != h.authConfig.Token {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	/*ctx := r.Context()
	roomInfo, err := h.roomInfoUsecase.Execute(ctx)
	if err != nil {
		slog.Error("RoomInfoUsecase.Execute failed", "err", err)
		w.Write([]byte("error!"))
	}*/
	roomInfo := usecase.AllRoomInfo{}
	svg, err := h.generatePNG(roomInfo) //GeneratePNGImage(roomInfo, h.clock.Now())
	if err != nil {
		slog.Error("GeneratePNGImage failed", "err", err)
		w.Write([]byte("error!"))
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(svg.Bytes())
}

func (h *RoomInfoHandler) generatePNG(roomInfo usecase.AllRoomInfo) (bytes.Buffer, error) {
	//now := h.clock.Now()
	//colors := CalculateColors(now)

	var buf bytes.Buffer
	weatherIconSrc, err := ConvertToIcon("01d") // roomInfo.Weather.Icon)
	if err != nil {
		return buf, err
	}
	weatherIcon := weatherIconSrc
	weatherIconInv := invertGray(weatherIconSrc)
	/*
		if colors.Bg == color.Black {
			weatherIcon = invertGray(weatherIconSrc)
		}*/

	img := image.NewGray(image.Rect(0, 0, Width, Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	rect := image.Rectangle{Min: image.Point{}, Max: image.Point{}.Add(weatherIcon.Bounds().Size())}
	draw.Draw(img, rect, weatherIcon, weatherIcon.Bounds().Min, draw.Over)

	rect2 := image.Rectangle{Min: image.Point{300, 300}, Max: image.Point{300, 300}.Add(weatherIcon.Bounds().Size())}
	draw.Draw(img, rect2, weatherIconInv, weatherIconInv.Bounds().Min, draw.Over)
	/*
		face := truetype.NewFace(h.font, &truetype.Options{
			Size: 140,
		})
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(colors.Text),
			Face: face,
			Dot:  fixed.P(50, 150),
		}*/

	err = png.Encode(&buf, img)
	return buf, err
}

func invertGray(src *image.Gray) *image.Gray {
	bounds := src.Bounds()
	dst := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			orig := src.GrayAt(x, y).Y
			inv := 255 - orig
			dst.SetGray(x, y, color.Gray{Y: inv})
		}
	}
	return dst
}
