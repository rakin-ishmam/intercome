package invite

import geo "github.com/kellydunn/golang-geo"

// Location holds gps location
type Location struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// KiloDist returns distance between two location
func (l Location) KiloDist(com Location) float64 {
	p := geo.NewPoint(l.Lat, l.Lng)
	p2 := geo.NewPoint(com.Lat, com.Lng)

	return p.GreatCircleDistance(p2)
}
