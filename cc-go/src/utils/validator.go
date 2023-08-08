package utils

import (
	"strconv"
)

func CreditCardValidator(input string) bool {
	arr := []int{}

	// Parse the input to integer then push it to arr (slice)
	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return false // Handle error or return an error flag
		}
		arr = append(arr, num)
	}

	// The logic for the Luhn Algorithm
	for i := len(arr) - 2; i >= 0; i = i - 2 {
		temp := arr[i]
		temp = temp * 2
		if temp > 9 {
			temp = (temp % 10) + 1
		}
		arr[i] = temp
	}

	// Sum
	count := 0
	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}

	return count%10 == 0
}
