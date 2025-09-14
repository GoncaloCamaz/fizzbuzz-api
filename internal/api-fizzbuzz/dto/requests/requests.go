/*
Package requests contains data transfer objects (DTOs) for handling incoming API requests.
*/
package requests

import (
	"fizzbuzz-api/internal/api-fizzbuzz/errors"
	"fmt"
)

// FizzBuzzRequest represents the request payload for the FizzBuzz operation
type FizzBuzzRequest struct {
	Limit                int64  `json:"limit"`
	FirstNumber          int64  `json:"first_number"`
	SecondNumber         int64  `json:"second_number"`
	FirstReplacementStr  string `json:"first_replacement_str"`
	SecondReplacementStr string `json:"second_replacement_str"`
}

// ValidateFizzBuzzParams validates the FizzBuzzRequest parameters
func (f FizzBuzzRequest) ValidateFizzBuzzParams() error {
	if f.Limit < 0 {
		return fmt.Errorf(errors.ErrLimitNegative)
	}

	if f.FirstNumber <= 0 || f.SecondNumber <= 0 {
		return fmt.Errorf(errors.ErrMultipleNonPositive)
	}

	return nil
}
