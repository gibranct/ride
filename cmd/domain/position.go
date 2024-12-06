package domain

import (
	"time"

	"github.com.br/gibranct/ride/cmd/domain/vo"
	"github.com/google/uuid"
)

type Position struct {
	positionId string
	rideId     string
	coord      *vo.Coord
	Date       *time.Time
}

func (p *Position) SetCoord(lat, long float64) error {
	newCoord, err := vo.NewCoord(lat, long)
	if err != nil {
		return err
	}
	p.coord = newCoord

	return nil
}

func (p *Position) GetPositionId() string {
	return p.positionId
}

func (p *Position) GetRideId() string {
	return p.rideId
}

func (p *Position) GetCoord() vo.Coord {
	return *p.coord
}

func NewPosition(positionId, rideId string, lat, long float64, date time.Time) (*Position, error) {
	newCoord, err := vo.NewCoord(lat, long)
	if err != nil {
		return nil, err
	}
	return &Position{
		positionId: positionId,
		rideId:     rideId,
		coord:      newCoord,
		Date:       &date,
	}, nil
}

func CreatePosition(rideId string, lat, long float64) (*Position, error) {
	positionId := uuid.NewString()
	date := time.Now()
	return NewPosition(positionId, rideId, lat, long, date)
}
