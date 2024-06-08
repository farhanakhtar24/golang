package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("Error reading file")
	}

	balanceData := string(data)
	value, err := strconv.ParseFloat(balanceData, 64)

	if err != nil {
		return 1000, errors.New("Failed to convert parse file data to float")
	}

	return value, nil
}
