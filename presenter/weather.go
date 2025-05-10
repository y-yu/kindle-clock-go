package presenter

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log/slog"
	"os"
	"sync"
)

type weather struct {
	sunny          *image.Gray
	sunnyWithCloud *image.Gray
	cloudyWithSun  *image.Gray
	cloudy         *image.Gray
	rainyWithSun   *image.Gray
	rainy          *image.Gray
	snowing        *image.Gray
	thunder        *image.Gray
	mist           *image.Gray
}

var weatherIcons = sync.OnceValue(func() weather {
	allWeatherIconFiles := []string{
		"sunny.png",
		"sunny_with_cloud.png",
		"cloudy_with_sun.png",
		"cloudy.png",
		"rainy_with_sun.png",
		"rainy.png",
		"thunder.png",
		"snowing.png",
		"mist.png",
	}
	icons := lo.Map(allWeatherIconFiles, func(fileName string, _ int) *image.Gray {
		f, err := os.Open(fmt.Sprintf("./etc/weather_icon/%s", fileName))
		if err != nil {
			slog.Error("Failed to open weather icon file", "fileName", fileName, "err", err)
			panic(err)
		}
		defer func(fileName string) {
			err := f.Close()
			if err != nil {
				slog.Error("Failed to close weather icon file", "fileName", fileName, "err", err)
			}
		}(fileName)
		img, err := png.Decode(f)
		if err != nil {
			slog.Error("Failed to decode weather icon file", "fileName", fileName, "err", err)
			panic(err)
		}
		gray := image.NewGray(img.Bounds())
		draw.Draw(gray, gray.Bounds(), &image.Uniform{color.White}, img.Bounds().Min, draw.Src)
		draw.Draw(gray, gray.Bounds(), img, img.Bounds().Min, draw.Over)
		return gray
	})
	return weather{
		sunny:          icons[0],
		sunnyWithCloud: icons[1],
		cloudyWithSun:  icons[2],
		cloudy:         icons[3],
		rainyWithSun:   icons[4],
		rainy:          icons[5],
		thunder:        icons[6],
		snowing:        icons[7],
		mist:           icons[8],
	}
})

func ConvertToIcon(iconString string) (*image.Gray, error) {
	icons := weatherIcons()
	switch iconString {
	case "01d", "01n":
		return icons.sunny, nil
	case "02d", "02n":
		return icons.sunnyWithCloud, nil
	case "03d", "03n":
		return icons.cloudyWithSun, nil
	case "04d", "04n":
		return icons.cloudy, nil
	case "09d", "09n":
		return icons.rainyWithSun, nil
	case "10d", "10n":
		return icons.rainy, nil
	case "11d", "11n":
		return icons.thunder, nil
	case "13d", "13n":
		return icons.snowing, nil
	case "50d", "50n":
		return icons.mist, nil
	default:
		return nil, errors.New("unknown icon format")
	}
}
