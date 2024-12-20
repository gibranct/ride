package domain

type InProgressStatus struct {
	ride  *Ride
	Value string
}

func (rq *InProgressStatus) Request() error {
	return errInvalidStatus
}

func (rq *InProgressStatus) Accept() error {
	return errInvalidStatus
}

func (rq *InProgressStatus) Start() error {
	return errInvalidStatus
}

func (rq *InProgressStatus) GetValue() string {
	return rq.Value
}

func (rq *InProgressStatus) Finish() error {
	rq.ride.setStatus(NewCompleteStatus(rq.ride))
	return nil
}

func NewInProgressStatus(ride *Ride) *InProgressStatus {
	return &InProgressStatus{
		ride:  ride,
		Value: IN_PROGRESS_RIDE_STATUS,
	}
}
