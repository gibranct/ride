package service

import "time"

type FareCalculator interface {
	Calculate(distance float64) float64
}

type NormalFareCalculator struct{}

func (f *NormalFareCalculator) Calculate(distance float64) float64 {
	return distance * 2.1
}

type OvernightFareCalculator struct{}

func (f *OvernightFareCalculator) Calculate(distance float64) float64 {
	return distance * 3.9
}

type SpecialDayFareCalculator struct{}

func (f *SpecialDayFareCalculator) Calculate(distance float64) float64 {
	return distance * 1
}

func NewFareCalculator(date time.Time) FareCalculator {
	if date.Day() == 1 {
		return &SpecialDayFareCalculator{}
	}

	if date.Hour() >= 22 || date.Hour() < 6 {
		return &OvernightFareCalculator{}
	}

	return &NormalFareCalculator{}
}
