package event

type RideCompletedEvent struct {
	RideId string
	Fare   float64
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
