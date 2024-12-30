package domain

type CompleteStatus struct {
	ride  *Ride
	Value string
}

func (rq *CompleteStatus) Request() error {
	return errInvalidStatus
}

func (rq *CompleteStatus) Accept() error {
	return errInvalidStatus
}

func (rq *CompleteStatus) Start() error {
	return errInvalidStatus
}

func (rq *CompleteStatus) GetValue() string {
	return rq.Value
}

func (rq *CompleteStatus) Finish() error {
	return errInvalidStatus
}

func NewCompleteStatus(ride *Ride) *CompleteStatus {
	return &CompleteStatus{
		ride:  ride,
		Value: COMPLETED_RIDE_STATUS,
	}
}
