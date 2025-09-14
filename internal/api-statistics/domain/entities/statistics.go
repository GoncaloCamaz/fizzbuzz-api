/*
Package entities defines the data structures used in the statistics domain.
*/
package entities

// Statistics represents the statistics entity
type Statistics struct {
	ID                   string `json:"id"`
	FirstNumber          int64  `json:"first_number"`
	SecondNumber         int64  `json:"second_number"`
	FirstReplacementStr  string `json:"first_replacement_str"`
	SecondReplacementStr string `json:"second_replacement_str"`
	Limit                int64  `json:"limit"`
}
