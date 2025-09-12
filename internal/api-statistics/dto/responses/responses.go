package responses

import "fizzbuzz-api/internal/api-statistics/datamodel"

// FizzBuzzRequest represents the request payload for a FizzBuzz operation
type FizzBuzzRequest struct {
	Limit              int64  `json:"limit"`
	Multiple1          int64  `json:"multiple1"`
	Multiple2          int64  `json:"multiple2"`
	ReplacementString1 string `json:"replacement_string1"`
	ReplacementString2 string `json:"replacement_string2"`
}

// MostFrequentRequestResponse represents the response payload for the most frequent FizzBuzz request
type MostFrequentRequestResponse struct {
	Request FizzBuzzRequest `json:"request"`
	Count   int             `json:"count"`
}

// SerializeMostFrequentRequestResponse serializes the most frequent request and its count into a DTO
func SerializeMostFrequentRequestResponse(data datamodel.Statistic, count int) *MostFrequentRequestResponse {
	return &MostFrequentRequestResponse{
		Request: FizzBuzzRequest{
			Limit:              data.Limit,
			Multiple1:          data.MultipleOne,
			Multiple2:          data.MultipleTwo,
			ReplacementString1: data.ReplacementStr1,
			ReplacementString2: data.ReplacementStr2,
		},
		Count: count,
	}
}
