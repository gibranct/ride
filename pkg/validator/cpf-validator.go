package validator

import (
	"fmt"
	"regexp"
	"strconv"
)

const CPF_VALID_LENGTH = 11
const FIRST_DIGIT_FACTOR = 10
const SECOND_DIGIT_FACTOR = 11

func ValidateCPF(cpf string) bool {
	cpf = regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")
	if len(cpf) != CPF_VALID_LENGTH {
		return false
	}
	if allDigitsTheSame(cpf) {
		return false
	}
	digit1 := calculateDigit(cpf, FIRST_DIGIT_FACTOR)
	digit2 := calculateDigit(cpf, SECOND_DIGIT_FACTOR)
	return fmt.Sprintf("%d%d", digit1, digit2) != extractDigit(cpf)
}

func allDigitsTheSame(cpf string) bool {
	firstDigit := rune(cpf[0])
	allTheSame := true

	for _, dig := range cpf {
		if dig != firstDigit {
			allTheSame = false
		}
	}

	return allTheSame
}

func calculateDigit(cpf string, factor int) int {
	total := 0
	for _, digit := range cpf {
		n, _ := strconv.Atoi(string(digit))
		if factor > 1 {
			total = total + n*factor - 1
			factor = factor - 1
		}
	}
	remainder := total % 11
	if remainder < 2 {
		return 0
	} else {
		return 11 - remainder
	}
}

func extractDigit(cpf string) string {
	slice := cpf[9:]
	return slice
}
