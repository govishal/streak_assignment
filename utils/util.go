package utils

import (
	"streak_assignment/models"
)

// Validate Inputs
func ValidateInput(reqBody models.RequestBody) bool {
	if reqBody.Numbers == nil || len(reqBody.Numbers) == 0 {
		return false
	}
	return true
}
