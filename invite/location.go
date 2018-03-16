package invite

import (
	"strconv"
	"strings"

	geo "github.com/kellydunn/golang-geo"
	"github.com/pkg/errors"
)

type number float64

func (n *number) UnmarshalJSON(b []byte) error {
	v := string(b)
	v = strings.Trim(v, "\"")
	f, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return errors.Wrap(err, "invit.number.UnmarshalJSON")
	}

	*n = number(f)
	return nil
}

// Location holds gps location
type Location struct {
	Lat number `json:"latitude"`
	Lng number `json:"longitude"`
}

// KiloDist returns distance between two location
func (l Location) KiloDist(com Location) float64 {
	p := geo.NewPoint(float64(l.Lat), float64(l.Lng))
	p2 := geo.NewPoint(float64(com.Lat), float64(com.Lng))

	return p.GreatCircleDistance(p2)
}
