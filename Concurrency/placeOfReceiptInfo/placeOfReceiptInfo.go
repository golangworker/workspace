package placeofreceiptinfo

import (
	"math/rand"
	"fmt"
)

type coordinates struct {
	Latitude float64 // широта: от -90 до 90
	Longitude float64 // долгота: от -180 до 180
}

func RandomPlace() string {
	lat := rand.Float64()*180 - 90 // от -90 да 90
	lon := rand.Float64()*360 - 180 // от -180 до 180

	c := coordinates{
		Latitude: lat,
		Longitude: lon,
	}

	return fmt.Sprintf("GEO: `lat: %.2f, lon: %.2f`", c.Latitude, c.Longitude)
}
