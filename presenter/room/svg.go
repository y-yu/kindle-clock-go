package room

import (
	"bytes"
	"fmt"
	"github.com/y-yu/kindle-clock-go/domain/usecase"
	"github.com/y-yu/kindle-clock-go/presenter"
	"time"
)

const svgStringFormat = `<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
	<style>
		text { fill: %s; }
		.imageColor { fill: %s; }
	</style>
	<title>Kindle Clock</title>
		<g font-family="Roboto Slab">
			<g transform="scale(3) translate(15 0)" class="imageColor">
				<svg height="100" width="100" xmlns="http://www.w3.org/2000/svg">
					%s
				</svg>
			</g>
			<text font-size="10px" y="270" x="490" text-anchor="end">
				%s
			</text>

			<rect y="25" x="485" width="145" height="60" rx="5" style="stroke:gray;stroke-width:5;fill-opacity:0;"/>
			<text font-size="40px" font-family="Roboto Mono" y="70" x="505">
				%s
			</text>
			
			<text font-size="35px" y="160" x="490" text-anchor="middle">Score:</text>
			<text font-size="90px" y="230" x="630" text-anchor="middle">%d</text>

			<text y="360" x="20" font-size="30px">
				<tspan x="20" dy="2em">AWAIR</tspan>
				<tspan x="20" dy="4em">Nature</tspan>
				<tspan x="40" dy="1em">Remo</tspan>
				<tspan x="20" dy="4em">SwitchBot</tspan>
			</text>
			
			<text y="360" x="170" font-size="90px">
				<tspan x="200" dy="-1em" dx="0.5em" font-size="30px">Temperature</tspan>
				<tspan x="200" dy="1.2em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">°C</tspan>
				</tspan>
				<tspan x="200" dy="1.5em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">°C</tspan>
				</tspan>
				<tspan x="200" dy="1.5em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">°C</tspan>
				</tspan>
			</text>

			<text y="360" x="520" font-size="90px">
				<tspan x="490" dy="-1em" dx="1.2em" font-size="30px">Humidity</tspan>
				<tspan x="490" dy="1.2em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">%%</tspan>
				</tspan>
				<tspan x="490" dy="1.5em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">%%</tspan>
				</tspan>
				<tspan x="490" dy="1.5em">
					%s<tspan font-size="40px" dx="-0.0em" text-anchor="end">%%</tspan>
				</tspan>
			</text>
			
			<text font-size="35px" y="780" x="250" text-anchor="middle">Electric Energy:</text>
			<text font-size="90px" y="850" x="565" text-anchor="end">%d</text>
			<text font-size="64px" y="850" x="570" text-anchor="start">W</text>

			<text font-size="35px" y="923" x="50" text-anchor="middle">CO<tspan baseline-shift="sub" font-size="25">2</tspan>:</text>
			<text font-size="64px" y="993" x="180" text-anchor="end">
				%d
			</text>
			<text font-size="35px" y="983" x="220" text-anchor="middle">ppm</text>

			<text font-size="35px" y="923" x="300" text-anchor="middle">VOC:</text>
			<text font-size="64px" y="993" x="440" text-anchor="end">
				%d
			</text>
			<text font-size="35px" y="983" x="475" text-anchor="middle">ppb</text>

			<text font-size="35px" y="923" x="575" text-anchor="middle">PM2.5:</text>
			<text font-size="64px" y="993" x="630" text-anchor="end">
				%s
			</text>
			<text font-size="35px" y="983" x="690" text-anchor="middle">μg/m<tspan baseline-shift="super" font-size="25px">3</tspan></text>
		</g>
</svg>`

func GeneratePNGImage(
	result usecase.AllRoomInfo,
	now time.Time,
) (bytes.Buffer, error) {
	var buf bytes.Buffer

	weatherIcon, err := ConvertToIcon(result.Weather.Icon)
	if err != nil {
		return buf, err
	}

	colors := presenter.CalculateColors(now)

	// Generate your SVG data here
	svgString := fmt.Sprintf(
		svgStringFormat,
		presenter.Width,
		presenter.Height,
		colors.Text,
		colors.Text,
		weatherIcon,
		result.Weather.Datetime.Format(time.RFC3339),
		now.Format("15:04"),
		result.AwairRoomInfo.Score,
		floatSawedOffString(result.AwairRoomInfo.Temperature),
		floatSawedOffString(result.NatureRemoRoomInfo.Temperature),
		floatSawedOffString(result.SwitchBotMeterInfo.Temperature),
		floatSawedOffString(result.AwairRoomInfo.Humidity),
		floatSawedOffString(result.NatureRemoRoomInfo.Humidity),
		floatSawedOffString(result.SwitchBotMeterInfo.Humidity),
		result.NatureRemoRoomInfo.ElectricEnergy,
		result.AwairRoomInfo.Co2,
		result.AwairRoomInfo.Voc,
		floatSawedOffString(result.AwairRoomInfo.Pm25),
	)
	buf, err = presenter.ConvertSVGToPNG(svgString, colors.Bg, presenter.Width, presenter.Height)

	return buf, nil
}

func floatSawedOffString[A ~float32](d A) string {
	return fmt.Sprintf("%2.1f", d)
}
