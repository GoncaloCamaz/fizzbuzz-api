package responses

import "fizzbuzz-api/internal/api-statistics/datamodel"

// FizzBuzzRequest represents the request payload for a FizzBuzz operation
type FizzBuzzRequest struct {
	Limit                int64  `json:"limit"`
	FirstNumber          int64  `json:"first_number"`
	SecondNumber         int64  `json:"second_number"`
	FirstReplacementStr  string `json:"first_replacement_str"`
	SecondReplacementStr string `json:"second_replacement_str"`
}

// MostFrequentRequestResponse represents the response payload for the most frequent FizzBuzz request
type MostFrequentRequestResponse struct {
	Request FizzBuzzRequest `json:"request"`
	Count   int             `json:"count"`
}

// SerializeMostFrequentRequestResponse serializes the most frequent request and its count into a DTO
func SerializeMostFrequentRequestResponse(data *datamodel.Statistic, count int) *MostFrequentRequestResponse {
	return &MostFrequentRequestResponse{
		Request: FizzBuzzRequest{
			Limit:                data.Limit,
			FirstNumber:          data.FirstNumber,
			SecondNumber:         data.SecondNumber,
			FirstReplacementStr:  data.FirstReplacementStr,
			SecondReplacementStr: data.SecondReplacementStr,
		},
		Count: count,
	}
}
