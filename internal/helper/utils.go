package helper

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {
	const numbers = "1234567890"

	if length <= 0 {
		return 0, nil
	}

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	numLength := len(numbers)

	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%numLength]
	}

	result, err := strconv.Atoi(string(buffer))
	if err != nil {
		return 0, err
	}

	return result, nil
}
