package model

import (
	"math"
	"route-planner/pkg/model"
)

const (
	EarthRadius = 6371.0
	toRad       = math.Pi / 180.0
)

type Location struct {
	model.BaseModel

	Name      string
	Latitude  float64
	Longitude float64
	Color     string

	Distance float64 `bun:"-"`
}

func (m *Location) HaversineDistance(lat, lon float64) {
	lat1 := m.Latitude * toRad
	lon1 := m.Longitude * toRad
	lat2 := lat * toRad
	lon2 := lon * toRad

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	m.Distance = EarthRadius * c
}
