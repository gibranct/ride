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
	Finish() error
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
	if status == COMPLETED_RIDE_STATUS {
		return NewCompleteStatus(ride), nil
	}
	return nil, errInvalidStatus
}
