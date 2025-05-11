package presenter

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/sethvargo/go-envconfig"
	"github.com/y-yu/kindle-clock-go/config"
	"github.com/y-yu/kindle-clock-go/domain"
	"github.com/y-yu/kindle-clock-go/domain/model"
	"github.com/y-yu/kindle-clock-go/domain/usecase"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"time"
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
			slog.Warn("invalid token", "token", token)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	ctx := r.Context()
	roomInfo, err := h.roomInfoUsecase.Execute(ctx)
	if err != nil {
		slog.Error("RoomInfoUsecase.Execute failed", "err", err)
		http.Error(w, "ServerError", http.StatusInternalServerError)
		return
	}
	svg, err := h.generatePNG(roomInfo)
	if err != nil {
		slog.Error("GeneratePNGImage failed", "err", err)
		http.Error(w, "ServerError", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(svg.Bytes())
	if err != nil {
		slog.Error("[RoomInfoHandler#Handle] failed to write image", "err", err)
	}
}

func (h *RoomInfoHandler) generatePNG(roomInfo usecase.AllRoomInfo) (bytes.Buffer, error) {
	now := h.clock.Now()
	colors := CalculateColors(now)

	var buf bytes.Buffer

	textUIFace := truetype.NewFace(h.font, &truetype.Options{
		Size: 30,
	})
	uiFace := truetype.NewFace(h.font, &truetype.Options{
		Size: 90,
	})
	smallTextUIFace := truetype.NewFace(h.font, &truetype.Options{
		Size: 20,
	})
	smallUIFace := truetype.NewFace(h.font, &truetype.Options{
		Size: 50,
	})
	defer func() {
		err := textUIFace.Close()
		if err != nil {
			slog.Error("textUIFace.Close()", "err", err)
		}
		err = smallTextUIFace.Close()
		if err != nil {
			slog.Error("smallTextUIFace.Close()", "err", err)
		}
		err = uiFace.Close()
		if err != nil {
			slog.Error("uiFace.Close()", "err", err)
		}
		err = smallUIFace.Close()
		if err != nil {
			slog.Error("smallUIFace.Close()", "err", err)
		}
	}()

	img := image.NewGray(image.Rect(0, 0, Width, Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{colors.Bg}, image.Point{}, draw.Src)

	err := weatherIcon(img, colors, smallTextUIFace, roomInfo.Weather.Icon, now)
	if err != nil {
		return buf, err
	}

	clock(img, colors, textUIFace, now)

	score(img, colors.Text, textUIFace, uiFace, roomInfo.AwairRoomInfo.Score)
	temperatureHumidityTable(
		img,
		colors.Text,
		textUIFace,
		uiFace,
		[3]model.Temperature{roomInfo.AwairRoomInfo.Temperature, roomInfo.NatureRemoRoomInfo.Temperature, roomInfo.SwitchBotMeterInfo.Temperature},
		[3]model.Humidity{roomInfo.AwairRoomInfo.Humidity, roomInfo.NatureRemoRoomInfo.Humidity, roomInfo.SwitchBotMeterInfo.Humidity},
	)
	electricEnergy(img, colors.Text, textUIFace, uiFace, roomInfo.NatureRemoRoomInfo.ElectricEnergy)
	airInfoTable(img, colors.Text, textUIFace, smallUIFace, smallTextUIFace, roomInfo.AwairRoomInfo)

	err = png.Encode(&buf, img)
	return buf, err
}

func weatherIcon(
	img draw.Image,
	colors Colors,
	smallTextUIFace font.Face,
	icon string,
	now time.Time,
) error {
	weatherIconSrc, err := ConvertToIcon(icon)
	if err != nil {
		return err
	}
	weatherIcon := weatherIconSrc
	if colors.Bg == color.Black {
		weatherIcon = invertGray(weatherIconSrc)
	}
	initialPoint := image.Point{20, 0}
	rect := image.Rectangle{Min: initialPoint, Max: initialPoint.Add(weatherIcon.Bounds().Size())}
	draw.Draw(img, rect, weatherIcon, weatherIcon.Bounds().Min, draw.Over)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(colors.Text),
		Face: smallTextUIFace,
		Dot:  fixed.P(320, 250),
	}
	d.DrawString(now.Format(time.RFC3339))

	return nil
}

func clock(
	img draw.Image,
	colors Colors,
	textUIFace font.Face,
	now time.Time,
) {
	const (
		clockReactX      = 512
		clockReactLength = 86
	)
	rect := image.Rect(510, 30, 600, 80)
	draw.Draw(img, rect, &image.Uniform{colors.Text}, rect.Bounds().Min, draw.Over)
	rect = image.Rect(clockReactX, 32, clockReactX+clockReactLength, 78)
	draw.Draw(img, rect, &image.Uniform{colors.Bg}, rect.Bounds().Min, draw.Over)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(colors.Text),
		Face: textUIFace,
		Dot:  fixed.P(clockReactX, 65),
	}
	DrawStringCentering(d, clockReactLength, now.Format("15:04"))
}

func score(
	img draw.Image,
	textColor color.Color,
	textUIFace font.Face,
	uiFace font.Face,
	score model.Score,
) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: textUIFace,
		Dot:  fixed.P(430, 150),
	}
	d.DrawString("Score")
	d.Dot.X += fixed.I(30)
	d.Dot.Y += fixed.I(60)
	d.Face = uiFace
	d.DrawString(fmt.Sprintf("%d", score))
}

func temperatureHumidityTable(
	img draw.Image,
	textColor color.Color,
	textUIFace font.Face,
	uiFace font.Face,
	temperature [3]model.Temperature,
	humidity [3]model.Humidity,
) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: textUIFace,
		Dot:  fixed.P(200, 340),
	}
	d.DrawString("Temperature")
	d.Dot.X += fixed.I(150)
	d.DrawString("Humidity")

	d.Dot.X = fixed.I(20)
	d.Dot.Y += fixed.I(70)

	d.DrawString("AWAIR")
	d.Dot.X = fixed.I(20)
	d.Dot.Y += fixed.I(110)
	d.DrawString("Nature")
	d.Dot.X = fixed.I(40)
	d.Dot.Y += fixed.I(30)
	d.DrawString("Remo")
	d.Dot.X = fixed.I(20)
	d.Dot.Y += fixed.I(100)
	d.DrawString("SwitchBot")

	tableY := 430

	d.Dot = fixed.P(210, tableY)
	for _, te := range temperature {
		d.Face = uiFace
		d.DrawString(fmt.Sprintf("%s", floatSawedOffString(te)))
		d.Face = textUIFace
		d.DrawString("°C")
		d.Dot.X = fixed.I(210)
		d.Dot.Y += fixed.I(120)
	}

	d.Dot = fixed.P(520, tableY)
	for _, hu := range humidity {
		d.Face = uiFace
		d.DrawString(fmt.Sprintf("%s", floatSawedOffString(hu)))
		d.Face = textUIFace
		d.DrawString("%")
		d.Dot.X = fixed.I(520)
		d.Dot.Y += fixed.I(120)
	}
}

func electricEnergy(
	img draw.Image,
	textColor color.Color,
	textUIFace font.Face,
	uiFace font.Face,
	energy model.ElectricEnergy,
) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: textUIFace,
		Dot:  fixed.P(150, 750),
	}
	d.DrawString("ElectricEnergy")
	d.Dot.X += fixed.I(30)
	d.Dot.Y += fixed.I(80)
	d.Face = uiFace
	d.DrawString(fmt.Sprintf("%d", energy))
	d.Face = textUIFace
	d.DrawString("W")
}

func airInfoTable(
	img draw.Image,
	textColor color.Color,
	textUIFace font.Face,
	smallUIFace font.Face,
	smallTextUIFace font.Face,
	info model.AwairRoomInfo,
) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: textUIFace,
		Dot:  fixed.P(40, 900),
	}
	d.DrawString("CO")
	d.Face = smallTextUIFace
	d.DrawString("2")
	d.Dot.X += fixed.I(200)
	d.Face = textUIFace
	d.DrawString("VOC")
	d.Dot.X += fixed.I(210)
	d.DrawString("PM2.5")

	d.Dot = fixed.P(60, 960)
	d.Face = smallUIFace
	d.DrawString(fmt.Sprintf("%d", info.Co2))
	d.Face = smallTextUIFace
	d.DrawString("ppm")
	d.Dot.X += fixed.I(150)
	d.Face = smallUIFace
	d.DrawString(fmt.Sprintf("%d", info.Voc))
	d.Face = smallTextUIFace
	d.DrawString("ppb")
	d.Dot.X += fixed.I(160)
	d.Face = smallUIFace
	d.DrawString(fmt.Sprintf("%s", floatSawedOffString(info.Pm25)))
	d.Face = smallTextUIFace
	d.DrawString("μg/m³")
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

func floatSawedOffString[A ~float32](d A) string {
	return fmt.Sprintf("%2.1f", d)
}
