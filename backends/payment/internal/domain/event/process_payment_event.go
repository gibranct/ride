package event

type ProcessPaymentEvent struct {
	RideId string  `json:"ride_id"`
	Fare   float64 `json:"fare"`
}

func (r *ProcessPaymentEvent) GetName() string {
	return "ride.completed"
}

func NewFinishRideEvent(rideId string, fare float64) *ProcessPaymentEvent {
	return &ProcessPaymentEvent{
		RideId: rideId,
		Fare:   fare,
	}
}
