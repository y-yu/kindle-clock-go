package domain

import "time"

type Clock interface {
	Now() time.Time
}

type SystemClock struct {
	timezone *time.Location
}

const locationAsiaTokyo = "Asia/Tokyo"

var jst = time.FixedZone(locationAsiaTokyo, 9*60*60)

func NewSystemClock() *SystemClock {
	return &SystemClock{
		timezone: jst,
	}
}

func (*SystemClock) Now() time.Time {
	return time.Now().In(jst)
}
