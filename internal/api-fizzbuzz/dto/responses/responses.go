/*
Package responses contains the response DTOs for the FizzBuzz API.
*/
package responses

// FizzBuzzResponse represents the response payload for the FizzBuzz operation
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

// SerializeFizzBuzzResponse serializes the FizzBuzz result into a FizzBuzzResponse DTO
func SerializeFizzBuzzResponse(result []string) *FizzBuzzResponse {
	return &FizzBuzzResponse{
		Result: result,
	}
}
