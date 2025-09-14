package utils

import "fmt"

func FizzBuzz(limit, firstNumber, secondNumber int64, firstReplacementStr, secondReplacementStr string) []string {
	result := make([]string, 0, limit)
	var i int64
	for i = 1; i <= limit; i++ {
		switch {
		case i%firstNumber == 0 && i%secondNumber == 0:
			result = append(result, firstReplacementStr+secondReplacementStr)
		case i%firstNumber == 0:
			result = append(result, firstReplacementStr)
		case i%secondNumber == 0:
			result = append(result, secondReplacementStr)
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}

	}

	return result
}
