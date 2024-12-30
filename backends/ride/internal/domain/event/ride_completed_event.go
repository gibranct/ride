package event

type RideCompletedEvent struct {
	RideId string  `json:"ride_id"`
	Fare   float64 `json:"fare"`
}

func (r *RideCompletedEvent) GetName() string {
	return "ride.completed"
}

func NewFinishRideEvent(rideId string, fare float64) *RideCompletedEvent {
	return &RideCompletedEvent{
		RideId: rideId,
		Fare:   fare,
	}
}
