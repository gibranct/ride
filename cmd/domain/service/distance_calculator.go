package service

import (
	"math"

	"github.com.br/gibranct/ride/cmd/domain/vo"
)

const earthRadiusKm = 6371.0
const degreeToRadian = math.Pi / 180

type DistanceCalculator struct{}

func NewDistanceCalculator() *DistanceCalculator {
	return &DistanceCalculator{}
}

func (dc *DistanceCalculator) Calculate(from, to vo.Coord) float64 {
	deltaLat := (to.GetLat() - from.GetLat()) * degreeToRadian
	deltaLong := (to.GetLong() - from.GetLong()) * degreeToRadian
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(from.GetLat()*degreeToRadian)*math.Cos(to.GetLat()*degreeToRadian)*
			math.Sin(deltaLong/2)*math.Sin(deltaLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadiusKm * c
	return math.Round(distance)
}
