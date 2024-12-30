package domain

type RequestedStatus struct {
	ride  *Ride
	Value string
}

func (rq *RequestedStatus) Request() error {
	return errInvalidStatus
}

func (rq *RequestedStatus) Accept() error {
	rq.ride.setStatus(NewAcceptedStatus(rq.ride))
	return nil
}

func (rq *RequestedStatus) Start() error {
	return errInvalidStatus
}

func (rq *RequestedStatus) Finish() error {
	return errInvalidStatus
}

func (rq *RequestedStatus) GetValue() string {
	return rq.Value
}

func NewRequestedStatus(ride *Ride) *RequestedStatus {
	return &RequestedStatus{
		ride:  ride,
		Value: REQUESTED_RIDE_STATUS,
	}
}
