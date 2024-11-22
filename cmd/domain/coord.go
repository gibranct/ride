package domain

import (
	"errors"
)

type Coord struct {
	lat  float64
	long float64
}

func NewCoord(lat, long float64) (*Coord, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("invalid latitude")
	}
	if long < -180 || long > 180 {
		return nil, errors.New("invalid longitude")
	}

	return &Coord{
		lat:  lat,
		long: long,
	}, nil
}

func (c *Coord) GetLat() float64 {
	return c.lat
}

func (c *Coord) GetLong() float64 {
	return c.long
}
