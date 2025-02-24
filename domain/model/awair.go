package model

type Score int32

type Co2 int32

type Voc int32

type Pm25 float32

type AwairRoomInfo struct {
	Score       Score
	Temperature Temperature
	Humidity    Humidity
	Co2         Co2
	Voc         Voc
	Pm25        Pm25
}
