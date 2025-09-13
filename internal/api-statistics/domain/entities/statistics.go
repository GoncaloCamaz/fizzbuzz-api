/*
Package entities defines the data structures used in the statistics domain.
*/
package entities

// Statistics represents the statistics entity
type Statistics struct {
	ID              string `json:"id"`
	MultipleOne     int64  `json:"multiple_one"`
	MultipleTwo     int64  `json:"multiple_two"`
	ReplacementStr1 string `json:"replacement_str1"`
	ReplacementStr2 string `json:"replacement_str2"`
	Limit           int64  `json:"limit"`
}
