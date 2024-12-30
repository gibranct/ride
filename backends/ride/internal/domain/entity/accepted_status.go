package domain

type AcceptedStatus struct {
	ride  *Ride
	Value string
}

func (rq *AcceptedStatus) Request() error {
	return errInvalidStatus
}

func (rq *AcceptedStatus) Accept() error {
	return errInvalidStatus
}

func (rq *AcceptedStatus) Start() error {
	rq.ride.setStatus(NewInProgressStatus(rq.ride))
	return nil
}

func (rq *AcceptedStatus) Finish() error {
	return errInvalidStatus
}

func (rq *AcceptedStatus) GetValue() string {
	return rq.Value
}

func NewAcceptedStatus(ride *Ride) *AcceptedStatus {
	return &AcceptedStatus{
		ride:  ride,
		Value: ACCEPTED_RIDE_STATUS,
	}
}
