package utility

import (
	"regexp"
)

func IsAlpha(data string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	if isAlpha(data) {
		return true
	} else {
		return false
	}
}

func IsDigit(data string) bool {
	isDigit := regexp.MustCompile(`^[0-9]+$`).MatchString

	if isDigit(data) {
		return true
	} else {
		return false
	}	
}

func IsCapital(data string) bool {
	isCapital := regexp.MustCompile(`^[A-Z]+$`).MatchString

	if isCapital(data) {
		return true
	} else {
		return false
	}		
}

func Swap [T any] ( data []T, i,j int) {
	data[i], data[j] = data[j], data[i]
}