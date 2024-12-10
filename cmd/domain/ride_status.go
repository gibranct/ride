package domain

import (
	"errors"
)

var (
	errInvalidStatus = errors.New("invalid status")
)

const (
	REQUESTED_RIDE_STATUS   string = "requested"
	ACCEPTED_RIDE_STATUS    string = "accepted"
	IN_PROGRESS_RIDE_STATUS string = "in_progress"
	COMPLETED_RIDE_STATUS   string = "completed"
)

type RideStatus interface {
	GetValue() string
	Request() error
	Accept() error
	Start() error
}

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

func (rq *RequestedStatus) GetValue() string {
	return rq.Value
}

func NewRequestedStatus(ride *Ride) *RequestedStatus {
	return &RequestedStatus{
		ride:  ride,
		Value: REQUESTED_RIDE_STATUS,
	}
}

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

func (rq *AcceptedStatus) GetValue() string {
	return rq.Value
}

func NewAcceptedStatus(ride *Ride) *AcceptedStatus {
	return &AcceptedStatus{
		ride:  ride,
		Value: ACCEPTED_RIDE_STATUS,
	}
}

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

func NewInProgressStatus(ride *Ride) *InProgressStatus {
	return &InProgressStatus{
		ride:  ride,
		Value: IN_PROGRESS_RIDE_STATUS,
	}
}

func NewRideStatus(status string, ride *Ride) (RideStatus, error) {
	if status == REQUESTED_RIDE_STATUS {
		return NewRequestedStatus(ride), nil
	}
	if status == ACCEPTED_RIDE_STATUS {
		return NewAcceptedStatus(ride), nil
	}
	if status == IN_PROGRESS_RIDE_STATUS {
		return NewInProgressStatus(ride), nil
	}
	return nil, errInvalidStatus
}
