package usecase

import "github.com.br/gibranct/ride/cmd/infra/repository"

type GetRide struct {
	rideRepository repository.RideRepository
}

type GetRideOutput struct {
	RideId      string
	PassengerId string
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
	Status      string
}

func (gr *GetRide) Execute(rideId string) (*GetRideOutput, error) {
	ride, err := gr.rideRepository.GetRideByID(rideId)
	if err != nil {
		return nil, err
	}

	return &GetRideOutput{
		RideId:      ride.GetRideId(),
		PassengerId: ride.GetPassengerId(),
		FromLat:     ride.GetFromCoord().GetLat(),
		FromLong:    ride.GetFromCoord().GetLat(),
		ToLat:       ride.GetToCoord().GetLat(),
		ToLong:      ride.GetToCoord().GetLat(),
		Status:      ride.GetStatus(),
	}, nil
}
