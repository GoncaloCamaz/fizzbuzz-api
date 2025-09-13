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
	Limit              int64  `json:"limit"`
	Multiple1          int64  `json:"multiple1"`
	Multiple2          int64  `json:"multiple2"`
	ReplacementString1 string `json:"replacement_string1"`
	ReplacementString2 string `json:"replacement_string2"`
}

// ValidateFizzBuzzParams validates the FizzBuzzRequest parameters
func (f FizzBuzzRequest) ValidateFizzBuzzParams() error {
	if f.Limit < 0 {
		return fmt.Errorf(errors.ErrLimitNegative)
	}

	if f.Multiple1 <= 0 || f.Multiple2 <= 0 {
		return fmt.Errorf(errors.ErrMultipleNonPositive)
	}

	return nil
}
