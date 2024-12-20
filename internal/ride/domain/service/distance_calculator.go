package service

import (
	"math"

	"github.com.br/gibranct/ride/internal/ride/domain/vo"
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

func (dc *DistanceCalculator) CalculateByPositions(positions []Position) float64 {
	distance := 0.0
	for idx, pos := range positions {
		if idx >= len(positions)-1 {
			continue
		}
		nextPosition := positions[idx+1]
		distance += dc.Calculate(pos.GetCoord(), nextPosition.GetCoord())
	}
	return distance
}

type Position interface {
	GetCoord() vo.Coord
}
