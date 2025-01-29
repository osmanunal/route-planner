package model

import (
	"math"
	"route-planner/pkg/model"

	"github.com/uber/h3-go"
)

const (
	EarthRadius = 6371.0
	toRad       = math.Pi / 180.0
	// H3Resolution https://h3geo.org/docs/core-library/restable
	H3Resolution = 8 // 0.737327598km^2
)

type Location struct {
	model.BaseModel

	Name      string
	Latitude  float64
	Longitude float64
	Color     string
	H3Index   h3.H3Index

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

// Konumun enlem ve boylam koordinatlarını yapılandırılmış H3Resolution seviyesinde
// H3 coğrafi indeksine dönüştürür ve H3Index alanında saklar
func (m *Location) SetH3Index() {
	geo := h3.GeoCoord{
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	}
	m.H3Index = h3.FromGeo(geo, H3Resolution)
}

// Konumun H3 indeksinden 3 halka (k=3) mesafe içindeki tüm hücrelerin H3 indekslerini,
// konumun kendi indeksi de dahil olmak üzere dizi olarak döndürür
func (m *Location) GetNeighborIndexes() []h3.H3Index {
	return h3.KRing(m.H3Index, 3)
}
