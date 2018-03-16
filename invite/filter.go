package invite

// Filter holads filter parameters
type Filter struct {
	offLoc  Location
	maxDist float64
}

// SetOffLoc sets office location
func (f *Filter) SetOffLoc(lat float64, lng float64) {
	f.offLoc.Lat = number(lat)
	f.offLoc.Lng = number(lng)
}

// SetMaxDist sets maximum distance from office location
func (f *Filter) SetMaxDist(dist float64) {
	f.maxDist = dist
}
