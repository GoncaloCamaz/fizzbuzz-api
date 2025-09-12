package utils

import "fmt"

func FizzBuzz(limit, multiple1, multiple2 int, replacementString1, replacementString2 string) []string {
	result := make([]string, 0, limit)
	for i := 1; i <= limit; i++ {
		switch {
		case i%multiple1 == 0 && i%multiple2 == 0:
			result = append(result, replacementString1+replacementString2)
		case i%multiple1 == 0:
			result = append(result, replacementString1)
		case i%multiple2 == 0:
			result = append(result, replacementString2)
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}

	}
	
	return result
}
