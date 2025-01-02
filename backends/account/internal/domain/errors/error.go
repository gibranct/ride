package errors

import "errors"

type ErrDomain struct {
	message string
}

func (e *ErrDomain) Error() string {
	return e.message
}

func (e *ErrDomain) Name() string {
	return "ErrDomain"
}

func NewErrorDomain(message string) *ErrDomain {
	return &ErrDomain{message: message}
}

var (
	ErrInvalidName       = NewErrorDomain("invalid name")
	ErrInvalidEmail      = NewErrorDomain("invalid email")
	ErrInvalidCPF        = NewErrorDomain("invalid CPF")
	ErrPasswordTooShort  = NewErrorDomain("password must be greater than 5 characters")
	ErrInvalidCarPlate   = NewErrorDomain("invalid car plate")
	ErrEmailAlreadyTaken = NewErrorDomain("email already taken")
)

type ErrInfrastructure struct {
	message string
}

func (e *ErrInfrastructure) Error() string {
	return e.message
}

func (e *ErrInfrastructure) Name() string {
	return "ErrorInfrastructure"
}

func NewErrorInfrastructure(message string) *ErrInfrastructure {
	return &ErrInfrastructure{message: message}
}

var (
	ErrSavingAccount = NewErrorInfrastructure("error saving account to database")
	ErrDatabase      = NewErrorInfrastructure("error connecting to database")
)

func AllDomainErrors() []error {
	return []error{
		ErrInvalidName,
		ErrInvalidEmail,
		ErrInvalidCPF,
		ErrPasswordTooShort,
		ErrInvalidCarPlate,
		ErrEmailAlreadyTaken,
	}
}

func AllInfraErrors() []error {
	return []error{
		ErrSavingAccount,
		ErrDatabase,
	}
}

var (
	ErrAccountNotFound = errors.New("account not found")
)
